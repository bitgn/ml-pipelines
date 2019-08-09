# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: events.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='events.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0c\x65vents.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"2\n\x0eProjectCreated\x12\x12\n\nproject_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"f\n\x0e\x44\x61tasetCreated\x12\x12\n\ndataset_id\x18\x01 \x01(\t\x12\x12\n\nproject_id\x18\x02 \x01(\t\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x1e\n\x04meta\x18\x04 \x01(\x0b\x32\x10.DatasetMetadata\"X\n\x0e\x44\x61tasetUpdated\x12\x12\n\ndataset_id\x18\x01 \x01(\t\x12\x12\n\nproject_id\x18\x02 \x01(\t\x12\x1e\n\x04meta\x18\x03 \x01(\x0b\x32\x10.DatasetMetadata\"a\n\x08JobAdded\x12\x0e\n\x06job_id\x18\x01 \x01(\t\x12\x10\n\x08job_name\x18\x02 \x01(\t\x12\x0e\n\x06inputs\x18\x03 \x03(\t\x12\x0f\n\x07outputs\x18\x04 \x03(\t\x12\x12\n\nproject_id\x18\x05 \x01(\t\"5\n\x0b\x45xpertAdded\x12\x11\n\texpert_id\x18\x01 \x01(\t\x12\x13\n\x0b\x65xpert_name\x18\x02 \x01(\t\"k\n\rDatasetSample\x12%\n\x06\x66ormat\x18\x01 \x01(\x0e\x32\x15.DatasetSample.FORMAT\x12\x0c\n\x04\x62ody\x18\x02 \x01(\x0c\"%\n\x06\x46ORMAT\x12\x08\n\x04TEXT\x10\x00\x12\x07\n\x03TSV\x10\x01\x12\x08\n\x04JSON\x10\x02\"\xcf\x04\n\x0f\x44\x61tasetMetadata\x12\x14\n\x0crecord_count\x18\x01 \x01(\x03\x12\"\n\x12record_count_state\x18\x02 \x01(\x0e\x32\x06.STATE\x12\x12\n\nfile_count\x18\x03 \x01(\x03\x12 \n\x10\x66ile_count_state\x18\x04 \x01(\x0e\x32\x06.STATE\x12\x15\n\rstorage_bytes\x18\x05 \x01(\x03\x12#\n\x13storage_bytes_state\x18\x06 \x01(\x0e\x32\x06.STATE\x12\x1e\n\x06sample\x18\x07 \x01(\x0b\x32\x0e.DatasetSample\x12\x1c\n\x0csample_state\x18\x08 \x01(\x0e\x32\x06.STATE\x12\x18\n\x10update_timestamp\x18\t \x01(\x03\x12&\n\x16update_timestamp_state\x18\n \x01(\x0e\x32\x06.STATE\x12\x13\n\x0b\x64\x61ta_format\x18\x0b \x01(\t\x12!\n\x11\x64\x61ta_format_state\x18\x0c \x01(\x0e\x32\x06.STATE\x12\x13\n\x0b\x64\x65scription\x18\r \x01(\t\x12!\n\x11\x64\x65scription_state\x18\x0e \x01(\x0e\x32\x06.STATE\x12\x13\n\x0blocation_id\x18\x0f \x01(\t\x12!\n\x11location_id_state\x18\x10 \x01(\x0e\x32\x06.STATE\x12\x14\n\x0clocation_uri\x18\x11 \x01(\t\x12\"\n\x12location_uri_state\x18\x12 \x01(\x0e\x32\x06.STATE\x12\x0f\n\x07\x65xperts\x18\x13 \x03(\t\x12\x1d\n\rexperts_state\x18\x14 \x01(\x0e\x32\x06.STATE\"*\n\x05\x45vent\x12\x0c\n\x04\x42ody\x18\x01 \x01(\x0c\x12\x13\n\x04Type\x18\x02 \x01(\x0e\x32\x05.Type*(\n\x05STATE\x12\x08\n\x04NONE\x10\x00\x12\t\n\x05VALUE\x10\x01\x12\n\n\x06\x44\x45LETE\x10\x02*\x89\x01\n\x04Type\x12\x08\n\x04None\x10\x00\x12\x18\n\x14\x45vent_ProjectCreated\x10\x01\x12\x18\n\x14\x45vent_DatasetCreated\x10\x02\x12\x18\n\x14\x45vent_DatasetUpdated\x10\x03\x12\x15\n\x11\x45vent_ExpertAdded\x10\x04\x12\x12\n\x0e\x45vent_JobAdded\x10\x05\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_STATE = _descriptor.EnumDescriptor(
  name='STATE',
  full_name='STATE',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='VALUE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DELETE', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1196,
  serialized_end=1236,
)
_sym_db.RegisterEnumDescriptor(_STATE)

