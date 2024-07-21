# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'PropertyInputArgs',
]

@pulumi.input_type
class PropertyInputArgs:
    def __init__(__self__, *,
                 property_id: pulumi.Input[str]):
        pulumi.set(__self__, "property_id", property_id)

    @property
    @pulumi.getter(name="propertyId")
    def property_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "property_id")

    @property_id.setter
    def property_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "property_id", value)

