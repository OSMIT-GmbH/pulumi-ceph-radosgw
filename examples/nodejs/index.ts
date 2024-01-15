import * as pulumi from "@pulumi/pulumi";
import * as cephRadosgw from "@osmit-gmbh/ceph-radosgw";

const myRandomResource = new cephRadosgw.Random("myRandomResource", {length: 24});
export const output = {
    value: myRandomResource.result,
};
