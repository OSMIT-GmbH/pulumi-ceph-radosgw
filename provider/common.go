// Copyright 2023, OSMIT GmbH
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package provider

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/ceph/go-ceph/rgw/admin"
	p "github.com/pulumi/pulumi-go-provider"
	"github.com/pulumi/pulumi-go-provider/infer"
	"github.com/pulumi/pulumi/sdk/v3/go/common/diag"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// I'm not sure what's wrong with boolean - see following error:
// error: pulumi:providers: resource 'xyz-provider': property assimilate value {false} has a problem: Field 'assimilate' on 'provider.ProviderConfig' must be a 'bool'; got 'string' instead
type ProviderConfig struct {
	AccessKeyID       string `pulumi:"accessKeyID"`
	SecretAccessKey   string `pulumi:"secretAccessKey" provider:"secret"`
	Endpoint          string `pulumi:"endpoint"`
	Insecure          string `pulumi:"insecure,optional"`
	insecure          bool
	Assimilate        string `pulumi:"assimilate,optional"`
	assimilate        bool
	DeleteAssimilated string `pulumi:"deleteAssimilated,optional"`
	deleteAssimilated bool
	Version           string `pulumi:"version,optional"` // version seems to be provided automatically
	cacheKey          string
}

var _ = (infer.Annotated)((*ProviderConfig)(nil))

func (c *ProviderConfig) Annotate(a infer.Annotator) {
	a.Describe(&c.AccessKeyID, "The username. It's important but not secret.")
	a.Describe(&c.SecretAccessKey, "The password. It is very secret.")
	a.Describe(&c.Endpoint, `The URI to the API`)
	a.Describe(&c.Insecure, `Don't validate server SSL certificate`)
	a.Describe(&c.Assimilate, `Assimilate an existing object during create`)
	a.Describe(&c.DeleteAssimilated, `Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)`)
	// a.SetDefault(&c.Insecure, false)
}

var _ = (infer.CustomConfigure)((*ProviderConfig)(nil))

func (c *ProviderConfig) Configure(ctx p.Context) error {
	// apiUrl, err := url.Parse(c.Uri)

	//if err != nil {
	//	// return nil, errors.Wrap(err, "could not parse ZtAPI from configuration as URI")
	//	// fmt.Errorf("no session token returned from login request to %v. Received: %v", c.Uri, zitiLogin.String())
	//	return err
	//}
	c.cacheKey = fmt.Sprintf("%s:%s:%s", c.Endpoint, c.AccessKeyID, c.SecretAccessKey)
	c.assimilate = strings.EqualFold(c.Assimilate, "true")
	c.deleteAssimilated = strings.EqualFold(c.DeleteAssimilated, "true")
	c.insecure = strings.EqualFold(c.Insecure, "true")

	//ctx.Log(diag.Info, msg)
	return nil
}

var _ = (infer.CustomCheck[ProviderConfig])((*ProviderConfig)(nil))

func (*ProviderConfig) Check(ctx p.Context, name string, oldInputs, newInputs resource.PropertyMap) (ProviderConfig, []p.CheckFailure, error) {
	ctx.Logf(diag.Warning, "Hello world from Check") // TODO dosn't get called

	if booleanValue, ok := newInputs["insecure"]; ok {
		newInputs["insecure"] = resource.NewBoolProperty(strings.EqualFold(booleanValue.StringValue(), "true"))
	}
	ret, failures, err := infer.DefaultCheck[ProviderConfig](newInputs)
	return ret, failures, err
}

// TODO https://github.com/pulumi/pulumi-go-provider/issues/121
//var _ = (infer.CustomDiff[*ProviderConfig, *ProviderConfig])((*ProviderConfig)(nil))
//
//func (*ProviderConfig) Diff(ctx p.Context, id string, olds *ProviderConfig, news *ProviderConfig) (p.DiffResponse, error) {
//	fmt.Printf("Config Diff called: %v => %v", olds, news)
//	return p.DiffResponse{
//		DeleteBeforeReplace: true,
//		HasChanges:          false,
//		DetailedDiff:        nil,
//	}, nil
//}

var IdPreviewPrefix = "~~preview~~"

type CacheEntry struct {
	client *admin.API
	s3     *s3.Client
	// configTypesMutex sync.Mutex
}

var cache = make(map[string]*CacheEntry)

var clientMutex sync.Mutex = sync.Mutex{}

