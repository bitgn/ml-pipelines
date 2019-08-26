# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: mlp_api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from . import vo_pb2 as vo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='mlp_api.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\rmlp_api.proto\x1a\x08vo.proto\"I\n\x14\x43reateProjectRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12#\n\x04meta\x18\x02 \x01(\x0b\x32\x15.ProjectMetadataDelta\">\n\x15\x43reateProjectResponse\x12\x18\n\x05\x65rror\x18\x01 \x01(\x0b\x32\t.ApiError\x12\x0b\n\x03uid\x18\x02 \x01(\x0c\">\n\x15\x43reateDatasetResponse\x12\x18\n\x05\x65rror\x18\x01 \x01(\x0b\x32\t.ApiError\x12\x0b\n\x03uid\x18\x02 \x01(\x0c\"\r\n\x0bStatRequest\"\x1f\n\x0cStatResponse\x12\x0f\n\x07version\x18\x01 \x01(\t\"V\n\x10\x43reateJobRequest\x12\x13\n\x0bproject_uid\x18\x01 \x01(\x0c\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x1f\n\x04meta\x18\x03 \x01(\x0b\x32\x11.JobMetadataDelta\":\n\x11\x43reateJobResponse\x12\x18\n\x05\x65rror\x18\x01 \x01(\x0b\x32\t.ApiError\x12\x0b\n\x03uid\x18\x02 \x01(\x0c\"\'\n\x0b\x41piResponse\x12\x18\n\x05\x65rror\x18\x01 \x01(\x0b\x32\t.ApiError\"^\n\x14\x43reateDatasetRequest\x12\x13\n\x0bproject_uid\x18\x01 \x01(\x0c\x12\x0c\n\x04name\x18\x02 \x01(\t\x12#\n\x04meta\x18\x03 \x01(\x0b\x32\x15.DatasetMetadataDelta\"H\n\x14UpdateDatasetRequest\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12#\n\x04meta\x18\x02 \x01(\x0b\x32\x15.DatasetMetadataDelta\"\x86\x01\n\x08\x41piError\x12\x19\n\x04\x63ode\x18\x01 \x01(\x0e\x32\x0b.StatusCode\x12\x0f\n\x07message\x18\x02 \x01(\t\x12\x13\n\x0bsubject_uid\x18\x03 \x01(\x0c\x12\x14\n\x0csubject_name\x18\x04 \x01(\t\x12\x12\n\nfield_name\x18\x05 \x01(\t\x12\x0f\n\x07\x64\x65tails\x18\x06 \x03(\t*\xdb\x02\n\nStatusCode\x12\x06\n\x02OK\x10\x00\x12\r\n\tCANCELLED\x10\x01\x12\x0b\n\x07UNKNOWN\x10\x02\x12\x14\n\x10INVALID_ARGUMENT\x10\x03\x12\x15\n\x11\x44\x45\x41\x44LINE_EXCEEDED\x10\x04\x12\r\n\tNOT_FOUND\x10\x05\x12\x12\n\x0e\x41LREADY_EXISTS\x10\x06\x12\x15\n\x11PERMISSION_DENIED\x10\x07\x12\x13\n\x0fUNAUTHENTICATED\x10\x10\x12\x16\n\x12RESOURCE_EXHAUSTED\x10\x08\x12\x17\n\x13\x46\x41ILED_PRECONDITION\x10\t\x12\x0b\n\x07\x41\x42ORTED\x10\n\x12\x10\n\x0cOUT_OF_RANGE\x10\x0b\x12\x11\n\rUNIMPLEMENTED\x10\x0c\x12\x0c\n\x08INTERNAL\x10\r\x12\x0f\n\x0bUNAVAILABLE\x10\x0e\x12\r\n\tDATA_LOSS\x10\x0f\x12\x0c\n\x08\x42\x41\x44_NAME\x10\x64\x12\x0e\n\nNAME_TAKEN\x10\x65\x32\xa2\x02\n\x07\x43\x61talog\x12@\n\rCreateProject\x12\x15.CreateProjectRequest\x1a\x16.CreateProjectResponse\"\x00\x12@\n\rCreateDataset\x12\x15.CreateDatasetRequest\x1a\x16.CreateDatasetResponse\"\x00\x12\x36\n\rUpdateDataset\x12\x15.UpdateDatasetRequest\x1a\x0c.ApiResponse\"\x00\x12\x34\n\tCreateJob\x12\x11.CreateJobRequest\x1a\x12.CreateJobResponse\"\x00\x12%\n\x04Stat\x12\x0c.StatRequest\x1a\r.StatResponse\"\x00\x62\x06proto3')
  ,
  dependencies=[vo__pb2.DESCRIPTOR,])

