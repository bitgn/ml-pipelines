# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='api.proto',
  package='',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n\tapi.proto\">\n\x14\x43reateProjectRequest\x12\x11\n\tProjectId\x18\x01 \x01(\t\x12\x13\n\x0bProjectName\x18\x02 \x01(\t\"\x17\n\x15\x43reateProjectResponse\"7\n\rCreateDataset\x12\x12\n\ndataset_id\x18\x01 \x01(\t\x12\x12\n\nproject_id\x18\x02 \x01(\t2K\n\x07\x43\x61talog\x12@\n\rCreateProject\x12\x15.CreateProjectRequest\x1a\x16.CreateProjectResponse\"\x00\x62\x06proto3')
)




_CREATEPROJECTREQUEST = _descriptor.Descriptor(
  name='CreateProjectRequest',
  full_name='CreateProjectRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ProjectId', full_name='CreateProjectRequest.ProjectId', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='ProjectName', full_name='CreateProjectRequest.ProjectName', index=1,
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
  serialized_start=13,
  serialized_end=75,
)


_CREATEPROJECTRESPONSE = _descriptor.Descriptor(
  name='CreateProjectResponse',
  full_name='CreateProjectResponse',
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
  serialized_start=77,
  serialized_end=100,
)


_CREATEDATASET = _descriptor.Descriptor(
  name='CreateDataset',
  full_name='CreateDataset',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='dataset_id', full_name='CreateDataset.dataset_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='project_id', full_name='CreateDataset.project_id', index=1,
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
  serialized_start=102,
  serialized_end=157,
)

DESCRIPTOR.message_types_by_name['CreateProjectRequest'] = _CREATEPROJECTREQUEST
DESCRIPTOR.message_types_by_name['CreateProjectResponse'] = _CREATEPROJECTRESPONSE
DESCRIPTOR.message_types_by_name['CreateDataset'] = _CREATEDATASET
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateProjectRequest = _reflection.GeneratedProtocolMessageType('CreateProjectRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTREQUEST,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:CreateProjectRequest)
  })
_sym_db.RegisterMessage(CreateProjectRequest)

CreateProjectResponse = _reflection.GeneratedProtocolMessageType('CreateProjectResponse', (_message.Message,), {
  'DESCRIPTOR' : _CREATEPROJECTRESPONSE,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:CreateProjectResponse)
  })
_sym_db.RegisterMessage(CreateProjectResponse)

CreateDataset = _reflection.GeneratedProtocolMessageType('CreateDataset', (_message.Message,), {
  'DESCRIPTOR' : _CREATEDATASET,
  '__module__' : 'api_pb2'
  # @@protoc_insertion_point(class_scope:CreateDataset)
  })
_sym_db.RegisterMessage(CreateDataset)



_CATALOG = _descriptor.ServiceDescriptor(
  name='Catalog',
  full_name='Catalog',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=159,
  serialized_end=234,
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
])
_sym_db.RegisterServiceDescriptor(_CATALOG)

DESCRIPTOR.services_by_name['Catalog'] = _CATALOG

# @@protoc_insertion_point(module_scope)
