# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 access_key_id: pulumi.Input[str],
                 endpoint: pulumi.Input[str],
                 secret_access_key: pulumi.Input[str],
                 assimilate: Optional[pulumi.Input[str]] = None,
                 delete_assimilated: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] access_key_id: The username. It's important but not secret.
        :param pulumi.Input[str] endpoint: The URI to the API
        :param pulumi.Input[str] secret_access_key: The password. It is very secret.
        :param pulumi.Input[str] assimilate: Assimilate an existing object during create
        :param pulumi.Input[str] delete_assimilated: Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
        :param pulumi.Input[str] insecure: Don't validate server SSL certificate
        """
        pulumi.set(__self__, "access_key_id", access_key_id)
        pulumi.set(__self__, "endpoint", endpoint)
        pulumi.set(__self__, "secret_access_key", secret_access_key)
        if assimilate is not None:
            pulumi.set(__self__, "assimilate", assimilate)
        if delete_assimilated is not None:
            pulumi.set(__self__, "delete_assimilated", delete_assimilated)
        if insecure is not None:
            pulumi.set(__self__, "insecure", insecure)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="accessKeyID")
    def access_key_id(self) -> pulumi.Input[str]:
        """
        The username. It's important but not secret.
        """
        return pulumi.get(self, "access_key_id")

    @access_key_id.setter
    def access_key_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "access_key_id", value)

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Input[str]:
        """
        The URI to the API
        """
        return pulumi.get(self, "endpoint")

    @endpoint.setter
    def endpoint(self, value: pulumi.Input[str]):
        pulumi.set(self, "endpoint", value)

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> pulumi.Input[str]:
        """
        The password. It is very secret.
        """
        return pulumi.get(self, "secret_access_key")

    @secret_access_key.setter
    def secret_access_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_access_key", value)

    @property
    @pulumi.getter
    def assimilate(self) -> Optional[pulumi.Input[str]]:
        """
        Assimilate an existing object during create
        """
        return pulumi.get(self, "assimilate")

    @assimilate.setter
    def assimilate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "assimilate", value)

    @property
    @pulumi.getter(name="deleteAssimilated")
    def delete_assimilated(self) -> Optional[pulumi.Input[str]]:
        """
        Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
        """
        return pulumi.get(self, "delete_assimilated")

    @delete_assimilated.setter
    def delete_assimilated(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delete_assimilated", value)

    @property
    @pulumi.getter
    def insecure(self) -> Optional[pulumi.Input[str]]:
        """
        Don't validate server SSL certificate
        """
        return pulumi.get(self, "insecure")

    @insecure.setter
    def insecure(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "insecure", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key_id: Optional[pulumi.Input[str]] = None,
                 assimilate: Optional[pulumi.Input[str]] = None,
                 delete_assimilated: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[str]] = None,
                 secret_access_key: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Ceph-radosgw resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key_id: The username. It's important but not secret.
        :param pulumi.Input[str] assimilate: Assimilate an existing object during create
        :param pulumi.Input[str] delete_assimilated: Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
        :param pulumi.Input[str] endpoint: The URI to the API
        :param pulumi.Input[str] insecure: Don't validate server SSL certificate
        :param pulumi.Input[str] secret_access_key: The password. It is very secret.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Ceph-radosgw resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key_id: Optional[pulumi.Input[str]] = None,
                 assimilate: Optional[pulumi.Input[str]] = None,
                 delete_assimilated: Optional[pulumi.Input[str]] = None,
                 endpoint: Optional[pulumi.Input[str]] = None,
                 insecure: Optional[pulumi.Input[str]] = None,
                 secret_access_key: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if access_key_id is None and not opts.urn:
                raise TypeError("Missing required property 'access_key_id'")
            __props__.__dict__["access_key_id"] = access_key_id
            __props__.__dict__["assimilate"] = assimilate
            __props__.__dict__["delete_assimilated"] = delete_assimilated
            if endpoint is None and not opts.urn:
                raise TypeError("Missing required property 'endpoint'")
            __props__.__dict__["endpoint"] = endpoint
            __props__.__dict__["insecure"] = insecure
            if secret_access_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_access_key'")
            __props__.__dict__["secret_access_key"] = None if secret_access_key is None else pulumi.Output.secret(secret_access_key)
            __props__.__dict__["version"] = version
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretAccessKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Provider, __self__).__init__(
            'ceph-radosgw',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="accessKeyID")
    def access_key_id(self) -> pulumi.Output[str]:
        """
        The username. It's important but not secret.
        """
        return pulumi.get(self, "access_key_id")

    @property
    @pulumi.getter
    def assimilate(self) -> pulumi.Output[Optional[str]]:
        """
        Assimilate an existing object during create
        """
        return pulumi.get(self, "assimilate")

    @property
    @pulumi.getter(name="deleteAssimilated")
    def delete_assimilated(self) -> pulumi.Output[Optional[str]]:
        """
        Delete assimilated objects during delete (otherwise they would be kept on OpenZiti)
        """
        return pulumi.get(self, "delete_assimilated")

    @property
    @pulumi.getter
    def endpoint(self) -> pulumi.Output[str]:
        """
        The URI to the API
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def insecure(self) -> pulumi.Output[Optional[str]]:
        """
        Don't validate server SSL certificate
        """
        return pulumi.get(self, "insecure")

    @property
    @pulumi.getter(name="secretAccessKey")
    def secret_access_key(self) -> pulumi.Output[str]:
        """
        The password. It is very secret.
        """
        return pulumi.get(self, "secret_access_key")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "version")

