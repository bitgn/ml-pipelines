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
  serialized_pb=_b('\n\x0c\x65vents.proto\x1a\x08vo.proto\"P\n\x0eProjectCreated\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12\x0c\n\x04name\x18\x02 \x01(\t\x12#\n\x04meta\x18\x03 \x01(\x0b\x32\x15.ProjectMetadataDelta\"{\n\x0e\x44\x61tasetCreated\x12\x13\n\x0bproject_uid\x18\x01 \x01(\x0c\x12\x0b\n\x03uid\x18\x02 \x01(\x0c\x12\x0c\n\x04name\x18\x03 \x01(\t\x12#\n\x04meta\x18\x04 \x01(\x0b\x32\x15.DatasetMetadataDelta\x12\x14\n\x0cproject_name\x18\x05 \x01(\t\"W\n\x0e\x44\x61tasetUpdated\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12\x13\n\x0bproject_uid\x18\x02 \x01(\x0c\x12#\n\x04meta\x18\x03 \x01(\x0b\x32\x15.DatasetMetadataDelta\"q\n\x08JobAdded\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12\x13\n\x0bproject_uid\x18\x02 \x01(\x0c\x12\x0c\n\x04name\x18\x03 \x01(\t\x12\x1f\n\x04meta\x18\x04 \x01(\x0b\x32\x11.JobMetadataDelta\x12\x14\n\x0cproject_name\x18\x05 \x01(\t\"{\n\x0eServiceCreated\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12\x13\n\x0bproject_uid\x18\x02 \x01(\x0c\x12\x0c\n\x04name\x18\x03 \x01(\t\x12#\n\x04meta\x18\x04 \x01(\x0b\x32\x15.ServiceMetadataDelta\x12\x14\n\x0cproject_name\x18\x05 \x01(\t\"\xd7\x01\n\x13\x44\x61tasetVersionAdded\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12\x13\n\x0bproject_uid\x18\x02 \x01(\x0c\x12\x12\n\nparent_uid\x18\x03 \x01(\x0c\x12\x14\n\x0cproject_name\x18\x04 \x01(\t\x12\r\n\x05title\x18\x05 \x01(\t\x12\x11\n\ttimestamp\x18\x06 \x01(\x03\x12\x1b\n\x05items\x18\x07 \x03(\x0b\x32\x0c.DatasetItem\x12 \n\x06inputs\x18\x08 \x03(\x0b\x32\x10.DatasetVerInput\x12\x13\n\x0b\x64\x61taset_uid\x18\t \x01(\x0c\"L\n\x0b\x45xpertAdded\x12\x0b\n\x03uid\x18\x01 \x01(\x0c\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\"\n\x04meta\x18\x03 \x01(\x0b\x32\x14.ExpertMetadataDelta\"*\n\x05\x45vent\x12\x0c\n\x04\x42ody\x18\x01 \x01(\x0c\x12\x13\n\x04Type\x18\x02 \x01(\x0e\x32\x05.Type*\xc2\x01\n\x04Type\x12\x08\n\x04None\x10\x00\x12\x18\n\x14\x45vent_ProjectCreated\x10\x01\x12\x18\n\x14\x45vent_DatasetCreated\x10\x02\x12\x18\n\x14\x45vent_DatasetUpdated\x10\x03\x12\x15\n\x11\x45vent_ExpertAdded\x10\x04\x12\x12\n\x0e\x45vent_JobAdded\x10\x05\x12\x18\n\x14\x45vent_ServiceCreated\x10\x06\x12\x1d\n\x19\x45vent_DatasetVersionAdded\x10\x07\x62\x06proto3')
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
      name='Event_ExpertAdded', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_JobAdded', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_ServiceCreated', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='Event_DatasetVersionAdded', index=7, number=7,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=903,
  serialized_end=1097,
)
_sym_db.RegisterEnumDescriptor(_TYPE)

Type = enum_type_wrapper.EnumTypeWrapper(_TYPE)
globals()['None'] = 0
Event_ProjectCreated = 1
Event_DatasetCreated = 2
Event_DatasetUpdated = 3
Event_ExpertAdded = 4
Event_JobAdded = 5
Event_ServiceCreated = 6
Event_DatasetVersionAdded = 7