func initClient(ctx p.Context) (*CacheEntry, ProviderConfig, error) {
	c := infer.GetConfig[ProviderConfig](ctx)
	ce, ok := cache[c.cacheKey]
	if !ok {
		// new entry - use mutex to limit to one session

		clientMutex.Lock()
		ce, ok = cache[c.cacheKey]
		if !ok {
			var httpClient *http.Client
			if c.insecure {
				tr := &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
				httpClient = &http.Client{Transport: tr}
			} else {
				httpClient = http.DefaultClient
			}
			client, err := admin.New(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, httpClient)
			if err != nil {
				clientMutex.Unlock()
				return nil, c, err
			}

			s3client := s3.New(s3.Options{
				Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
					return aws.Credentials{
						AccessKeyID:     c.AccessKeyID,
						SecretAccessKey: c.SecretAccessKey,
					}, nil
				}),
				EndpointResolver: s3.EndpointResolverFromURL(c.Endpoint),
				UsePathStyle:     true,
				HTTPClient:       httpClient,
			})
			// fmt.Printf("identity name: %#v; token: s\n", client)
			ce = &CacheEntry{
				client: client,
				s3:     s3client,
			}
			cache[c.cacheKey] = ce
		}
		clientMutex.Unlock()
	}
	return ce, c, nil
}

func initS3(ctx p.Context) (*CacheEntry, ProviderConfig, error) {
	c := infer.GetConfig[ProviderConfig](ctx)
	ce, ok := cache[c.cacheKey]
	if !ok {
		// new entry - use mutex to limit to one session

		clientMutex.Lock()
		ce, ok = cache[c.cacheKey]
		if !ok {
			var httpClient *http.Client
			if c.insecure {
				tr := &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				}
				httpClient = &http.Client{Transport: tr}
			} else {
				httpClient = http.DefaultClient
			}
			client, err := admin.New(c.Endpoint, c.AccessKeyID, c.SecretAccessKey, httpClient)
			if err != nil {
				clientMutex.Unlock()
				return nil, c, err
			}

			// fmt.Printf("identity name: %#v; token: s\n", client)
			ce = &CacheEntry{
				client: client,
			}
			cache[c.cacheKey] = ce
		}
		clientMutex.Unlock()
	}
	return ce, c, nil
}

//	func formatApiErr(ctx p.Context, err error, apiError *rest_model.APIErrorEnvelope) error {
//		// if errors.Is(err, config.CreateConfigBadRequest) {
//		errOut, err2 := json.Marshal(apiError.Error)
//		if err2 != nil {
//			ctx.Logf(diag.Error, "ERROR: type: ErrorString: %v, MarshallingError: %s Payload.Error=%+v  PalyloadErrorCause=%#v PayloadMetadata: %+v\n", err.Error(), err2.Error(), apiError.Error, apiError.Error.Cause, apiError.Meta)
//			return errors.Join(err, err2)
//		}
//		return fmt.Errorf("ERROR: type: ErrorString: %s, Payload=%s\n", err.Error(), string(errOut))
//	}
//
//	func formatApiErrDupeCheck(ctx p.Context, err error, apiError *rest_model.APIErrorEnvelope) (error, bool) {
//		errRet := formatApiErr(ctx, err, apiError)
//		match, _ := regexp.MatchString(" Payload=\\{\"cause\":\\{\"field\":\"name\",\"reason\":\"duplicate value '[^\"']+' in unique index on identities store\",\"value\":\"[^\"']+\"},", errRet.Error())
//		if !match {
//			match, _ = regexp.MatchString(" Payload={\"cause\":{\"field\":\"name\",\"reason\":\"name is must be unique\",\"value\":\"[^\"']+\"},\"code\":\"COULD_NOT_VALIDATE\",", errRet.Error())
//		}
//		return errRet, match
//	}
// func handleDeleteErr(ctx p.Context, err error, id string, typeName string) error {
// 	var apiError *runtime.APIError
// 	if errors.As(err, &apiError) {
// 		if apiError.Code == 404 {
// 			ctx.Logf(diag.Warning, "DELETE on %s[%s] returned 404 - assuming already deleted!", typeName, id)
// 			return nil
// 		}
// 	}
// 	return err
// }

