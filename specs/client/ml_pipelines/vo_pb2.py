# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: vo.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='vo.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x08vo.proto\"\x97\x01\n\x14\x44\x61tasetMetadataDelta\x12\x17\n\x06sample\x18\x01 \x01(\x0b\x32\x07.Sample\x12\x12\n\nsample_set\x18\x02 \x01(\x08\x12\x13\n\x0b\x64\x65scription\x18\x05 \x01(\t\x12\x17\n\x0f\x64\x65scription_set\x18\x06 \x01(\x08\x12\x0f\n\x07summary\x18\x07 \x01(\t\x12\x13\n\x0bsummary_set\x18\x08 \x01(\x08\"g\n\x06Sample\x12\x1e\n\x06\x66ormat\x18\x01 \x01(\x0e\x32\x0e.Sample.FORMAT\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t\"/\n\x06\x46ORMAT\x12\x08\n\x04NONE\x10\x00\x12\x08\n\x04TEXT\x10\x01\x12\x07\n\x03TSV\x10\x02\x12\x08\n\x04JSON\x10\x03*c\n\x0e\x41\x43TIVITY_LEVEL\x12\x14\n\x10\x41\x43TIVITY_VERBOSE\x10\x00\x12\x11\n\rACTIVITY_INFO\x10\x01\x12\x12\n\x0e\x41\x43TIVITY_ERROR\x10\x02\x12\x14\n\x10\x41\x43TIVITY_SUCCESS\x10\x03*G\n\x0e\x44\x41TASET_STATUS\x12\x0f\n\x0bSTATUS_NONE\x10\x00\x12\x10\n\x0cSTATUS_ERROR\x10\x01\x12\x12\n\x0eSTATUS_SUCCESS\x10\x02\x62\x06proto3')
)

_ACTIVITY_LEVEL = _descriptor.EnumDescriptor(
  name='ACTIVITY_LEVEL',
  full_name='ACTIVITY_LEVEL',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='ACTIVITY_VERBOSE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ACTIVITY_INFO', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ACTIVITY_ERROR', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ACTIVITY_SUCCESS', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=271,
  serialized_end=370,
)
_sym_db.RegisterEnumDescriptor(_ACTIVITY_LEVEL)

ACTIVITY_LEVEL = enum_type_wrapper.EnumTypeWrapper(_ACTIVITY_LEVEL)
_DATASET_STATUS = _descriptor.EnumDescriptor(
  name='DATASET_STATUS',
  full_name='DATASET_STATUS',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='STATUS_NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STATUS_ERROR', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='STATUS_SUCCESS', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=372,
  serialized_end=443,
)
_sym_db.RegisterEnumDescriptor(_DATASET_STATUS)

DATASET_STATUS = enum_type_wrapper.EnumTypeWrapper(_DATASET_STATUS)
ACTIVITY_VERBOSE = 0
ACTIVITY_INFO = 1
ACTIVITY_ERROR = 2
ACTIVITY_SUCCESS = 3
STATUS_NONE = 0
STATUS_ERROR = 1
STATUS_SUCCESS = 2


_SAMPLE_FORMAT = _descriptor.EnumDescriptor(
  name='FORMAT',
  full_name='Sample.FORMAT',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TEXT', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TSV', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JSON', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=222,
  serialized_end=269,
)
_sym_db.RegisterEnumDescriptor(_SAMPLE_FORMAT)


_DATASETMETADATADELTA = _descriptor.Descriptor(
  name='DatasetMetadataDelta',
  full_name='DatasetMetadataDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sample', full_name='DatasetMetadataDelta.sample', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_set', full_name='DatasetMetadataDelta.sample_set', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='DatasetMetadataDelta.description', index=2,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description_set', full_name='DatasetMetadataDelta.description_set', index=3,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='summary', full_name='DatasetMetadataDelta.summary', index=4,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='summary_set', full_name='DatasetMetadataDelta.summary_set', index=5,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=13,
  serialized_end=164,
)


_SAMPLE = _descriptor.Descriptor(
  name='Sample',
  full_name='Sample',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='format', full_name='Sample.format', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='body', full_name='Sample.body', index=1,
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
    _SAMPLE_FORMAT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=166,
  serialized_end=269,
)

_DATASETMETADATADELTA.fields_by_name['sample'].message_type = _SAMPLE
_SAMPLE.fields_by_name['format'].enum_type = _SAMPLE_FORMAT
_SAMPLE_FORMAT.containing_type = _SAMPLE
DESCRIPTOR.message_types_by_name['DatasetMetadataDelta'] = _DATASETMETADATADELTA
DESCRIPTOR.message_types_by_name['Sample'] = _SAMPLE
DESCRIPTOR.enum_types_by_name['ACTIVITY_LEVEL'] = _ACTIVITY_LEVEL
DESCRIPTOR.enum_types_by_name['DATASET_STATUS'] = _DATASET_STATUS
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

DatasetMetadataDelta = _reflection.GeneratedProtocolMessageType('DatasetMetadataDelta', (_message.Message,), {
  'DESCRIPTOR' : _DATASETMETADATADELTA,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:DatasetMetadataDelta)
  })
_sym_db.RegisterMessage(DatasetMetadataDelta)

Sample = _reflection.GeneratedProtocolMessageType('Sample', (_message.Message,), {
  'DESCRIPTOR' : _SAMPLE,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:Sample)
  })
_sym_db.RegisterMessage(Sample)


# @@protoc_insertion_point(module_scope)
