# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: proto/hello.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x11proto/hello.proto\x12\x05hello\"\x1c\n\x0cHelloRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"\x1e\n\rHelloResponse\x12\r\n\x05hello\x18\x01 \x01(\t\"(\n\x18HelloClientStreamRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"/\n\x19HelloClientStreamResponse\x12\x12\n\nhello_list\x18\x01 \x03(\t\"(\n\x18HelloServerStreamRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"*\n\x19HelloServerStreamResponse\x12\r\n\x05hello\x18\x01 \x01(\t\")\n\x19HelloBidirectionalRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\"+\n\x1aHelloBidirectionalResponse\x12\r\n\x05hello\x18\x01 \x01(\t2\xdd\x02\n\x0cHelloService\x12\x34\n\x05Hello\x12\x13.hello.HelloRequest\x1a\x14.hello.HelloResponse\"\x00\x12Z\n\x11HelloClientStream\x12\x1f.hello.HelloClientStreamRequest\x1a .hello.HelloClientStreamResponse\"\x00(\x01\x12Z\n\x11HelloServerStream\x12\x1f.hello.HelloServerStreamRequest\x1a .hello.HelloServerStreamResponse\"\x00\x30\x01\x12_\n\x12HelloBidirectional\x12 .hello.HelloBidirectionalRequest\x1a!.hello.HelloBidirectionalResponse\"\x00(\x01\x30\x01\x42\x06Z\x04./pbb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'proto.hello_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\004./pb'
  _globals['_HELLOREQUEST']._serialized_start=28
  _globals['_HELLOREQUEST']._serialized_end=56
  _globals['_HELLORESPONSE']._serialized_start=58
  _globals['_HELLORESPONSE']._serialized_end=88
  _globals['_HELLOCLIENTSTREAMREQUEST']._serialized_start=90
  _globals['_HELLOCLIENTSTREAMREQUEST']._serialized_end=130
  _globals['_HELLOCLIENTSTREAMRESPONSE']._serialized_start=132
  _globals['_HELLOCLIENTSTREAMRESPONSE']._serialized_end=179
  _globals['_HELLOSERVERSTREAMREQUEST']._serialized_start=181
  _globals['_HELLOSERVERSTREAMREQUEST']._serialized_end=221
  _globals['_HELLOSERVERSTREAMRESPONSE']._serialized_start=223
  _globals['_HELLOSERVERSTREAMRESPONSE']._serialized_end=265
  _globals['_HELLOBIDIRECTIONALREQUEST']._serialized_start=267
  _globals['_HELLOBIDIRECTIONALREQUEST']._serialized_end=308
  _globals['_HELLOBIDIRECTIONALRESPONSE']._serialized_start=310
  _globals['_HELLOBIDIRECTIONALRESPONSE']._serialized_end=353
  _globals['_HELLOSERVICE']._serialized_start=356
  _globals['_HELLOSERVICE']._serialized_end=705
# @@protoc_insertion_point(module_scope)
