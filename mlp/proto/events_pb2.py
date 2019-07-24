# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/events.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='proto/events.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x12proto/events.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"2\n\x0eProjectCreated\x12\x12\n\nproject_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"j\n\x0e\x44\x61tasetCreated\x12\x12\n\ndataset_id\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x12\n\nproject_id\x18\x03 \x01(\t\x12\"\n\x08metadata\x18\x04 \x01(\x0b\x32\x10.DatasetMetadata\"\\\n\x0e\x44\x61tasetUpdated\x12\x12\n\ndataset_id\x18\x01 \x01(\t\x12\x12\n\nproject_id\x18\x02 \x01(\t\x12\"\n\x08metadata\x18\x04 \x01(\x0b\x32\x10.DatasetMetadata\"W\n\x10\x44\x61tasetLinkAdded\x12\x0f\n\x07link_id\x18\x01 \x01(\t\x12\x11\n\tlink_name\x18\x02 \x01(\t\x12\x0e\n\x06inputs\x18\x03 \x03(\t\x12\x0f\n\x07outputs\x18\x04 \x03(\t\"\xa1\x02\n\x0f\x44\x61tasetMetadata\x12 \n\nset_fields\x18\x01 \x03(\x0e\x32\x0c.FIELD_TYPES\x12 \n\ndel_fields\x18\x02 \x03(\x0e\x32\x0c.FIELD_TYPES\x12\x14\n\x0crecord_count\x18\x03 \x01(\x03\x12\x12\n\nfile_count\x18\x04 \x01(\x03\x12\x11\n\traw_bytes\x18\x05 \x01(\x03\x12\x11\n\tzip_bytes\x18\x06 \x01(\x03\x12\x13\n\x0bsample_body\x18\x07 \x01(\x0c\x12!\n\x0bsample_kind\x18\x08 \x01(\x0e\x32\x0c.SAMPLE_KIND\x12\x18\n\x10update_timestamp\x18\t \x01(\x03\x12\x13\n\x0b\x64\x61ta_format\x18\n \x01(\t\x12\x13\n\x0b\x64\x65scription\x18\x0b \x01(\t*\xd1\x01\n\x0b\x46IELD_TYPES\x12\x0e\n\nFIELD_NONE\x10\x00\x12\x16\n\x12\x46IELD_RECORD_COUNT\x10\x01\x12\x13\n\x0f\x46IELD_RAW_BYTES\x10\x02\x12\x13\n\x0f\x46IELD_ZIP_BYTES\x10\x03\x12\x10\n\x0c\x46IELD_SAMPLE\x10\x04\x12\x1a\n\x16\x46IELD_UPDATE_TIMESTAMP\x10\x05\x12\x14\n\x10\x46IELD_FILE_COUNT\x10\x06\x12\x15\n\x11\x46IELD_DATA_FORMAT\x10\x07\x12\x15\n\x11\x46IELD_DESCRIPTION\x10\x08**\n\x0bSAMPLE_KIND\x12\x08\n\x04TEXT\x10\x00\x12\x07\n\x03TSV\x10\x01\x12\x08\n\x04JSON\x10\x02\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_FIELD_TYPES = _descriptor.EnumDescriptor(
  name='FIELD_TYPES',
  full_name='FIELD_TYPES',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='FIELD_NONE', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_RECORD_COUNT', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_RAW_BYTES', index=2, number=2,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_ZIP_BYTES', index=3, number=3,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_SAMPLE', index=4, number=4,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_UPDATE_TIMESTAMP', index=5, number=5,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_FILE_COUNT', index=6, number=6,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_DATA_FORMAT', index=7, number=7,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FIELD_DESCRIPTION', index=8, number=8,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=691,
  serialized_end=900,
)
_sym_db.RegisterEnumDescriptor(_FIELD_TYPES)

FIELD_TYPES = enum_type_wrapper.EnumTypeWrapper(_FIELD_TYPES)
_SAMPLE_KIND = _descriptor.EnumDescriptor(
  name='SAMPLE_KIND',
  full_name='SAMPLE_KIND',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='TEXT', index=0, number=0,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TSV', index=1, number=1,
      options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='JSON', index=2, number=2,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=902,
  serialized_end=944,
)
_sym_db.RegisterEnumDescriptor(_SAMPLE_KIND)

