# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: greet.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='greet.proto',
  package='greet',
  syntax='proto3',
  serialized_options=b'Z\007greetpb',
  serialized_pb=b'\n\x0bgreet.proto\x12\x05greet\"0\n\x08Greeting\x12\x11\n\tfirt_name\x18\x01 \x01(\t\x12\x11\n\tlast_name\x18\x02 \x01(\t\"4\n\x0fGreetingRequest\x12!\n\x08greeting\x18\x01 \x01(\x0b\x32\x0f.greet.Greeting\"\"\n\x10GreetingResponse\x12\x0e\n\x06result\x18\x01 \x01(\t\":\n\x15GreetManyTimesRequest\x12!\n\x08greeting\x18\x01 \x01(\x0b\x32\x0f.greet.Greeting\"(\n\x16GreetManyTimesResponse\x12\x0e\n\x06result\x18\x01 \x01(\t\"5\n\x10LongGreetRequest\x12!\n\x08greeting\x18\x01 \x01(\x0b\x32\x0f.greet.Greeting\"#\n\x11LongGreetResponse\x12\x0e\n\x06result\x18\x01 \x01(\t\"9\n\x14GreetEveryOneRequest\x12!\n\x08greeting\x18\x01 \x01(\x0b\x32\x0f.greet.Greeting\"\'\n\x15GreetEveryOneResponse\x12\x0e\n\x06result\x18\x01 \x01(\t\"8\n\x13WithDeadLineRequest\x12!\n\x08greeting\x18\x01 \x01(\x0b\x32\x0f.greet.Greeting\"&\n\x14WithDeadLineResponse\x12\x0e\n\x06result\x18\x01 \x01(\t2\x84\x03\n\x0cGreetService\x12:\n\x05Greet\x12\x16.greet.GreetingRequest\x1a\x17.greet.GreetingResponse\"\x00\x12Q\n\x0eGreetManyTimes\x12\x1c.greet.GreetManyTimesRequest\x1a\x1d.greet.GreetManyTimesResponse\"\x00\x30\x01\x12\x42\n\tLongGreet\x12\x17.greet.LongGreetRequest\x1a\x18.greet.LongGreetResponse\"\x00(\x01\x12P\n\rGreetEveryOne\x12\x1b.greet.GreetEveryOneRequest\x1a\x1c.greet.GreetEveryOneResponse\"\x00(\x01\x30\x01\x12O\n\x12GreetWithDeadLines\x12\x1a.greet.WithDeadLineRequest\x1a\x1b.greet.WithDeadLineResponse\"\x00\x42\tZ\x07greetpbb\x06proto3'
)




_GREETING = _descriptor.Descriptor(
  name='Greeting',
  full_name='greet.Greeting',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='firt_name', full_name='greet.Greeting.firt_name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='last_name', full_name='greet.Greeting.last_name', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=22,
  serialized_end=70,
)


_GREETINGREQUEST = _descriptor.Descriptor(
  name='GreetingRequest',
  full_name='greet.GreetingRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='greeting', full_name='greet.GreetingRequest.greeting', index=0,
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
  serialized_start=72,
  serialized_end=124,
)


_GREETINGRESPONSE = _descriptor.Descriptor(
  name='GreetingResponse',
  full_name='greet.GreetingResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='greet.GreetingResponse.result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=126,
  serialized_end=160,
)


_GREETMANYTIMESREQUEST = _descriptor.Descriptor(
  name='GreetManyTimesRequest',
  full_name='greet.GreetManyTimesRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='greeting', full_name='greet.GreetManyTimesRequest.greeting', index=0,
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
  serialized_start=162,
  serialized_end=220,
)


_GREETMANYTIMESRESPONSE = _descriptor.Descriptor(
  name='GreetManyTimesResponse',
  full_name='greet.GreetManyTimesResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='greet.GreetManyTimesResponse.result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=222,
  serialized_end=262,
)


