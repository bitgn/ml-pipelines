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
  serialized_pb=_b('\n\x08vo.proto\"f\n\x14ProjectMetadataDelta\x12\r\n\x05title\x18\x01 \x01(\t\x12\x11\n\ttitle_set\x18\x02 \x01(\x08\x12\x13\n\x0b\x64\x65scription\x18\x03 \x01(\t\x12\x17\n\x0f\x64\x65scription_set\x18\x04 \x01(\x08\"\xea\x01\n\x14ServiceMetadataDelta\x12\r\n\x05title\x18\x01 \x01(\t\x12\x11\n\ttitle_set\x18\x02 \x01(\x08\x12\x13\n\x0blocation_id\x18\x03 \x01(\t\x12\x17\n\x0flocation_id_set\x18\x04 \x01(\x08\x12\x14\n\x0clocation_uri\x18\x05 \x01(\t\x12\x18\n\x10location_uri_set\x18\x06 \x01(\x08\x12\x13\n\x0b\x64\x65scription\x18\x07 \x01(\t\x12\x17\n\x0f\x64\x65scription_set\x18\x08 \x01(\x08\x12\x0f\n\x07\x65xperts\x18\t \x03(\t\x12\x13\n\x0b\x65xperts_set\x18\n \x01(\x08\"a\n\x13\x45xpertMetadataDelta\x12\x11\n\tfull_name\x18\x01 \x01(\t\x12\x15\n\rfull_name_set\x18\x02 \x01(\x08\x12\r\n\x05\x65mail\x18\x03 \x01(\t\x12\x11\n\temail_set\x18\x04 \x01(\x08\";\n\x08Relation\x12\x11\n\ttarget_id\x18\x01 \x01(\x0c\x12\x1c\n\x0btarget_type\x18\x02 \x01(\x0e\x32\x07.ENTITY\"\x95\x01\n\x10JobMetadataDelta\x12\r\n\x05title\x18\x01 \x01(\t\x12\x11\n\ttitle_set\x18\x02 \x01(\x08\x12\x19\n\x06inputs\x18\x03 \x03(\x0b\x32\t.JobInput\x12\x12\n\ninputs_set\x18\x04 \x01(\x08\x12\x1b\n\x07outputs\x18\x05 \x03(\x0b\x32\n.JobOutput\x12\x13\n\x0boutputs_set\x18\x06 \x01(\x08\"Z\n\x08JobInput\x12\x11\n\tsource_id\x18\x01 \x01(\x0c\x12\x1c\n\x04type\x18\x02 \x01(\x0e\x32\x0e.JobInput.Type\"\x1d\n\x04Type\x12\x08\n\x04None\x10\x00\x12\x0b\n\x07\x44\x61taset\x10\x01\"\\\n\tJobOutput\x12\x11\n\ttarget_id\x18\x01 \x01(\x0c\x12\x1d\n\x04type\x18\x02 \x01(\x0e\x32\x0f.JobOutput.Type\"\x1d\n\x04Type\x12\x08\n\x04None\x10\x00\x12\x0b\n\x07\x44\x61taset\x10\x01\"b\n\x0f\x44\x61tasetVerInput\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12#\n\x04type\x18\x02 \x01(\x0e\x32\x15.DatasetVerInput.TYPE\"\x1d\n\x04TYPE\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07JOB_RUN\x10\x01\"\x92\x04\n\x14\x44\x61tasetMetadataDelta\x12\x14\n\x0crecord_count\x18\x01 \x01(\x03\x12\x18\n\x10record_count_set\x18\x02 \x01(\x08\x12\x12\n\nfile_count\x18\x03 \x01(\x03\x12\x16\n\x0e\x66ile_count_set\x18\x04 \x01(\x08\x12\x15\n\rstorage_bytes\x18\x05 \x01(\x03\x12\x19\n\x11storage_bytes_set\x18\x06 \x01(\x08\x12\x1e\n\x06sample\x18\x07 \x01(\x0b\x32\x0e.DatasetSample\x12\x12\n\nsample_set\x18\x08 \x01(\x08\x12\x18\n\x10update_timestamp\x18\t \x01(\x03\x12\x1c\n\x14update_timestamp_set\x18\n \x01(\x08\x12\x13\n\x0b\x64\x61ta_format\x18\x0b \x01(\t\x12\x17\n\x0f\x64\x61ta_format_set\x18\x0c \x01(\x08\x12\x13\n\x0b\x64\x65scription\x18\r \x01(\t\x12\x17\n\x0f\x64\x65scription_set\x18\x0e \x01(\x08\x12\x13\n\x0blocation_id\x18\x0f \x01(\t\x12\x17\n\x0flocation_id_set\x18\x10 \x01(\x08\x12\x14\n\x0clocation_uri\x18\x11 \x01(\t\x12\x18\n\x10location_uri_set\x18\x12 \x01(\x08\x12\x0f\n\x07\x65xperts\x18\x13 \x03(\x0c\x12\x13\n\x0b\x65xperts_set\x18\x14 \x01(\x08\x12\r\n\x05title\x18\x15 \x01(\t\x12\x11\n\ttitle_set\x18\x16 \x01(\x08\"P\n\x0b\x44\x61tasetItem\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x15\n\rstorage_bytes\x18\x02 \x01(\x03\x12\x0f\n\x07records\x18\x03 \x01(\x03\x12\x0b\n\x03uid\x18\x04 \x01(\x0c\"k\n\rDatasetSample\x12%\n\x06\x66ormat\x18\x01 \x01(\x0e\x32\x15.DatasetSample.FORMAT\x12\x0c\n\x04\x62ody\x18\x02 \x01(\t\"%\n\x06\x46ORMAT\x12\x08\n\x04TEXT\x10\x00\x12\x07\n\x03TSV\x10\x01\x12\x08\n\x04JSON\x10\x02*M\n\x06\x45NTITY\x12\x08\n\x04NONE\x10\x00\x12\x0b\n\x07\x44\x41TASET\x10\x01\x12\x07\n\x03JOB\x10\x02\x12\x0b\n\x07SERVICE\x10\x03\x12\t\n\x05MODEL\x10\x04\x12\x0b\n\x07PROJECT\x10\x05\x62\x06proto3')
)

