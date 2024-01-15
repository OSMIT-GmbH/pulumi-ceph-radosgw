# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['KeyArgs', 'Key']

@pulumi.input_type
class KeyArgs:
    def __init__(__self__, *,
                 user_id: pulumi.Input[str],
                 access_key: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sub_user_name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Key resource.
        :param pulumi.Input[str] user_id: User-ID of 'parent' user
        :param pulumi.Input[str] access_key: The access key
        :param pulumi.Input[str] key_type: key type - either 's3' or 'swift'
        :param pulumi.Input[str] secret_key: The secret key
        :param pulumi.Input[str] sub_user_name: Name of sub-user (optional)
        """
        pulumi.set(__self__, "user_id", user_id)
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sub_user_name is not None:
            pulumi.set(__self__, "sub_user_name", sub_user_name)

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
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The access key
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        key type - either 's3' or 'swift'
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The secret key
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="subUserName")
    def sub_user_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of sub-user (optional)
        """
        return pulumi.get(self, "sub_user_name")

    @sub_user_name.setter
    def sub_user_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sub_user_name", value)


class Key(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sub_user_name: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Key resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The access key
        :param pulumi.Input[str] key_type: key type - either 's3' or 'swift'
        :param pulumi.Input[str] secret_key: The secret key
        :param pulumi.Input[str] sub_user_name: Name of sub-user (optional)
        :param pulumi.Input[str] user_id: User-ID of 'parent' user
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: KeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Key resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param KeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(KeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
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
            __props__ = KeyArgs.__new__(KeyArgs)

            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["key_type"] = key_type
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["sub_user_name"] = sub_user_name
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
            __props__.__dict__["_assimilated"] = None
        super(Key, __self__).__init__(
            'ceph-radosgw:index:Key',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Key':
        """
        Get an existing Key resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = KeyArgs.__new__(KeyArgs)

        __props__.__dict__["_assimilated"] = None
        __props__.__dict__["access_key"] = None
        __props__.__dict__["key_type"] = None
        __props__.__dict__["secret_key"] = None
        __props__.__dict__["sub_user_name"] = None
        __props__.__dict__["user_id"] = None
        return Key(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def _assimilated(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "_assimilated")

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[Optional[str]]:
        """
        The access key
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[Optional[str]]:
        """
        key type - either 's3' or 'swift'
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[Optional[str]]:
        """
        The secret key
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="subUserName")
    def sub_user_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of sub-user (optional)
        """
        return pulumi.get(self, "sub_user_name")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        User-ID of 'parent' user
        """
        return pulumi.get(self, "user_id")