_LONGGREETREQUEST = _descriptor.Descriptor(
  name='LongGreetRequest',
  full_name='greet.LongGreetRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='greeting', full_name='greet.LongGreetRequest.greeting', index=0,
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
  serialized_start=264,
  serialized_end=317,
)


_LONGGREETRESPONSE = _descriptor.Descriptor(
  name='LongGreetResponse',
  full_name='greet.LongGreetResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='greet.LongGreetResponse.result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=319,
  serialized_end=354,
)


_GREETEVERYONEREQUEST = _descriptor.Descriptor(
  name='GreetEveryOneRequest',
  full_name='greet.GreetEveryOneRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='greeting', full_name='greet.GreetEveryOneRequest.greeting', index=0,
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
  serialized_start=356,
  serialized_end=413,
)


_GREETEVERYONERESPONSE = _descriptor.Descriptor(
  name='GreetEveryOneResponse',
  full_name='greet.GreetEveryOneResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='greet.GreetEveryOneResponse.result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=415,
  serialized_end=454,
)


_WITHDEADLINEREQUEST = _descriptor.Descriptor(
  name='WithDeadLineRequest',
  full_name='greet.WithDeadLineRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='greeting', full_name='greet.WithDeadLineRequest.greeting', index=0,
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
  serialized_start=456,
  serialized_end=512,
)


_WITHDEADLINERESPONSE = _descriptor.Descriptor(
  name='WithDeadLineResponse',
  full_name='greet.WithDeadLineResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='result', full_name='greet.WithDeadLineResponse.result', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_end=552,
)

_GREETINGREQUEST.fields_by_name['greeting'].message_type = _GREETING
_GREETMANYTIMESREQUEST.fields_by_name['greeting'].message_type = _GREETING
_LONGGREETREQUEST.fields_by_name['greeting'].message_type = _GREETING
_GREETEVERYONEREQUEST.fields_by_name['greeting'].message_type = _GREETING
_WITHDEADLINEREQUEST.fields_by_name['greeting'].message_type = _GREETING
DESCRIPTOR.message_types_by_name['Greeting'] = _GREETING
DESCRIPTOR.message_types_by_name['GreetingRequest'] = _GREETINGREQUEST
DESCRIPTOR.message_types_by_name['GreetingResponse'] = _GREETINGRESPONSE
DESCRIPTOR.message_types_by_name['GreetManyTimesRequest'] = _GREETMANYTIMESREQUEST
DESCRIPTOR.message_types_by_name['GreetManyTimesResponse'] = _GREETMANYTIMESRESPONSE
DESCRIPTOR.message_types_by_name['LongGreetRequest'] = _LONGGREETREQUEST
DESCRIPTOR.message_types_by_name['LongGreetResponse'] = _LONGGREETRESPONSE
DESCRIPTOR.message_types_by_name['GreetEveryOneRequest'] = _GREETEVERYONEREQUEST
DESCRIPTOR.message_types_by_name['GreetEveryOneResponse'] = _GREETEVERYONERESPONSE
DESCRIPTOR.message_types_by_name['WithDeadLineRequest'] = _WITHDEADLINEREQUEST
DESCRIPTOR.message_types_by_name['WithDeadLineResponse'] = _WITHDEADLINERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Greeting = _reflection.GeneratedProtocolMessageType('Greeting', (_message.Message,), {
  'DESCRIPTOR' : _GREETING,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.Greeting)
  })
_sym_db.RegisterMessage(Greeting)

GreetingRequest = _reflection.GeneratedProtocolMessageType('GreetingRequest', (_message.Message,), {
  'DESCRIPTOR' : _GREETINGREQUEST,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.GreetingRequest)
  })
_sym_db.RegisterMessage(GreetingRequest)

GreetingResponse = _reflection.GeneratedProtocolMessageType('GreetingResponse', (_message.Message,), {
  'DESCRIPTOR' : _GREETINGRESPONSE,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.GreetingResponse)
  })
_sym_db.RegisterMessage(GreetingResponse)

