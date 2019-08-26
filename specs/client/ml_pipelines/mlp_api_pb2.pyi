# @generated by generate_proto_mypy_stubs.py.  Do not edit!
import sys
from google.protobuf.descriptor import (
    Descriptor as google___protobuf___descriptor___Descriptor,
    EnumDescriptor as google___protobuf___descriptor___EnumDescriptor,
)

from google.protobuf.internal.containers import (
    RepeatedScalarFieldContainer as google___protobuf___internal___containers___RepeatedScalarFieldContainer,
)

from google.protobuf.message import (
    Message as google___protobuf___message___Message,
)

from typing import (
    Iterable as typing___Iterable,
    List as typing___List,
    Optional as typing___Optional,
    Text as typing___Text,
    Tuple as typing___Tuple,
    cast as typing___cast,
)

from typing_extensions import (
    Literal as typing_extensions___Literal,
)

from vo_pb2 import (
    DatasetMetadataDelta as vo_pb2___DatasetMetadataDelta,
    ProjectMetadataDelta as vo_pb2___ProjectMetadataDelta,
)


class StatusCode(int):
    DESCRIPTOR: google___protobuf___descriptor___EnumDescriptor = ...
    @classmethod
    def Name(cls, number: int) -> str: ...
    @classmethod
    def Value(cls, name: str) -> StatusCode: ...
    @classmethod
    def keys(cls) -> typing___List[str]: ...
    @classmethod
    def values(cls) -> typing___List[StatusCode]: ...
    @classmethod
    def items(cls) -> typing___List[typing___Tuple[str, StatusCode]]: ...
    OK = typing___cast(StatusCode, 0)
    CANCELLED = typing___cast(StatusCode, 1)
    UNKNOWN = typing___cast(StatusCode, 2)
    INVALID_ARGUMENT = typing___cast(StatusCode, 3)
    DEADLINE_EXCEEDED = typing___cast(StatusCode, 4)
    NOT_FOUND = typing___cast(StatusCode, 5)
    ALREADY_EXISTS = typing___cast(StatusCode, 6)
    PERMISSION_DENIED = typing___cast(StatusCode, 7)
    UNAUTHENTICATED = typing___cast(StatusCode, 16)
    RESOURCE_EXHAUSTED = typing___cast(StatusCode, 8)
    FAILED_PRECONDITION = typing___cast(StatusCode, 9)
    ABORTED = typing___cast(StatusCode, 10)
    OUT_OF_RANGE = typing___cast(StatusCode, 11)
    UNIMPLEMENTED = typing___cast(StatusCode, 12)
    INTERNAL = typing___cast(StatusCode, 13)
    UNAVAILABLE = typing___cast(StatusCode, 14)
    DATA_LOSS = typing___cast(StatusCode, 15)
    BAD_NAME = typing___cast(StatusCode, 100)
    NAME_TAKEN = typing___cast(StatusCode, 101)
OK = typing___cast(StatusCode, 0)
CANCELLED = typing___cast(StatusCode, 1)
UNKNOWN = typing___cast(StatusCode, 2)
INVALID_ARGUMENT = typing___cast(StatusCode, 3)
DEADLINE_EXCEEDED = typing___cast(StatusCode, 4)
NOT_FOUND = typing___cast(StatusCode, 5)
ALREADY_EXISTS = typing___cast(StatusCode, 6)
PERMISSION_DENIED = typing___cast(StatusCode, 7)
UNAUTHENTICATED = typing___cast(StatusCode, 16)
RESOURCE_EXHAUSTED = typing___cast(StatusCode, 8)
FAILED_PRECONDITION = typing___cast(StatusCode, 9)
ABORTED = typing___cast(StatusCode, 10)
OUT_OF_RANGE = typing___cast(StatusCode, 11)
UNIMPLEMENTED = typing___cast(StatusCode, 12)
INTERNAL = typing___cast(StatusCode, 13)
UNAVAILABLE = typing___cast(StatusCode, 14)
DATA_LOSS = typing___cast(StatusCode, 15)
BAD_NAME = typing___cast(StatusCode, 100)
NAME_TAKEN = typing___cast(StatusCode, 101)

class CreateProjectRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    name = ... # type: typing___Text

    @property
    def meta(self) -> vo_pb2___ProjectMetadataDelta: ...

    def __init__(self,
        *,
        name : typing___Optional[typing___Text] = None,
        meta : typing___Optional[vo_pb2___ProjectMetadataDelta] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> CreateProjectRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"meta"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"meta",u"name"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"meta",b"meta"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"meta",b"meta",u"name",b"name"]) -> None: ...

class CreateProjectResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    uid = ... # type: bytes

    @property
    def error(self) -> ApiError: ...

    def __init__(self,
        *,
        error : typing___Optional[ApiError] = None,
        uid : typing___Optional[bytes] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> CreateProjectResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"error",u"uid"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"error",b"error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"error",b"error",u"uid",b"uid"]) -> None: ...

class CreateDatasetResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    uid = ... # type: bytes

    @property
    def error(self) -> ApiError: ...

    def __init__(self,
        *,
        error : typing___Optional[ApiError] = None,
        uid : typing___Optional[bytes] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> CreateDatasetResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"error",u"uid"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"error",b"error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"error",b"error",u"uid",b"uid"]) -> None: ...

class StatRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    def __init__(self,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> StatRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...

class StatResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    version = ... # type: typing___Text

    def __init__(self,
        *,
        version : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> StatResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"version"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"version",b"version"]) -> None: ...

class ApiResponse(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...

    @property
    def error(self) -> ApiError: ...

    def __init__(self,
        *,
        error : typing___Optional[ApiError] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> ApiResponse: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"error"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"error",b"error"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"error",b"error"]) -> None: ...

class CreateDatasetRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    project_uid = ... # type: bytes
    name = ... # type: typing___Text

    @property
    def meta(self) -> vo_pb2___DatasetMetadataDelta: ...

    def __init__(self,
        *,
        project_uid : typing___Optional[bytes] = None,
        name : typing___Optional[typing___Text] = None,
        meta : typing___Optional[vo_pb2___DatasetMetadataDelta] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> CreateDatasetRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"meta"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"meta",u"name",u"project_uid"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"meta",b"meta"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"meta",b"meta",u"name",b"name",u"project_uid",b"project_uid"]) -> None: ...

class UpdateDatasetRequest(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    uid = ... # type: bytes

    @property
    def meta(self) -> vo_pb2___DatasetMetadataDelta: ...

    def __init__(self,
        *,
        uid : typing___Optional[bytes] = None,
        meta : typing___Optional[vo_pb2___DatasetMetadataDelta] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> UpdateDatasetRequest: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"meta"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"meta",u"uid"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"meta",b"meta"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"meta",b"meta",u"uid",b"uid"]) -> None: ...

class ApiError(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    code = ... # type: StatusCode
    message = ... # type: typing___Text
    subject_uid = ... # type: bytes
    subject_name = ... # type: typing___Text
    field_name = ... # type: typing___Text
    details = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text]

    def __init__(self,
        *,
        code : typing___Optional[StatusCode] = None,
        message : typing___Optional[typing___Text] = None,
        subject_uid : typing___Optional[bytes] = None,
        subject_name : typing___Optional[typing___Text] = None,
        field_name : typing___Optional[typing___Text] = None,
        details : typing___Optional[typing___Iterable[typing___Text]] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> ApiError: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"code",u"details",u"field_name",u"message",u"subject_name",u"subject_uid"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"code",b"code",u"details",b"details",u"field_name",b"field_name",u"message",b"message",u"subject_name",b"subject_name",u"subject_uid",b"subject_uid"]) -> None: ...
