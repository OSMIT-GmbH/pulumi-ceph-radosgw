package main

import (
	radosgw "github.com/OSMIT-GmbH/pulumi-ceph-radosgw/sdk/go/ceph-radosgw"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		myRandomResource, err := radosgw.NewRandom(ctx, "myRandomResource", &radosgw.RandomArgs{
			Length: pulumi.Int(24),
		})
		if err != nil {
			return err
		}
		ctx.Export("output", map[string]interface{}{
			"value": myRandomResource.Result,
		})
		return nil
	})
}
