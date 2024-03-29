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
from ._enums import *

__all__ = ['SubUserArgs', 'SubUser']

@pulumi.input_type
class SubUserArgs:
    def __init__(__self__, *,
                 permissions: pulumi.Input['SubUserPermission'],
                 sub_user_name: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 generate_key: Optional[pulumi.Input[bool]] = None,
                 key_type: Optional[pulumi.Input['KeyType']] = None,
                 purge_keys: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SubUser resource.
        :param pulumi.Input[str] sub_user_name: Name of this subuser
        :param pulumi.Input[str] user_id: User-ID of 'parent' user
        """
        pulumi.set(__self__, "permissions", permissions)
        pulumi.set(__self__, "sub_user_name", sub_user_name)
        pulumi.set(__self__, "user_id", user_id)
        if generate_key is not None:
            pulumi.set(__self__, "generate_key", generate_key)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if purge_keys is not None:
            pulumi.set(__self__, "purge_keys", purge_keys)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Input['SubUserPermission']:
        return pulumi.get(self, "permissions")

    @permissions.setter
    def permissions(self, value: pulumi.Input['SubUserPermission']):
        pulumi.set(self, "permissions", value)

    @property
    @pulumi.getter(name="subUserName")
    def sub_user_name(self) -> pulumi.Input[str]:
        """
        Name of this subuser
        """
        return pulumi.get(self, "sub_user_name")

    @sub_user_name.setter
    def sub_user_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "sub_user_name", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        """
        User-ID of 'parent' user
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter(name="generateKey")
    def generate_key(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "generate_key")

    @generate_key.setter
    def generate_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "generate_key", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input['KeyType']]:
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input['KeyType']]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter(name="purgeKeys")
    def purge_keys(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "purge_keys")

    @purge_keys.setter
    def purge_keys(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "purge_keys", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)


class SubUser(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 generate_key: Optional[pulumi.Input[bool]] = None,
                 key_type: Optional[pulumi.Input['KeyType']] = None,
                 permissions: Optional[pulumi.Input['SubUserPermission']] = None,
                 purge_keys: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sub_user_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a SubUser resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] sub_user_name: Name of this subuser
        :param pulumi.Input[str] user_id: User-ID of 'parent' user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SubUserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SubUser resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SubUserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SubUserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 generate_key: Optional[pulumi.Input[bool]] = None,
                 key_type: Optional[pulumi.Input['KeyType']] = None,
                 permissions: Optional[pulumi.Input['SubUserPermission']] = None,
                 purge_keys: Optional[pulumi.Input[bool]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sub_user_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SubUserArgs.__new__(SubUserArgs)

            __props__.__dict__["generate_key"] = generate_key
            __props__.__dict__["key_type"] = key_type
            if permissions is None and not opts.urn:
                raise TypeError("Missing required property 'permissions'")
            __props__.__dict__["permissions"] = permissions
            __props__.__dict__["purge_keys"] = purge_keys
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            if sub_user_name is None and not opts.urn:
                raise TypeError("Missing required property 'sub_user_name'")
            __props__.__dict__["sub_user_name"] = sub_user_name
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["_assimilated"] = None
            __props__.__dict__["fullname"] = None
            __props__.__dict__["keys"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SubUser, __self__).__init__(
            'ceph-radosgw:index:SubUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'SubUser':
        """
        Get an existing SubUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = SubUserArgs.__new__(SubUserArgs)

        __props__.__dict__["_assimilated"] = None
        __props__.__dict__["fullname"] = None
        __props__.__dict__["generate_key"] = None
        __props__.__dict__["key_type"] = None
        __props__.__dict__["keys"] = None
        __props__.__dict__["permissions"] = None
        __props__.__dict__["purge_keys"] = None
        __props__.__dict__["secret_key"] = None
        __props__.__dict__["sub_user_name"] = None
        __props__.__dict__["user_id"] = None
        return SubUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def _assimilated(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "_assimilated")

    @property
    @pulumi.getter
    def fullname(self) -> pulumi.Output[str]:
        return pulumi.get(self, "fullname")

    @property
    @pulumi.getter(name="generateKey")
    def generate_key(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "generate_key")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[Optional['KeyType']]:
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Output[Sequence['outputs.KeyEntry']]:
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def permissions(self) -> pulumi.Output['SubUserPermission']:
        return pulumi.get(self, "permissions")

    @property
    @pulumi.getter(name="purgeKeys")
    def purge_keys(self) -> pulumi.Output[Optional[bool]]:
        return pulumi.get(self, "purge_keys")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="subUserName")
    def sub_user_name(self) -> pulumi.Output[str]:
        """
        Name of this subuser
        """
        return pulumi.get(self, "sub_user_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        User-ID of 'parent' user
        """
        return pulumi.get(self, "user_id")

