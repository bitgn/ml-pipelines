# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
    EnumDescriptor as google___protobuf___descriptor___EnumDescriptor,
)

from google.protobuf.message import (
    Message as google___protobuf___message___Message,
)

from typing import (
    List as typing___List,
    Optional as typing___Optional,
    Text as typing___Text,
    Tuple as typing___Tuple,
    cast as typing___cast,
)

from typing_extensions import (
    Literal as typing_extensions___Literal,
)


class DatasetMetadataDelta(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    sample_set = ... # type: bool
    description = ... # type: typing___Text
    description_set = ... # type: bool
    summary = ... # type: typing___Text
    summary_set = ... # type: bool

    @property
    def sample(self) -> Sample: ...

    def __init__(self,
        *,
        sample : typing___Optional[Sample] = None,
        sample_set : typing___Optional[bool] = None,
        description : typing___Optional[typing___Text] = None,
        description_set : typing___Optional[bool] = None,
        summary : typing___Optional[typing___Text] = None,
        summary_set : typing___Optional[bool] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> DatasetMetadataDelta: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"sample"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"description",u"description_set",u"sample",u"sample_set",u"summary",u"summary_set"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"sample",b"sample"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"description",b"description",u"description_set",b"description_set",u"sample",b"sample",u"sample_set",b"sample_set",u"summary",b"summary",u"summary_set",b"summary_set"]) -> None: ...

class Sample(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    class FORMAT(int):
        DESCRIPTOR: google___protobuf___descriptor___EnumDescriptor = ...
        @classmethod
        def Name(cls, number: int) -> str: ...
        @classmethod
        def Value(cls, name: str) -> Sample.FORMAT: ...
        @classmethod
        def keys(cls) -> typing___List[str]: ...
        @classmethod
        def values(cls) -> typing___List[Sample.FORMAT]: ...
        @classmethod
        def items(cls) -> typing___List[typing___Tuple[str, Sample.FORMAT]]: ...
        NONE = typing___cast(Sample.FORMAT, 0)
        TEXT = typing___cast(Sample.FORMAT, 1)
        TSV = typing___cast(Sample.FORMAT, 2)
        JSON = typing___cast(Sample.FORMAT, 3)
    NONE = typing___cast(Sample.FORMAT, 0)
    TEXT = typing___cast(Sample.FORMAT, 1)
    TSV = typing___cast(Sample.FORMAT, 2)
    JSON = typing___cast(Sample.FORMAT, 3)

    format = ... # type: Sample.FORMAT
    body = ... # type: typing___Text

    def __init__(self,
        *,
        format : typing___Optional[Sample.FORMAT] = None,
        body : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> Sample: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"body",u"format"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"body",b"body",u"format",b"format"]) -> None: ...