STATE = enum_type_wrapper.EnumTypeWrapper(_STATE)
_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='None', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_ProjectCreated', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_DatasetCreated', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_DatasetUpdated', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_ExpertAdded', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_JobAdded', index=5, number=5,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1239,
  serialized_end=1376,
)
_sym_db.RegisterEnumDescriptor(_TYPE)

Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
NONE = 0
VALUE = 1
DELETE = 2
globals()['None'] = 0
Event_ProjectCreated = 1
Event_DatasetCreated = 2
Event_DatasetUpdated = 3
Event_ExpertAdded = 4
Event_JobAdded = 5


_DATASETSAMPLE_FORMAT = _descriptor.EnumDescriptor(
  name='FORMAT',
  full_name='DatasetSample.FORMAT',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TEXT', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TSV', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JSON', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=519,
  serialized_end=556,
)
_sym_db.RegisterEnumDescriptor(_DATASETSAMPLE_FORMAT)


_PROJECTCREATED = _descriptor.Descriptor(
  name='ProjectCreated',
  full_name='ProjectCreated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_id', full_name='ProjectCreated.project_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='ProjectCreated.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=49,
  serialized_end=99,
)


_DATASETCREATED = _descriptor.Descriptor(
  name='DatasetCreated',
  full_name='DatasetCreated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='DatasetCreated.dataset_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetCreated.project_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='DatasetCreated.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='DatasetCreated.meta', index=3,
      number=4, type=11, cpp_type=10, label=1,
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
  serialized_start=101,
  serialized_end=203,
)


_DATASETUPDATED = _descriptor.Descriptor(
  name='DatasetUpdated',
  full_name='DatasetUpdated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='DatasetUpdated.dataset_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetUpdated.project_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='DatasetUpdated.meta', index=2,
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
  serialized_start=205,
  serialized_end=293,
)


_JOBADDED = _descriptor.Descriptor(
  name='JobAdded',
  full_name='JobAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='job_id', full_name='JobAdded.job_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='job_name', full_name='JobAdded.job_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs', full_name='JobAdded.inputs', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='JobAdded.outputs', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='JobAdded.project_id', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=295,
  serialized_end=392,
)


_EXPERTADDED = _descriptor.Descriptor(
  name='ExpertAdded',
  full_name='ExpertAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='expert_id', full_name='ExpertAdded.expert_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='expert_name', full_name='ExpertAdded.expert_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=394,
  serialized_end=447,
)


_DATASETSAMPLE = _descriptor.Descriptor(
  name='DatasetSample',
  full_name='DatasetSample',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='format', full_name='DatasetSample.format', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body', full_name='DatasetSample.body', index=1,
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
    _DATASETSAMPLE_FORMAT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=449,
  serialized_end=556,
)


_DATASETMETADATA = _descriptor.Descriptor(
  name='DatasetMetadata',
  full_name='DatasetMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='record_count', full_name='DatasetMetadata.record_count', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record_count_state', full_name='DatasetMetadata.record_count_state', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count', full_name='DatasetMetadata.file_count', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count_state', full_name='DatasetMetadata.file_count_state', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_bytes', full_name='DatasetMetadata.storage_bytes', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_bytes_state', full_name='DatasetMetadata.storage_bytes_state', index=5,
      number=6, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample', full_name='DatasetMetadata.sample', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_state', full_name='DatasetMetadata.sample_state', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp', full_name='DatasetMetadata.update_timestamp', index=8,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp_state', full_name='DatasetMetadata.update_timestamp_state', index=9,
      number=10, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format', full_name='DatasetMetadata.data_format', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format_state', full_name='DatasetMetadata.data_format_state', index=11,
      number=12, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='DatasetMetadata.description', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description_state', full_name='DatasetMetadata.description_state', index=13,
      number=14, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_id', full_name='DatasetMetadata.location_id', index=14,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_id_state', full_name='DatasetMetadata.location_id_state', index=15,
      number=16, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_uri', full_name='DatasetMetadata.location_uri', index=16,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_uri_state', full_name='DatasetMetadata.location_uri_state', index=17,
      number=18, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts', full_name='DatasetMetadata.experts', index=18,
      number=19, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts_state', full_name='DatasetMetadata.experts_state', index=19,
      number=20, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=559,
  serialized_end=1150,
)


