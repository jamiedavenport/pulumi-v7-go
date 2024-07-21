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
    'PropertyInput',
]

@pulumi.output_type
class PropertyInput(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "propertyId":
            suggest = "property_id"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in PropertyInput. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        PropertyInput.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        PropertyInput.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 property_id: str):
        pulumi.set(__self__, "property_id", property_id)

    @property
    @pulumi.getter(name="propertyId")
    def property_id(self) -> str:
        return pulumi.get(self, "property_id")