SAMPLE_KIND = enum_type_wrapper.EnumTypeWrapper(_SAMPLE_KIND)
FIELD_NONE = 0
FIELD_RECORD_COUNT = 1
FIELD_RAW_BYTES = 2
FIELD_ZIP_BYTES = 3
FIELD_SAMPLE = 4
FIELD_UPDATE_TIMESTAMP = 5
FIELD_FILE_COUNT = 6
FIELD_DATA_FORMAT = 7
FIELD_DESCRIPTION = 8
TEXT = 0
TSV = 1
JSON = 2



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
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='ProjectCreated.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=55,
  serialized_end=105,
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
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='DatasetCreated.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetCreated.project_id', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='DatasetCreated.metadata', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=107,
  serialized_end=213,
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
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetUpdated.project_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metadata', full_name='DatasetUpdated.metadata', index=2,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=215,
  serialized_end=307,
)


_DATASETLINKADDED = _descriptor.Descriptor(
  name='DatasetLinkAdded',
  full_name='DatasetLinkAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='link_id', full_name='DatasetLinkAdded.link_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='link_name', full_name='DatasetLinkAdded.link_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs', full_name='DatasetLinkAdded.inputs', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='DatasetLinkAdded.outputs', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=309,
  serialized_end=396,
)


_DATASETMETADATA = _descriptor.Descriptor(
  name='DatasetMetadata',
  full_name='DatasetMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='set_fields', full_name='DatasetMetadata.set_fields', index=0,
      number=1, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='del_fields', full_name='DatasetMetadata.del_fields', index=1,
      number=2, type=14, cpp_type=8, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='record_count', full_name='DatasetMetadata.record_count', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='file_count', full_name='DatasetMetadata.file_count', index=3,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='raw_bytes', full_name='DatasetMetadata.raw_bytes', index=4,
      number=5, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='zip_bytes', full_name='DatasetMetadata.zip_bytes', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_body', full_name='DatasetMetadata.sample_body', index=6,
      number=7, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sample_kind', full_name='DatasetMetadata.sample_kind', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp', full_name='DatasetMetadata.update_timestamp', index=8,
      number=9, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data_format', full_name='DatasetMetadata.data_format', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='description', full_name='DatasetMetadata.description', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=399,
  serialized_end=688,
)

_DATASETCREATED.fields_by_name['metadata'].message_type = _DATASETMETADATA
_DATASETUPDATED.fields_by_name['metadata'].message_type = _DATASETMETADATA
_DATASETMETADATA.fields_by_name['set_fields'].enum_type = _FIELD_TYPES
_DATASETMETADATA.fields_by_name['del_fields'].enum_type = _FIELD_TYPES
_DATASETMETADATA.fields_by_name['sample_kind'].enum_type = _SAMPLE_KIND
DESCRIPTOR.message_types_by_name['ProjectCreated'] = _PROJECTCREATED
DESCRIPTOR.message_types_by_name['DatasetCreated'] = _DATASETCREATED
DESCRIPTOR.message_types_by_name['DatasetUpdated'] = _DATASETUPDATED
DESCRIPTOR.message_types_by_name['DatasetLinkAdded'] = _DATASETLINKADDED
DESCRIPTOR.message_types_by_name['DatasetMetadata'] = _DATASETMETADATA
DESCRIPTOR.enum_types_by_name['FIELD_TYPES'] = _FIELD_TYPES
DESCRIPTOR.enum_types_by_name['SAMPLE_KIND'] = _SAMPLE_KIND
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ProjectCreated = _reflection.GeneratedProtocolMessageType('ProjectCreated', (_message.Message,), dict(
  DESCRIPTOR = _PROJECTCREATED,
  __module__ = 'proto.events_pb2'
  # @@protoc_insertion_point(class_scope:ProjectCreated)
  ))
_sym_db.RegisterMessage(ProjectCreated)

DatasetCreated = _reflection.GeneratedProtocolMessageType('DatasetCreated', (_message.Message,), dict(
  DESCRIPTOR = _DATASETCREATED,
  __module__ = 'proto.events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetCreated)
  ))
_sym_db.RegisterMessage(DatasetCreated)

DatasetUpdated = _reflection.GeneratedProtocolMessageType('DatasetUpdated', (_message.Message,), dict(
  DESCRIPTOR = _DATASETUPDATED,
  __module__ = 'proto.events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetUpdated)
  ))
_sym_db.RegisterMessage(DatasetUpdated)

DatasetLinkAdded = _reflection.GeneratedProtocolMessageType('DatasetLinkAdded', (_message.Message,), dict(
  DESCRIPTOR = _DATASETLINKADDED,
  __module__ = 'proto.events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetLinkAdded)
  ))
_sym_db.RegisterMessage(DatasetLinkAdded)

DatasetMetadata = _reflection.GeneratedProtocolMessageType('DatasetMetadata', (_message.Message,), dict(
  DESCRIPTOR = _DATASETMETADATA,
  __module__ = 'proto.events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetMetadata)
  ))
_sym_db.RegisterMessage(DatasetMetadata)


# @@protoc_insertion_point(module_scope)