_STATUSCODE = _descriptor.EnumDescriptor(
  name='StatusCode',
  full_name='StatusCode',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='OK', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CANCELLED', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INVALID_ARGUMENT', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEADLINE_EXCEEDED', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NOT_FOUND', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ALREADY_EXISTS', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PERMISSION_DENIED', index=7, number=7,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNAUTHENTICATED', index=8, number=16,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RESOURCE_EXHAUSTED', index=9, number=8,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FAILED_PRECONDITION', index=10, number=9,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ABORTED', index=11, number=10,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='OUT_OF_RANGE', index=12, number=11,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNIMPLEMENTED', index=13, number=12,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='INTERNAL', index=14, number=13,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='UNAVAILABLE', index=15, number=14,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DATA_LOSS', index=16, number=15,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BAD_NAME', index=17, number=100,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NAME_TAKEN', index=18, number=101,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=775,
  serialized_end=1122,
)
_sym_db.RegisterEnumDescriptor(_STATUSCODE)

StatusCode = enum_type_wrapper.EnumTypeWrapper(_STATUSCODE)
OK = 0
CANCELLED = 1
UNKNOWN = 2
INVALID_ARGUMENT = 3
DEADLINE_EXCEEDED = 4
NOT_FOUND = 5
ALREADY_EXISTS = 6
PERMISSION_DENIED = 7
UNAUTHENTICATED = 16
RESOURCE_EXHAUSTED = 8
FAILED_PRECONDITION = 9
ABORTED = 10
OUT_OF_RANGE = 11
UNIMPLEMENTED = 12
INTERNAL = 13
UNAVAILABLE = 14
DATA_LOSS = 15
BAD_NAME = 100
NAME_TAKEN = 101



_CREATEPROJECTREQUEST = _descriptor.Descriptor(
  name='CreateProjectRequest',
  full_name='CreateProjectRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='CreateProjectRequest.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='CreateProjectRequest.meta', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=27,
  serialized_end=100,
)


_CREATEPROJECTRESPONSE = _descriptor.Descriptor(
  name='CreateProjectResponse',
  full_name='CreateProjectResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='CreateProjectResponse.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uid', full_name='CreateProjectResponse.uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=102,
  serialized_end=164,
)


_CREATEDATASETRESPONSE = _descriptor.Descriptor(
  name='CreateDatasetResponse',
  full_name='CreateDatasetResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='CreateDatasetResponse.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uid', full_name='CreateDatasetResponse.uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=166,
  serialized_end=228,
)


_STATREQUEST = _descriptor.Descriptor(
  name='StatRequest',
  full_name='StatRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=230,
  serialized_end=243,
)


_STATRESPONSE = _descriptor.Descriptor(
  name='StatResponse',
  full_name='StatResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='version', full_name='StatResponse.version', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=245,
  serialized_end=276,
)


_CREATEJOBREQUEST = _descriptor.Descriptor(
  name='CreateJobRequest',
  full_name='CreateJobRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='CreateJobRequest.project_uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='CreateJobRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='CreateJobRequest.meta', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=278,
  serialized_end=364,
)


_CREATEJOBRESPONSE = _descriptor.Descriptor(
  name='CreateJobResponse',
  full_name='CreateJobResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='CreateJobResponse.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uid', full_name='CreateJobResponse.uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=366,
  serialized_end=424,
)


_APIRESPONSE = _descriptor.Descriptor(
  name='ApiResponse',
  full_name='ApiResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='error', full_name='ApiResponse.error', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=426,
  serialized_end=465,
)


_CREATEDATASETREQUEST = _descriptor.Descriptor(
  name='CreateDatasetRequest',
  full_name='CreateDatasetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='CreateDatasetRequest.project_uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='CreateDatasetRequest.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='CreateDatasetRequest.meta', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=467,
  serialized_end=561,
)


_UPDATEDATASETREQUEST = _descriptor.Descriptor(
  name='UpdateDatasetRequest',
  full_name='UpdateDatasetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='UpdateDatasetRequest.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='UpdateDatasetRequest.meta', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=563,
  serialized_end=635,
)


_APIERROR = _descriptor.Descriptor(
  name='ApiError',
  full_name='ApiError',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='code', full_name='ApiError.code', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='message', full_name='ApiError.message', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subject_uid', full_name='ApiError.subject_uid', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subject_name', full_name='ApiError.subject_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='field_name', full_name='ApiError.field_name', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='details', full_name='ApiError.details', index=5,
      number=6, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=638,
  serialized_end=772,
)