func diffWalk(ctx p.Context, diff map[string]p.PropertyDiff, path string, old reflect.Value, new reflect.Value) {
	ctx.Log(diag.Debug, fmt.Sprintf("diffWalk: visiting %s: old: %s new: %s", path, old.String(), new.String()))
	// Indirect through pointers and interfaces
	for old.Kind() == reflect.Ptr || old.Kind() == reflect.Interface {
		old = old.Elem()
	}
	for new.Kind() == reflect.Ptr || new.Kind() == reflect.Interface {
		new = new.Elem()
	}
	if new.Kind() != old.Kind() {
		ctx.Log(diag.Info, fmt.Sprintf("diffWalk: visiting %s: Kind changed: old: %s new: %s", path, old.Kind().String(), new.Kind().String()))
		diff[path] = p.PropertyDiff{Kind: p.Update}
		return
	}
	switch old.Kind() {
	case reflect.Array, reflect.Slice:
		mv := min(old.Len(), new.Len())
		if old.Len() != new.Len() {
			diff[fmt.Sprintf("%s[%d]", path, mv+1)] = p.PropertyDiff{Kind: p.Update}
		}
		for i := 0; i < mv; i++ {
			diffWalk(ctx, diff, fmt.Sprintf("%s[%d]", path, i), old.Index(i), new.Index(i))
		}
	case reflect.Map:
		if old.Len() != new.Len() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
		for _, k := range old.MapKeys() {
			diffWalk(ctx, diff, fmt.Sprintf("%s.%s", path, k.String()), old.MapIndex(k), new.MapIndex(k))
		}
	case reflect.Struct:
		for b := 0; b < new.NumField(); b++ {
			v2 := new.Field(b)
			o2 := old.Field(b)
			fld := new.Type().Field(b)
			name := fld.Name
			if tag, ok := fld.Tag.Lookup("pulumi"); ok {
				tagList := strings.Split(tag, ",")
				name = tagList[0]
				if len(name) == 0 {
					name = fld.Name
				}
			}
			// ctx.Log(diag.Info, fmt.Sprintf("comparing struct %s.%s (%d) [%s]<%s>: %s<>%s, %s != %s", path, name, b, fld.Name, old.Type().Field(b).Name, o2.Kind().String(), v2.Kind().String(), o2.String(), v2.String()))
			diffWalk(ctx, diff, fmt.Sprintf("%s.%s", path, name), o2, v2)
		}
	case reflect.String:
		if old.String() != new.String() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	case reflect.Int:
		if old.Int() != new.Int() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	case reflect.Bool:
		if old.Bool() != new.Bool() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	case reflect.Float64:
		if old.Float() != new.Float() {
			diff[path] = p.PropertyDiff{Kind: p.Update}
		}
	default:
		// handle other types
		diff[path] = p.PropertyDiff{Kind: p.Update}
		ctx.Log(diag.Warning, fmt.Sprintf("Unhandled types comparing %s: %s<>%s, %s != %s", path, old.Kind().String(), new.Kind().String(), old.String(), new.String()))
	}
}

func diffStrArrayIgnoreOrder(ctx p.Context, diff map[string]p.PropertyDiff, path string, old []string, new []string) {
	oldLen := len(old)
	newLen := len(new)
	minLen := min(oldLen, newLen)
	if oldLen > newLen {
		//noop
		diff[fmt.Sprintf("%s[%d]", path, minLen+1)] = p.PropertyDiff{Kind: p.Add}
	}
	if oldLen < newLen {
		diff[fmt.Sprintf("%s[%d]", path, minLen+1)] = p.PropertyDiff{Kind: p.Delete}
	}
	sort.Strings(old)
	sort.Strings(new)
	for i := 0; i < minLen; i++ {
		if old[i] != new[i] {
			diff[fmt.Sprintf("%s[%d]", path, i)] = p.PropertyDiff{Kind: p.Update}
		}
	}
}

func buildNameFilter(name string) *string {
	filter := "name=\"" + url.QueryEscape(name) + "\""
	return &filter
}

func ifte[T interface{}](cond bool, trueVal T, falseVal T) T {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}

func iftfe[T interface{}](cond bool, trueFunc func() T, falseVal T) T {
	if cond {
		return trueFunc()
	} else {
		return falseVal
	}
}

func ifted[T interface{}](cond bool, trueVal T, falseVal T) *T {
	if cond {
		return &trueVal
	} else {
		return &falseVal
	}
}
func iftden[T interface{}](cond bool, trueVal T) *T {
	if cond {
		return &trueVal
	} else {
		return nil
	}
}

func iftfden[T interface{}](cond bool, trueFunc func() T) *T {
	if cond {
		ret := trueFunc()
		return &ret
	} else {
		return nil
	}
}
