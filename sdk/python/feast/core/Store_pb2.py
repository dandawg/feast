# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: feast/core/Store.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='feast/core/Store.proto',
  package='feast.core',
  syntax='proto3',
  serialized_options=_b('\n\nfeast.coreB\nStoreProtoZ5github.com/gojek/feast/protos/generated/go/feast/core'),
  serialized_pb=_b('\n\x16\x66\x65\x61st/core/Store.proto\x12\nfeast.core\"\xae\x01\n\x05Store\x12\x0c\n\x04name\x18\x01 \x01(\t\x12)\n\x04type\x18\x02 \x01(\x0e\x32\x1b.feast.core.Store.StoreType\x12\x13\n\x0b\x64\x61tabaseURI\x18\x03 \x01(\t\x12\x15\n\rsubscriptions\x18\x04 \x03(\t\"@\n\tStoreType\x12\x0b\n\x07UNKNOWN\x10\x00\x12\t\n\x05REDIS\x10\x01\x12\x0c\n\x08\x42IGQUERY\x10\x02\x12\r\n\tCASSANDRA\x10\x03\x42O\n\nfeast.coreB\nStoreProtoZ5github.com/gojek/feast/protos/generated/go/feast/coreb\x06proto3')
)



_STORE_STORETYPE = _descriptor.EnumDescriptor(
  name='StoreType',
  full_name='feast.core.Store.StoreType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='REDIS', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='BIGQUERY', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CASSANDRA', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=149,
  serialized_end=213,
)
_sym_db.RegisterEnumDescriptor(_STORE_STORETYPE)


_STORE = _descriptor.Descriptor(
  name='Store',
  full_name='feast.core.Store',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='feast.core.Store.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='type', full_name='feast.core.Store.type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='databaseURI', full_name='feast.core.Store.databaseURI', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='subscriptions', full_name='feast.core.Store.subscriptions', index=3,
      number=4, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _STORE_STORETYPE,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=39,
  serialized_end=213,
)

_STORE.fields_by_name['type'].enum_type = _STORE_STORETYPE
_STORE_STORETYPE.containing_type = _STORE
DESCRIPTOR.message_types_by_name['Store'] = _STORE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Store = _reflection.GeneratedProtocolMessageType('Store', (_message.Message,), {
  'DESCRIPTOR' : _STORE,
  '__module__' : 'feast.core.Store_pb2'
  # @@protoc_insertion_point(class_scope:feast.core.Store)
  })
_sym_db.RegisterMessage(Store)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