_CREATEPROJECTREQUEST.fields_by_name['meta'].message_type = vo__pb2._PROJECTMETADATADELTA
_CREATEPROJECTRESPONSE.fields_by_name['error'].message_type = _APIERROR
_CREATEDATASETRESPONSE.fields_by_name['error'].message_type = _APIERROR
_CREATEJOBREQUEST.fields_by_name['meta'].message_type = vo__pb2._JOBMETADATADELTA
_CREATEJOBRESPONSE.fields_by_name['error'].message_type = _APIERROR
_APIRESPONSE.fields_by_name['error'].message_type = _APIERROR
_CREATEDATASETREQUEST.fields_by_name['meta'].message_type = vo__pb2._DATASETMETADATADELTA
_UPDATEDATASETREQUEST.fields_by_name['meta'].message_type = vo__pb2._DATASETMETADATADELTA
_APIERROR.fields_by_name['code'].enum_type = _STATUSCODE
DESCRIPTOR.message_types_by_name['CreateProjectRequest'] = _CREATEPROJECTREQUEST
DESCRIPTOR.message_types_by_name['CreateProjectResponse'] = _CREATEPROJECTRESPONSE
DESCRIPTOR.message_types_by_name['CreateDatasetResponse'] = _CREATEDATASETRESPONSE
DESCRIPTOR.message_types_by_name['StatRequest'] = _STATREQUEST
DESCRIPTOR.message_types_by_name['StatResponse'] = _STATRESPONSE
DESCRIPTOR.message_types_by_name['CreateJobRequest'] = _CREATEJOBREQUEST
DESCRIPTOR.message_types_by_name['CreateJobResponse'] = _CREATEJOBRESPONSE
DESCRIPTOR.message_types_by_name['ApiResponse'] = _APIRESPONSE
DESCRIPTOR.message_types_by_name['CreateDatasetRequest'] = _CREATEDATASETREQUEST
DESCRIPTOR.message_types_by_name['UpdateDatasetRequest'] = _UPDATEDATASETREQUEST
DESCRIPTOR.message_types_by_name['ApiError'] = _APIERROR
DESCRIPTOR.enum_types_by_name['StatusCode'] = _STATUSCODE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateProjectRequest = _reflection.GeneratedProtocolMessageType('CreateProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTREQUEST,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:CreateProjectRequest)
  })
_sym_db.RegisterMessage(CreateProjectRequest)

CreateProjectResponse = _reflection.GeneratedProtocolMessageType('CreateProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTRESPONSE,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:CreateProjectResponse)
  })
_sym_db.RegisterMessage(CreateProjectResponse)

CreateDatasetResponse = _reflection.GeneratedProtocolMessageType('CreateDatasetResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDATASETRESPONSE,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:CreateDatasetResponse)
  })
_sym_db.RegisterMessage(CreateDatasetResponse)

StatRequest = _reflection.GeneratedProtocolMessageType('StatRequest', (_message.Message,), {
  'DESCRIPTOR' : _STATREQUEST,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:StatRequest)
  })
_sym_db.RegisterMessage(StatRequest)

StatResponse = _reflection.GeneratedProtocolMessageType('StatResponse', (_message.Message,), {
  'DESCRIPTOR' : _STATRESPONSE,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:StatResponse)
  })
_sym_db.RegisterMessage(StatResponse)

CreateJobRequest = _reflection.GeneratedProtocolMessageType('CreateJobRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEJOBREQUEST,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:CreateJobRequest)
  })
_sym_db.RegisterMessage(CreateJobRequest)

CreateJobResponse = _reflection.GeneratedProtocolMessageType('CreateJobResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEJOBRESPONSE,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:CreateJobResponse)
  })
_sym_db.RegisterMessage(CreateJobResponse)

ApiResponse = _reflection.GeneratedProtocolMessageType('ApiResponse', (_message.Message,), {
  'DESCRIPTOR' : _APIRESPONSE,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:ApiResponse)
  })
_sym_db.RegisterMessage(ApiResponse)

CreateDatasetRequest = _reflection.GeneratedProtocolMessageType('CreateDatasetRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDATASETREQUEST,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:CreateDatasetRequest)
  })
_sym_db.RegisterMessage(CreateDatasetRequest)

UpdateDatasetRequest = _reflection.GeneratedProtocolMessageType('UpdateDatasetRequest', (_message.Message,), {
  'DESCRIPTOR' : _UPDATEDATASETREQUEST,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:UpdateDatasetRequest)
  })
_sym_db.RegisterMessage(UpdateDatasetRequest)

ApiError = _reflection.GeneratedProtocolMessageType('ApiError', (_message.Message,), {
  'DESCRIPTOR' : _APIERROR,
  '__module__' : 'mlp_api_pb2'
  # @@protoc_insertion_point(class_scope:ApiError)
  })
_sym_db.RegisterMessage(ApiError)



_CATALOG = _descriptor.ServiceDescriptor(
  name='Catalog',
  full_name='Catalog',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1125,
  serialized_end=1415,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateProject',
    full_name='Catalog.CreateProject',
    index=0,
    containing_service=None,
    input_type=_CREATEPROJECTREQUEST,
    output_type=_CREATEPROJECTRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CreateDataset',
    full_name='Catalog.CreateDataset',
    index=1,
    containing_service=None,
    input_type=_CREATEDATASETREQUEST,
    output_type=_CREATEDATASETRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='UpdateDataset',
    full_name='Catalog.UpdateDataset',
    index=2,
    containing_service=None,
    input_type=_UPDATEDATASETREQUEST,
    output_type=_APIRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='CreateJob',
    full_name='Catalog.CreateJob',
    index=3,
    containing_service=None,
    input_type=_CREATEJOBREQUEST,
    output_type=_CREATEJOBRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Stat',
    full_name='Catalog.Stat',
    index=4,
    containing_service=None,
    input_type=_STATREQUEST,
    output_type=_STATRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_CATALOG)

DESCRIPTOR.services_by_name['Catalog'] = _CATALOG

# @@protoc_insertion_point(module_scope)