GreetManyTimesRequest = _reflection.GeneratedProtocolMessageType('GreetManyTimesRequest', (_message.Message,), {
  'DESCRIPTOR' : _GREETMANYTIMESREQUEST,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.GreetManyTimesRequest)
  })
_sym_db.RegisterMessage(GreetManyTimesRequest)

GreetManyTimesResponse = _reflection.GeneratedProtocolMessageType('GreetManyTimesResponse', (_message.Message,), {
  'DESCRIPTOR' : _GREETMANYTIMESRESPONSE,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.GreetManyTimesResponse)
  })
_sym_db.RegisterMessage(GreetManyTimesResponse)

LongGreetRequest = _reflection.GeneratedProtocolMessageType('LongGreetRequest', (_message.Message,), {
  'DESCRIPTOR' : _LONGGREETREQUEST,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.LongGreetRequest)
  })
_sym_db.RegisterMessage(LongGreetRequest)

LongGreetResponse = _reflection.GeneratedProtocolMessageType('LongGreetResponse', (_message.Message,), {
  'DESCRIPTOR' : _LONGGREETRESPONSE,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.LongGreetResponse)
  })
_sym_db.RegisterMessage(LongGreetResponse)

GreetEveryOneRequest = _reflection.GeneratedProtocolMessageType('GreetEveryOneRequest', (_message.Message,), {
  'DESCRIPTOR' : _GREETEVERYONEREQUEST,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.GreetEveryOneRequest)
  })
_sym_db.RegisterMessage(GreetEveryOneRequest)

GreetEveryOneResponse = _reflection.GeneratedProtocolMessageType('GreetEveryOneResponse', (_message.Message,), {
  'DESCRIPTOR' : _GREETEVERYONERESPONSE,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.GreetEveryOneResponse)
  })
_sym_db.RegisterMessage(GreetEveryOneResponse)

WithDeadLineRequest = _reflection.GeneratedProtocolMessageType('WithDeadLineRequest', (_message.Message,), {
  'DESCRIPTOR' : _WITHDEADLINEREQUEST,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.WithDeadLineRequest)
  })
_sym_db.RegisterMessage(WithDeadLineRequest)

WithDeadLineResponse = _reflection.GeneratedProtocolMessageType('WithDeadLineResponse', (_message.Message,), {
  'DESCRIPTOR' : _WITHDEADLINERESPONSE,
  '__module__' : 'greet_pb2'
  # @@protoc_insertion_point(class_scope:greet.WithDeadLineResponse)
  })
_sym_db.RegisterMessage(WithDeadLineResponse)


DESCRIPTOR._options = None

_GREETSERVICE = _descriptor.ServiceDescriptor(
  name='GreetService',
  full_name='greet.GreetService',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=555,
  serialized_end=943,
  methods=[
  _descriptor.MethodDescriptor(
    name='Greet',
    full_name='greet.GreetService.Greet',
    index=0,
    containing_service=None,
    input_type=_GREETINGREQUEST,
    output_type=_GREETINGRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GreetManyTimes',
    full_name='greet.GreetService.GreetManyTimes',
    index=1,
    containing_service=None,
    input_type=_GREETMANYTIMESREQUEST,
    output_type=_GREETMANYTIMESRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='LongGreet',
    full_name='greet.GreetService.LongGreet',
    index=2,
    containing_service=None,
    input_type=_LONGGREETREQUEST,
    output_type=_LONGGREETRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GreetEveryOne',
    full_name='greet.GreetService.GreetEveryOne',
    index=3,
    containing_service=None,
    input_type=_GREETEVERYONEREQUEST,
    output_type=_GREETEVERYONERESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GreetWithDeadLines',
    full_name='greet.GreetService.GreetWithDeadLines',
    index=4,
    containing_service=None,
    input_type=_WITHDEADLINEREQUEST,
    output_type=_WITHDEADLINERESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_GREETSERVICE)

DESCRIPTOR.services_by_name['GreetService'] = _GREETSERVICE

# @@protoc_insertion_point(module_scope)