_ENTITY = _descriptor.EnumDescriptor(
  name='ENTITY',
  full_name='ENTITY',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DATASET', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JOB', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='SERVICE', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MODEL', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='PROJECT', index=5, number=5,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1675,
  serialized_end=1752,
)
_sym_db.RegisterEnumDescriptor(_ENTITY)

ENTITY = enum_type_wrapper.EnumTypeWrapper(_ENTITY)
NONE = 0
DATASET = 1
JOB = 2
SERVICE = 3
MODEL = 4
PROJECT = 5


_JOBINPUT_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='JobInput.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='None', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Dataset', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=726,
  serialized_end=755,
)
_sym_db.RegisterEnumDescriptor(_JOBINPUT_TYPE)

_JOBOUTPUT_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='JobOutput.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='None', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Dataset', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=726,
  serialized_end=755,
)
_sym_db.RegisterEnumDescriptor(_JOBOUTPUT_TYPE)

_DATASETVERINPUT_TYPE = _descriptor.EnumDescriptor(
  name='TYPE',
  full_name='DatasetVerInput.TYPE',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JOB_RUN', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=920,
  serialized_end=949,
)
_sym_db.RegisterEnumDescriptor(_DATASETVERINPUT_TYPE)

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
  serialized_start=1636,
  serialized_end=1673,
)
_sym_db.RegisterEnumDescriptor(_DATASETSAMPLE_FORMAT)


_PROJECTMETADATADELTA = _descriptor.Descriptor(
  name='ProjectMetadataDelta',
  full_name='ProjectMetadataDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='title', full_name='ProjectMetadataDelta.title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title_set', full_name='ProjectMetadataDelta.title_set', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='ProjectMetadataDelta.description', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description_set', full_name='ProjectMetadataDelta.description_set', index=3,
      number=4, type=8, cpp_type=7, label=1,
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
  serialized_start=12,
  serialized_end=114,
)


_SERVICEMETADATADELTA = _descriptor.Descriptor(
  name='ServiceMetadataDelta',
  full_name='ServiceMetadataDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='title', full_name='ServiceMetadataDelta.title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title_set', full_name='ServiceMetadataDelta.title_set', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_id', full_name='ServiceMetadataDelta.location_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_id_set', full_name='ServiceMetadataDelta.location_id_set', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_uri', full_name='ServiceMetadataDelta.location_uri', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_uri_set', full_name='ServiceMetadataDelta.location_uri_set', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='ServiceMetadataDelta.description', index=6,
      number=7, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description_set', full_name='ServiceMetadataDelta.description_set', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts', full_name='ServiceMetadataDelta.experts', index=8,
      number=9, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts_set', full_name='ServiceMetadataDelta.experts_set', index=9,
      number=10, type=8, cpp_type=7, label=1,
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
  serialized_start=117,
  serialized_end=351,
)


_EXPERTMETADATADELTA = _descriptor.Descriptor(
  name='ExpertMetadataDelta',
  full_name='ExpertMetadataDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='full_name', full_name='ExpertMetadataDelta.full_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='full_name_set', full_name='ExpertMetadataDelta.full_name_set', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email', full_name='ExpertMetadataDelta.email', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='email_set', full_name='ExpertMetadataDelta.email_set', index=3,
      number=4, type=8, cpp_type=7, label=1,
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
  serialized_start=353,
  serialized_end=450,
)


