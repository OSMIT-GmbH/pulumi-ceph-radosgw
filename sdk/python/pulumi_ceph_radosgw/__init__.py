# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .provider import *
from .user import *

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_ceph_radosgw.config as __config
    config = __config
else:
    config = _utilities.lazy_import('pulumi_ceph_radosgw.config')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "ceph-radosgw",
  "mod": "index",
  "fqn": "pulumi_ceph_radosgw",
  "classes": {
   "ceph-radosgw:index:User": "User"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "ceph-radosgw",
  "token": "pulumi:providers:ceph-radosgw",
  "fqn": "pulumi_ceph_radosgw",
  "class": "Provider"
 }
]
"""
)