_EVENT = _descriptor.Descriptor(
  name='Event',
  full_name='Event',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='Body', full_name='Event.Body', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='Type', full_name='Event.Type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1152,
  serialized_end=1194,
)

_DATASETCREATED.fields_by_name['meta'].message_type = _DATASETMETADATA
_DATASETUPDATED.fields_by_name['meta'].message_type = _DATASETMETADATA
_DATASETSAMPLE.fields_by_name['format'].enum_type = _DATASETSAMPLE_FORMAT
_DATASETSAMPLE_FORMAT.containing_type = _DATASETSAMPLE
_DATASETMETADATA.fields_by_name['record_count_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['file_count_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['storage_bytes_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['sample'].message_type = _DATASETSAMPLE
_DATASETMETADATA.fields_by_name['sample_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['update_timestamp_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['data_format_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['description_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['location_id_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['location_uri_state'].enum_type = _STATE
_DATASETMETADATA.fields_by_name['experts_state'].enum_type = _STATE
_EVENT.fields_by_name['Type'].enum_type = _TYPE
DESCRIPTOR.message_types_by_name['ProjectCreated'] = _PROJECTCREATED
DESCRIPTOR.message_types_by_name['DatasetCreated'] = _DATASETCREATED
DESCRIPTOR.message_types_by_name['DatasetUpdated'] = _DATASETUPDATED
DESCRIPTOR.message_types_by_name['JobAdded'] = _JOBADDED
DESCRIPTOR.message_types_by_name['ExpertAdded'] = _EXPERTADDED
DESCRIPTOR.message_types_by_name['DatasetSample'] = _DATASETSAMPLE
DESCRIPTOR.message_types_by_name['DatasetMetadata'] = _DATASETMETADATA
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
DESCRIPTOR.enum_types_by_name['STATE'] = _STATE
DESCRIPTOR.enum_types_by_name['Type'] = _TYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProjectCreated = _reflection.GeneratedProtocolMessageType('ProjectCreated', (_message.Message,), {
  'DESCRIPTOR' : _PROJECTCREATED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:ProjectCreated)
  })
_sym_db.RegisterMessage(ProjectCreated)

DatasetCreated = _reflection.GeneratedProtocolMessageType('DatasetCreated', (_message.Message,), {
  'DESCRIPTOR' : _DATASETCREATED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetCreated)
  })
_sym_db.RegisterMessage(DatasetCreated)

DatasetUpdated = _reflection.GeneratedProtocolMessageType('DatasetUpdated', (_message.Message,), {
  'DESCRIPTOR' : _DATASETUPDATED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetUpdated)
  })
_sym_db.RegisterMessage(DatasetUpdated)

JobAdded = _reflection.GeneratedProtocolMessageType('JobAdded', (_message.Message,), {
  'DESCRIPTOR' : _JOBADDED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:JobAdded)
  })
_sym_db.RegisterMessage(JobAdded)

ExpertAdded = _reflection.GeneratedProtocolMessageType('ExpertAdded', (_message.Message,), {
  'DESCRIPTOR' : _EXPERTADDED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:ExpertAdded)
  })
_sym_db.RegisterMessage(ExpertAdded)

DatasetSample = _reflection.GeneratedProtocolMessageType('DatasetSample', (_message.Message,), {
  'DESCRIPTOR' : _DATASETSAMPLE,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetSample)
  })
_sym_db.RegisterMessage(DatasetSample)

DatasetMetadata = _reflection.GeneratedProtocolMessageType('DatasetMetadata', (_message.Message,), {
  'DESCRIPTOR' : _DATASETMETADATA,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetMetadata)
  })
_sym_db.RegisterMessage(DatasetMetadata)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), {
  'DESCRIPTOR' : _EVENT,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:Event)
  })
_sym_db.RegisterMessage(Event)


# @@protoc_insertion_point(module_scope)