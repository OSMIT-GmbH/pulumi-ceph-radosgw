# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BucketArgs', 'Bucket']

@pulumi.input_type
class BucketArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 object_locking: Optional[pulumi.Input[bool]] = None,
                 purge_on_delete: Optional[pulumi.Input[bool]] = None,
                 quota: Optional[pulumi.Input['QuotaArgsArgs']] = None,
                 versioning: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Bucket resource.
        :param pulumi.Input[str] name: Bucket name
        :param pulumi.Input[bool] object_locking: Bucket object locking enabled
        :param pulumi.Input[bool] purge_on_delete: Purge bucket on delete
        :param pulumi.Input['QuotaArgsArgs'] quota: Bucket quota configuration
        """
        pulumi.set(__self__, "name", name)
        if object_locking is not None:
            pulumi.set(__self__, "object_locking", object_locking)
        if purge_on_delete is not None:
            pulumi.set(__self__, "purge_on_delete", purge_on_delete)
        if quota is not None:
            pulumi.set(__self__, "quota", quota)
        if versioning is not None:
            pulumi.set(__self__, "versioning", versioning)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Bucket name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="objectLocking")
    def object_locking(self) -> Optional[pulumi.Input[bool]]:
        """
        Bucket object locking enabled
        """
        return pulumi.get(self, "object_locking")

    @object_locking.setter
    def object_locking(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "object_locking", value)

    @property
    @pulumi.getter(name="purgeOnDelete")
    def purge_on_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Purge bucket on delete
        """
        return pulumi.get(self, "purge_on_delete")

    @purge_on_delete.setter
    def purge_on_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "purge_on_delete", value)

    @property
    @pulumi.getter
    def quota(self) -> Optional[pulumi.Input['QuotaArgsArgs']]:
        """
        Bucket quota configuration
        """
        return pulumi.get(self, "quota")

    @quota.setter
    def quota(self, value: Optional[pulumi.Input['QuotaArgsArgs']]):
        pulumi.set(self, "quota", value)

    @property
    @pulumi.getter
    def versioning(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "versioning")

    @versioning.setter
    def versioning(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "versioning", value)


class Bucket(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_locking: Optional[pulumi.Input[bool]] = None,
                 purge_on_delete: Optional[pulumi.Input[bool]] = None,
                 quota: Optional[pulumi.Input[pulumi.InputType['QuotaArgsArgs']]] = None,
                 versioning: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a Bucket resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Bucket name
        :param pulumi.Input[bool] object_locking: Bucket object locking enabled
        :param pulumi.Input[bool] purge_on_delete: Purge bucket on delete
        :param pulumi.Input[pulumi.InputType['QuotaArgsArgs']] quota: Bucket quota configuration
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BucketArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Bucket resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param BucketArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BucketArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 object_locking: Optional[pulumi.Input[bool]] = None,
                 purge_on_delete: Optional[pulumi.Input[bool]] = None,
                 quota: Optional[pulumi.Input[pulumi.InputType['QuotaArgsArgs']]] = None,
                 versioning: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BucketArgs.__new__(BucketArgs)

            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["object_locking"] = object_locking
            __props__.__dict__["purge_on_delete"] = purge_on_delete
            __props__.__dict__["quota"] = quota
            __props__.__dict__["versioning"] = versioning
            __props__.__dict__["_assimilated"] = None
            __props__.__dict__["_location"] = None
            __props__.__dict__["ubid"] = None
        super(Bucket, __self__).__init__(
            'ceph-radosgw:index:Bucket',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Bucket':
        """
        Get an existing Bucket resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = BucketArgs.__new__(BucketArgs)

        __props__.__dict__["_assimilated"] = None
        __props__.__dict__["_location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["object_locking"] = None
        __props__.__dict__["purge_on_delete"] = None
        __props__.__dict__["quota"] = None
        __props__.__dict__["ubid"] = None
        __props__.__dict__["versioning"] = None
        return Bucket(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def _assimilated(self) -> pulumi.Output[bool]:
        """
        Bucket was 'assimilated' - managing an existing bucket.
        """
        return pulumi.get(self, "_assimilated")

    @property
    @pulumi.getter
    def _location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "_location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Bucket name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="objectLocking")
    def object_locking(self) -> pulumi.Output[Optional[bool]]:
        """
        Bucket object locking enabled
        """
        return pulumi.get(self, "object_locking")

    @property
    @pulumi.getter(name="purgeOnDelete")
    def purge_on_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Purge bucket on delete.
        """
        return pulumi.get(self, "purge_on_delete")

    @property
    @pulumi.getter
    def quota(self) -> pulumi.Output[Optional['outputs.QuotaArgs']]:
        """
        Bucket quota configuration
        """
        return pulumi.get(self, "quota")

    @property
    @pulumi.getter
    def ubid(self) -> pulumi.Output[str]:
        """
        The unique bucket id.
        """
        return pulumi.get(self, "ubid")

    @property
    @pulumi.getter
    def versioning(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "versioning")