_RELATION = _descriptor.Descriptor(
  name='Relation',
  full_name='Relation',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='target_id', full_name='Relation.target_id', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='target_type', full_name='Relation.target_type', index=1,
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
  serialized_start=452,
  serialized_end=511,
)


_JOBMETADATADELTA = _descriptor.Descriptor(
  name='JobMetadataDelta',
  full_name='JobMetadataDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='title', full_name='JobMetadataDelta.title', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title_set', full_name='JobMetadataDelta.title_set', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs', full_name='JobMetadataDelta.inputs', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs_set', full_name='JobMetadataDelta.inputs_set', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='JobMetadataDelta.outputs', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs_set', full_name='JobMetadataDelta.outputs_set', index=5,
      number=6, type=8, cpp_type=7, label=1,
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
  serialized_start=514,
  serialized_end=663,
)


_JOBINPUT = _descriptor.Descriptor(
  name='JobInput',
  full_name='JobInput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='source_id', full_name='JobInput.source_id', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='JobInput.type', index=1,
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
    _JOBINPUT_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=665,
  serialized_end=755,
)


_JOBOUTPUT = _descriptor.Descriptor(
  name='JobOutput',
  full_name='JobOutput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='target_id', full_name='JobOutput.target_id', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='JobOutput.type', index=1,
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
    _JOBOUTPUT_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=757,
  serialized_end=849,
)


_DATASETVERINPUT = _descriptor.Descriptor(
  name='DatasetVerInput',
  full_name='DatasetVerInput',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='DatasetVerInput.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='DatasetVerInput.type', index=1,
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
    _DATASETVERINPUT_TYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=851,
  serialized_end=949,
)


_DATASETMETADATADELTA = _descriptor.Descriptor(
  name='DatasetMetadataDelta',
  full_name='DatasetMetadataDelta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='record_count', full_name='DatasetMetadataDelta.record_count', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record_count_set', full_name='DatasetMetadataDelta.record_count_set', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count', full_name='DatasetMetadataDelta.file_count', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count_set', full_name='DatasetMetadataDelta.file_count_set', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_bytes', full_name='DatasetMetadataDelta.storage_bytes', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_bytes_set', full_name='DatasetMetadataDelta.storage_bytes_set', index=5,
      number=6, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample', full_name='DatasetMetadataDelta.sample', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_set', full_name='DatasetMetadataDelta.sample_set', index=7,
      number=8, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp', full_name='DatasetMetadataDelta.update_timestamp', index=8,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp_set', full_name='DatasetMetadataDelta.update_timestamp_set', index=9,
      number=10, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format', full_name='DatasetMetadataDelta.data_format', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format_set', full_name='DatasetMetadataDelta.data_format_set', index=11,
      number=12, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='DatasetMetadataDelta.description', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description_set', full_name='DatasetMetadataDelta.description_set', index=13,
      number=14, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_id', full_name='DatasetMetadataDelta.location_id', index=14,
      number=15, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_id_set', full_name='DatasetMetadataDelta.location_id_set', index=15,
      number=16, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_uri', full_name='DatasetMetadataDelta.location_uri', index=16,
      number=17, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='location_uri_set', full_name='DatasetMetadataDelta.location_uri_set', index=17,
      number=18, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts', full_name='DatasetMetadataDelta.experts', index=18,
      number=19, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='experts_set', full_name='DatasetMetadataDelta.experts_set', index=19,
      number=20, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='DatasetMetadataDelta.title', index=20,
      number=21, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title_set', full_name='DatasetMetadataDelta.title_set', index=21,
      number=22, type=8, cpp_type=7, label=1,
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
  serialized_start=952,
  serialized_end=1482,
)


_DATASETITEM = _descriptor.Descriptor(
  name='DatasetItem',
  full_name='DatasetItem',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='DatasetItem.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='storage_bytes', full_name='DatasetItem.storage_bytes', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='records', full_name='DatasetItem.records', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uid', full_name='DatasetItem.uid', index=3,
      number=4, type=12, cpp_type=9, label=1,
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
  serialized_start=1484,
  serialized_end=1564,
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
    _DATASETSAMPLE_FORMAT,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1566,
  serialized_end=1673,
)

