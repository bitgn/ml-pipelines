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


class FIELD_TYPES(int):
    DESCRIPTOR: google___protobuf___descriptor___EnumDescriptor = ...
    @classmethod
    def Name(cls, number: int) -> str: ...
    @classmethod
    def Value(cls, name: str) -> FIELD_TYPES: ...
    @classmethod
    def keys(cls) -> typing___List[str]: ...
    @classmethod
    def values(cls) -> typing___List[FIELD_TYPES]: ...
    @classmethod
    def items(cls) -> typing___List[typing___Tuple[str, FIELD_TYPES]]: ...
    FIELD_NONE = typing___cast(FIELD_TYPES, 0)
    FIELD_RECORD_COUNT = typing___cast(FIELD_TYPES, 1)
    FIELD_RAW_BYTES = typing___cast(FIELD_TYPES, 2)
    FIELD_ZIP_BYTES = typing___cast(FIELD_TYPES, 3)
    FIELD_SAMPLE = typing___cast(FIELD_TYPES, 4)
    FIELD_UPDATE_TIMESTAMP = typing___cast(FIELD_TYPES, 5)
    FIELD_FILE_COUNT = typing___cast(FIELD_TYPES, 6)
    FIELD_DATA_FORMAT = typing___cast(FIELD_TYPES, 7)
    FIELD_DESCRIPTION = typing___cast(FIELD_TYPES, 8)
FIELD_NONE = typing___cast(FIELD_TYPES, 0)
FIELD_RECORD_COUNT = typing___cast(FIELD_TYPES, 1)
FIELD_RAW_BYTES = typing___cast(FIELD_TYPES, 2)
FIELD_ZIP_BYTES = typing___cast(FIELD_TYPES, 3)
FIELD_SAMPLE = typing___cast(FIELD_TYPES, 4)
FIELD_UPDATE_TIMESTAMP = typing___cast(FIELD_TYPES, 5)
FIELD_FILE_COUNT = typing___cast(FIELD_TYPES, 6)
FIELD_DATA_FORMAT = typing___cast(FIELD_TYPES, 7)
FIELD_DESCRIPTION = typing___cast(FIELD_TYPES, 8)

class SAMPLE_KIND(int):
    DESCRIPTOR: google___protobuf___descriptor___EnumDescriptor = ...
    @classmethod
    def Name(cls, number: int) -> str: ...
    @classmethod
    def Value(cls, name: str) -> SAMPLE_KIND: ...
    @classmethod
    def keys(cls) -> typing___List[str]: ...
    @classmethod
    def values(cls) -> typing___List[SAMPLE_KIND]: ...
    @classmethod
    def items(cls) -> typing___List[typing___Tuple[str, SAMPLE_KIND]]: ...
    TEXT = typing___cast(SAMPLE_KIND, 0)
    TSV = typing___cast(SAMPLE_KIND, 1)
    JSON = typing___cast(SAMPLE_KIND, 2)
TEXT = typing___cast(SAMPLE_KIND, 0)
TSV = typing___cast(SAMPLE_KIND, 1)
JSON = typing___cast(SAMPLE_KIND, 2)

class ProjectCreated(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    project_id = ... # type: typing___Text
    name = ... # type: typing___Text

    def __init__(self,
        *,
        project_id : typing___Optional[typing___Text] = None,
        name : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> ProjectCreated: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"name",u"project_id"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"name",b"name",u"project_id",b"project_id"]) -> None: ...

class DatasetCreated(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    dataset_id = ... # type: typing___Text
    name = ... # type: typing___Text
    project_id = ... # type: typing___Text

    @property
    def metadata(self) -> DatasetMetadata: ...

    def __init__(self,
        *,
        dataset_id : typing___Optional[typing___Text] = None,
        name : typing___Optional[typing___Text] = None,
        project_id : typing___Optional[typing___Text] = None,
        metadata : typing___Optional[DatasetMetadata] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> DatasetCreated: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"metadata"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"dataset_id",u"metadata",u"name",u"project_id"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"metadata",b"metadata"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"dataset_id",b"dataset_id",u"metadata",b"metadata",u"name",b"name",u"project_id",b"project_id"]) -> None: ...