_PROJECTCREATED = _descriptor.Descriptor(
  name='ProjectCreated',
  full_name='ProjectCreated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='ProjectCreated.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
    _descriptor.FieldDescriptor(
      name='meta', full_name='ProjectCreated.meta', index=2,
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
  serialized_start=26,
  serialized_end=106,
)


_DATASETCREATED = _descriptor.Descriptor(
  name='DatasetCreated',
  full_name='DatasetCreated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='DatasetCreated.project_uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uid', full_name='DatasetCreated.uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
    _descriptor.FieldDescriptor(
      name='project_name', full_name='DatasetCreated.project_name', index=4,
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
  serialized_start=108,
  serialized_end=231,
)


_DATASETUPDATED = _descriptor.Descriptor(
  name='DatasetUpdated',
  full_name='DatasetUpdated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='DatasetUpdated.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='DatasetUpdated.project_uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=233,
  serialized_end=320,
)


_JOBADDED = _descriptor.Descriptor(
  name='JobAdded',
  full_name='JobAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='JobAdded.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='JobAdded.project_uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='JobAdded.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='JobAdded.meta', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='JobAdded.project_name', index=4,
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
  serialized_start=322,
  serialized_end=435,
)


_SERVICECREATED = _descriptor.Descriptor(
  name='ServiceCreated',
  full_name='ServiceCreated',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='ServiceCreated.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='ServiceCreated.project_uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='ServiceCreated.name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='ServiceCreated.meta', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='ServiceCreated.project_name', index=4,
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
  serialized_start=437,
  serialized_end=560,
)


_DATASETVERSIONADDED = _descriptor.Descriptor(
  name='DatasetVersionAdded',
  full_name='DatasetVersionAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='DatasetVersionAdded.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_uid', full_name='DatasetVersionAdded.project_uid', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_uid', full_name='DatasetVersionAdded.parent_uid', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_name', full_name='DatasetVersionAdded.project_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='title', full_name='DatasetVersionAdded.title', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='timestamp', full_name='DatasetVersionAdded.timestamp', index=5,
      number=6, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='items', full_name='DatasetVersionAdded.items', index=6,
      number=7, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='inputs', full_name='DatasetVersionAdded.inputs', index=7,
      number=8, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='dataset_uid', full_name='DatasetVersionAdded.dataset_uid', index=8,
      number=9, type=12, cpp_type=9, label=1,
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
  serialized_start=563,
  serialized_end=778,
)


_EXPERTADDED = _descriptor.Descriptor(
  name='ExpertAdded',
  full_name='ExpertAdded',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='uid', full_name='ExpertAdded.uid', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='name', full_name='ExpertAdded.name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='meta', full_name='ExpertAdded.meta', index=2,
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
  serialized_start=780,
  serialized_end=856,
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
  serialized_start=858,
  serialized_end=900,
)

_PROJECTCREATED.fields_by_name['meta'].message_type = vo__pb2._PROJECTMETADATADELTA
_DATASETCREATED.fields_by_name['meta'].message_type = vo__pb2._DATASETMETADATADELTA
_DATASETUPDATED.fields_by_name['meta'].message_type = vo__pb2._DATASETMETADATADELTA
_JOBADDED.fields_by_name['meta'].message_type = vo__pb2._JOBMETADATADELTA
_SERVICECREATED.fields_by_name['meta'].message_type = vo__pb2._SERVICEMETADATADELTA
_DATASETVERSIONADDED.fields_by_name['items'].message_type = vo__pb2._DATASETITEM
_DATASETVERSIONADDED.fields_by_name['inputs'].message_type = vo__pb2._DATASETVERINPUT
_EXPERTADDED.fields_by_name['meta'].message_type = vo__pb2._EXPERTMETADATADELTA
_EVENT.fields_by_name['Type'].enum_type = _TYPE
DESCRIPTOR.message_types_by_name['ProjectCreated'] = _PROJECTCREATED
DESCRIPTOR.message_types_by_name['DatasetCreated'] = _DATASETCREATED
DESCRIPTOR.message_types_by_name['DatasetUpdated'] = _DATASETUPDATED
DESCRIPTOR.message_types_by_name['JobAdded'] = _JOBADDED
DESCRIPTOR.message_types_by_name['ServiceCreated'] = _SERVICECREATED
DESCRIPTOR.message_types_by_name['DatasetVersionAdded'] = _DATASETVERSIONADDED
DESCRIPTOR.message_types_by_name['ExpertAdded'] = _EXPERTADDED
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

JobAdded = _reflection.GeneratedProtocolMessageType('JobAdded', (_message.Message,), {
  'DESCRIPTOR' : _JOBADDED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:JobAdded)
  })
_sym_db.RegisterMessage(JobAdded)

ServiceCreated = _reflection.GeneratedProtocolMessageType('ServiceCreated', (_message.Message,), {
  'DESCRIPTOR' : _SERVICECREATED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:ServiceCreated)
  })
_sym_db.RegisterMessage(ServiceCreated)

DatasetVersionAdded = _reflection.GeneratedProtocolMessageType('DatasetVersionAdded', (_message.Message,), {
  'DESCRIPTOR' : _DATASETVERSIONADDED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:DatasetVersionAdded)
  })
_sym_db.RegisterMessage(DatasetVersionAdded)

ExpertAdded = _reflection.GeneratedProtocolMessageType('ExpertAdded', (_message.Message,), {
  'DESCRIPTOR' : _EXPERTADDED,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:ExpertAdded)
  })
_sym_db.RegisterMessage(ExpertAdded)

Event = _reflection.GeneratedProtocolMessageType('Event', (_message.Message,), {
  'DESCRIPTOR' : _EVENT,
  '__module__' : 'events_pb2'
  # @@protoc_insertion_point(class_scope:Event)
  })
_sym_db.RegisterMessage(Event)


# @@protoc_insertion_point(module_scope)
