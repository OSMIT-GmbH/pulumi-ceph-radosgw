# Pulumi Provider For Ceph radosgw

This provider supports provisioning of users, subusers, keys and S3 buckets on ceph rados gw.
Currently, it only supports some of the APIs we need for our purpose.
Feel free to extend it ;)

**WARNING:** This is in a quite an early stage - things my change ;)

## Usage

You have to set following config options:

```bash
AccessKeyID       string `pulumi:"accessKeyID"`
	SecretAccessKey   string `pulumi:"secretAccessKey" provider:"secret"`
	Endpoint          string `pulumi:"endpoint"`
	Insecure          string `pulumi:"insecure,optional"`
	insecure          bool

pulumi config set ceph-radosgw:endpoint https://rook-ceph-rgw-rook-s3.rados-ceph.svc:80
# you could use the credentals of the r/ceph 'cosi' user
pulumi config set ceph-radosgw:accessKeyID <access-key of user>
pulumi config set ceph-radosgw:secretAccessKey <secret-key of user>
# 'insecure' disables certificate checks when communicating with the rados-gw
# pulumi config set ceph-radosgw:insecure "true"
# `assimilate` allows to integrate existing objects. Great for migration
# pulumi config set ceph-radosgw:assimilate "true"
```

<!--
For samples how to create the ceph-radosgw objects have a look on the example [index.ts](examples/simple/index.ts)
-->

## Developing

This provider is based on the [pulumi-provider-boilerplate](https://github.com/pulumi/pulumi-provider-boilerplate).
Just use the README.md documentation from the boilerplate to set up your development environment.

This boilerplate creates a working Pulumi-owned provider named `ceph-radosgw`.
It implements a random number generator that you can [build and test out for yourself](#test-against-the-example) and
then replace the Random code with code specific to your provider.