class DatasetUpdated(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    dataset_id = ... # type: typing___Text
    project_id = ... # type: typing___Text

    @property
    def metadata(self) -> DatasetMetadata: ...

    def __init__(self,
        *,
        dataset_id : typing___Optional[typing___Text] = None,
        project_id : typing___Optional[typing___Text] = None,
        metadata : typing___Optional[DatasetMetadata] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> DatasetUpdated: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def HasField(self, field_name: typing_extensions___Literal[u"metadata"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"dataset_id",u"metadata",u"project_id"]) -> None: ...
    else:
        def HasField(self, field_name: typing_extensions___Literal[u"metadata",b"metadata"]) -> bool: ...
        def ClearField(self, field_name: typing_extensions___Literal[u"dataset_id",b"dataset_id",u"metadata",b"metadata",u"project_id",b"project_id"]) -> None: ...

class JobAdded(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    job_id = ... # type: typing___Text
    job_name = ... # type: typing___Text
    inputs = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text]
    outputs = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[typing___Text]
    project_id = ... # type: typing___Text

    def __init__(self,
        *,
        job_id : typing___Optional[typing___Text] = None,
        job_name : typing___Optional[typing___Text] = None,
        inputs : typing___Optional[typing___Iterable[typing___Text]] = None,
        outputs : typing___Optional[typing___Iterable[typing___Text]] = None,
        project_id : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> JobAdded: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"inputs",u"job_id",u"job_name",u"outputs",u"project_id"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"inputs",b"inputs",u"job_id",b"job_id",u"job_name",b"job_name",u"outputs",b"outputs",u"project_id",b"project_id"]) -> None: ...

class DatasetMetadata(google___protobuf___message___Message):
    DESCRIPTOR: google___protobuf___descriptor___Descriptor = ...
    set_fields = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[FIELD_TYPES]
    del_fields = ... # type: google___protobuf___internal___containers___RepeatedScalarFieldContainer[FIELD_TYPES]
    record_count = ... # type: int
    file_count = ... # type: int
    raw_bytes = ... # type: int
    zip_bytes = ... # type: int
    sample_body = ... # type: bytes
    sample_kind = ... # type: SAMPLE_KIND
    update_timestamp = ... # type: int
    data_format = ... # type: typing___Text
    description = ... # type: typing___Text

    def __init__(self,
        *,
        set_fields : typing___Optional[typing___Iterable[FIELD_TYPES]] = None,
        del_fields : typing___Optional[typing___Iterable[FIELD_TYPES]] = None,
        record_count : typing___Optional[int] = None,
        file_count : typing___Optional[int] = None,
        raw_bytes : typing___Optional[int] = None,
        zip_bytes : typing___Optional[int] = None,
        sample_body : typing___Optional[bytes] = None,
        sample_kind : typing___Optional[SAMPLE_KIND] = None,
        update_timestamp : typing___Optional[int] = None,
        data_format : typing___Optional[typing___Text] = None,
        description : typing___Optional[typing___Text] = None,
        ) -> None: ...
    @classmethod
    def FromString(cls, s: bytes) -> DatasetMetadata: ...
    def MergeFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    def CopyFrom(self, other_msg: google___protobuf___message___Message) -> None: ...
    if sys.version_info >= (3,):
        def ClearField(self, field_name: typing_extensions___Literal[u"data_format",u"del_fields",u"description",u"file_count",u"raw_bytes",u"record_count",u"sample_body",u"sample_kind",u"set_fields",u"update_timestamp",u"zip_bytes"]) -> None: ...
    else:
        def ClearField(self, field_name: typing_extensions___Literal[u"data_format",b"data_format",u"del_fields",b"del_fields",u"description",b"description",u"file_count",b"file_count",u"raw_bytes",b"raw_bytes",u"record_count",b"record_count",u"sample_body",b"sample_body",u"sample_kind",b"sample_kind",u"set_fields",b"set_fields",u"update_timestamp",b"update_timestamp",u"zip_bytes",b"zip_bytes"]) -> None: ...