_RELATION.fields_by_name['target_type'].enum_type = _ENTITY
_JOBMETADATADELTA.fields_by_name['inputs'].message_type = _JOBINPUT
_JOBMETADATADELTA.fields_by_name['outputs'].message_type = _JOBOUTPUT
_JOBINPUT.fields_by_name['type'].enum_type = _JOBINPUT_TYPE
_JOBINPUT_TYPE.containing_type = _JOBINPUT
_JOBOUTPUT.fields_by_name['type'].enum_type = _JOBOUTPUT_TYPE
_JOBOUTPUT_TYPE.containing_type = _JOBOUTPUT
_DATASETVERINPUT.fields_by_name['type'].enum_type = _DATASETVERINPUT_TYPE
_DATASETVERINPUT_TYPE.containing_type = _DATASETVERINPUT
_DATASETMETADATADELTA.fields_by_name['sample'].message_type = _DATASETSAMPLE
_DATASETSAMPLE.fields_by_name['format'].enum_type = _DATASETSAMPLE_FORMAT
_DATASETSAMPLE_FORMAT.containing_type = _DATASETSAMPLE
DESCRIPTOR.message_types_by_name['ProjectMetadataDelta'] = _PROJECTMETADATADELTA
DESCRIPTOR.message_types_by_name['ServiceMetadataDelta'] = _SERVICEMETADATADELTA
DESCRIPTOR.message_types_by_name['ExpertMetadataDelta'] = _EXPERTMETADATADELTA
DESCRIPTOR.message_types_by_name['Relation'] = _RELATION
DESCRIPTOR.message_types_by_name['JobMetadataDelta'] = _JOBMETADATADELTA
DESCRIPTOR.message_types_by_name['JobInput'] = _JOBINPUT
DESCRIPTOR.message_types_by_name['JobOutput'] = _JOBOUTPUT
DESCRIPTOR.message_types_by_name['DatasetVerInput'] = _DATASETVERINPUT
DESCRIPTOR.message_types_by_name['DatasetMetadataDelta'] = _DATASETMETADATADELTA
DESCRIPTOR.message_types_by_name['DatasetItem'] = _DATASETITEM
DESCRIPTOR.message_types_by_name['DatasetSample'] = _DATASETSAMPLE
DESCRIPTOR.enum_types_by_name['ENTITY'] = _ENTITY
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProjectMetadataDelta = _reflection.GeneratedProtocolMessageType('ProjectMetadataDelta', (_message.Message,), {
  'DESCRIPTOR' : _PROJECTMETADATADELTA,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:ProjectMetadataDelta)
  })
_sym_db.RegisterMessage(ProjectMetadataDelta)

ServiceMetadataDelta = _reflection.GeneratedProtocolMessageType('ServiceMetadataDelta', (_message.Message,), {
  'DESCRIPTOR' : _SERVICEMETADATADELTA,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:ServiceMetadataDelta)
  })
_sym_db.RegisterMessage(ServiceMetadataDelta)

ExpertMetadataDelta = _reflection.GeneratedProtocolMessageType('ExpertMetadataDelta', (_message.Message,), {
  'DESCRIPTOR' : _EXPERTMETADATADELTA,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:ExpertMetadataDelta)
  })
_sym_db.RegisterMessage(ExpertMetadataDelta)

Relation = _reflection.GeneratedProtocolMessageType('Relation', (_message.Message,), {
  'DESCRIPTOR' : _RELATION,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:Relation)
  })
_sym_db.RegisterMessage(Relation)

JobMetadataDelta = _reflection.GeneratedProtocolMessageType('JobMetadataDelta', (_message.Message,), {
  'DESCRIPTOR' : _JOBMETADATADELTA,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:JobMetadataDelta)
  })
_sym_db.RegisterMessage(JobMetadataDelta)

JobInput = _reflection.GeneratedProtocolMessageType('JobInput', (_message.Message,), {
  'DESCRIPTOR' : _JOBINPUT,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:JobInput)
  })
_sym_db.RegisterMessage(JobInput)

JobOutput = _reflection.GeneratedProtocolMessageType('JobOutput', (_message.Message,), {
  'DESCRIPTOR' : _JOBOUTPUT,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:JobOutput)
  })
_sym_db.RegisterMessage(JobOutput)

DatasetVerInput = _reflection.GeneratedProtocolMessageType('DatasetVerInput', (_message.Message,), {
  'DESCRIPTOR' : _DATASETVERINPUT,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:DatasetVerInput)
  })
_sym_db.RegisterMessage(DatasetVerInput)

DatasetMetadataDelta = _reflection.GeneratedProtocolMessageType('DatasetMetadataDelta', (_message.Message,), {
  'DESCRIPTOR' : _DATASETMETADATADELTA,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:DatasetMetadataDelta)
  })
_sym_db.RegisterMessage(DatasetMetadataDelta)

DatasetItem = _reflection.GeneratedProtocolMessageType('DatasetItem', (_message.Message,), {
  'DESCRIPTOR' : _DATASETITEM,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:DatasetItem)
  })
_sym_db.RegisterMessage(DatasetItem)

DatasetSample = _reflection.GeneratedProtocolMessageType('DatasetSample', (_message.Message,), {
  'DESCRIPTOR' : _DATASETSAMPLE,
  '__module__' : 'vo_pb2'
  # @@protoc_insertion_point(class_scope:DatasetSample)
  })
_sym_db.RegisterMessage(DatasetSample)


# @@protoc_insertion_point(module_scope)
