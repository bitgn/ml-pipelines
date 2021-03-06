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


from . import vo_pb2 as vo__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='events.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\x0c\x65vents.proto\x1a\x08vo.proto\"$\n\x0eProjectCreated\x12\x12\n\nproject_id\x18\x01 \x01(\t\"]\n\x0e\x44\x61tasetCreated\x12\x12\n\nproject_id\x18\x01 \x01(\t\x12\x12\n\ndataset_id\x18\x02 \x01(\t\x12#\n\x04meta\x18\x04 \x01(\x0b\x32\x15.DatasetMetadataDelta\"]\n\x0e\x44\x61tasetUpdated\x12\x12\n\nproject_id\x18\x01 \x01(\t\x12\x12\n\ndataset_id\x18\x02 \x01(\t\x12#\n\x04meta\x18\x03 \x01(\x0b\x32\x15.DatasetMetadataDelta\"\x90\x01\n\x14\x44\x61tasetActivityAdded\x12\x12\n\nproject_id\x18\x01 \x01(\t\x12\x12\n\ndataset_id\x18\x02 \x01(\t\x12\x18\n\x10update_timestamp\x18\x04 \x01(\x03\x12\x16\n\x0emultiline_text\x18\x05 \x01(\t\x12\x1e\n\x05level\x18\x06 \x01(\x0e\x32\x0f.ACTIVITY_LEVEL\"*\n\x05\x45vent\x12\x0c\n\x04\x42ody\x18\x01 \x01(\x0c\x12\x13\n\x04Type\x18\x02 \x01(\x0e\x32\x05.Type*~\n\x04Type\x12\x08\n\x04None\x10\x00\x12\x18\n\x14\x45vent_ProjectCreated\x10\x01\x12\x18\n\x14\x45vent_DatasetCreated\x10\x02\x12\x18\n\x14\x45vent_DatasetUpdated\x10\x03\x12\x1e\n\x1a\x45vent_DatasetActivityAdded\x10\x04\x62\x06proto3')
  ,
  dependencies=[vo__pb2.DESCRIPTOR,])

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
      name='Event_DatasetActivityAdded', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=445,
  serialized_end=571,
)
_sym_db.RegisterEnumDescriptor(_TYPE)

Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
globals()['None'] = 0
Event_ProjectCreated = 1
Event_DatasetCreated = 2
Event_DatasetUpdated = 3
Event_DatasetActivityAdded = 4



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
  serialized_start=26,
  serialized_end=62,
)


_DATASETCREATED = _descriptor.Descriptor(
  name='DatasetCreated',
  full_name='DatasetCreated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetCreated.project_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='DatasetCreated.dataset_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='DatasetCreated.meta', index=2,
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
  serialized_start=64,
  serialized_end=157,
)


_DATASETUPDATED = _descriptor.Descriptor(
  name='DatasetUpdated',
  full_name='DatasetUpdated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetUpdated.project_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='DatasetUpdated.dataset_id', index=1,
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
  serialized_start=159,
  serialized_end=252,
)


_DATASETACTIVITYADDED = _descriptor.Descriptor(
  name='DatasetActivityAdded',
  full_name='DatasetActivityAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_id', full_name='DatasetActivityAdded.project_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='DatasetActivityAdded.dataset_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='update_timestamp', full_name='DatasetActivityAdded.update_timestamp', index=2,
      number=4, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='multiline_text', full_name='DatasetActivityAdded.multiline_text', index=3,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='level', full_name='DatasetActivityAdded.level', index=4,
      number=6, type=14, cpp_type=8, label=1,
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
  serialized_start=255,
  serialized_end=399,
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
  serialized_start=401,
  serialized_end=443,
)

_DATASETCREATED.fields_by_name['meta'].message_type = vo__pb2._DATASETMETADATADELTA
_DATASETUPDATED.fields_by_name['meta'].message_type = vo__pb2._DATASETMETADATADELTA
_DATASETACTIVITYADDED.fields_by_name['level'].enum_type = vo__pb2._ACTIVITY_LEVEL
_EVENT.fields_by_name['Type'].enum_type = _TYPE
DESCRIPTOR.message_types_by_name['ProjectCreated'] = _PROJECTCREATED
DESCRIPTOR.message_types_by_name['DatasetCreated'] = _DATASETCREATED
DESCRIPTOR.message_types_by_name['DatasetUpdated'] = _DATASETUPDATED
DESCRIPTOR.message_types_by_name['DatasetActivityAdded'] = _DATASETACTIVITYADDED
DESCRIPTOR.message_types_by_name['Event'] = _EVENT
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

DatasetActivityAdded = _reflection.GeneratedProtocolMessageType('DatasetActivityAdded', (_message.Message,), {
  'DESCRIPTOR' : _DATASETACTIVITYADDED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetActivityAdded)
  })
_sym_db.RegisterMessage(DatasetActivityAdded)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), {
  'DESCRIPTOR' : _EVENT,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:Event)
  })
_sym_db.RegisterMessage(Event)


# @@protoc_insertion_point(module_scope)
