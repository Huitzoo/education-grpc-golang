# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

import greet_pb2 as greet__pb2


class GreetServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.Greet = channel.unary_unary(
        '/greet.GreetService/Greet',
        request_serializer=greet__pb2.GreetingRequest.SerializeToString,
        response_deserializer=greet__pb2.GreetingResponse.FromString,
        )
    self.GreetManyTimes = channel.unary_stream(
        '/greet.GreetService/GreetManyTimes',
        request_serializer=greet__pb2.GreetManyTimesRequest.SerializeToString,
        response_deserializer=greet__pb2.GreetManyTimesResponse.FromString,
        )
    self.LongGreet = channel.stream_unary(
        '/greet.GreetService/LongGreet',
        request_serializer=greet__pb2.LongGreetRequest.SerializeToString,
        response_deserializer=greet__pb2.LongGreetResponse.FromString,
        )
    self.GreetEveryOne = channel.stream_stream(
        '/greet.GreetService/GreetEveryOne',
        request_serializer=greet__pb2.GreetEveryOneRequest.SerializeToString,
        response_deserializer=greet__pb2.GreetEveryOneResponse.FromString,
        )
    self.GreetWithDeadLines = channel.unary_unary(
        '/greet.GreetService/GreetWithDeadLines',
        request_serializer=greet__pb2.WithDeadLineRequest.SerializeToString,
        response_deserializer=greet__pb2.WithDeadLineResponse.FromString,
        )


class GreetServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def Greet(self, request, context):
    """Unary
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GreetManyTimes(self, request, context):
    """Server streaming
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def LongGreet(self, request_iterator, context):
    """Client streaming
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GreetEveryOne(self, request_iterator, context):
    """bi direccional stream
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def GreetWithDeadLines(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_GreetServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'Greet': grpc.unary_unary_rpc_method_handler(
          servicer.Greet,
          request_deserializer=greet__pb2.GreetingRequest.FromString,
          response_serializer=greet__pb2.GreetingResponse.SerializeToString,
      ),
      'GreetManyTimes': grpc.unary_stream_rpc_method_handler(
          servicer.GreetManyTimes,
          request_deserializer=greet__pb2.GreetManyTimesRequest.FromString,
          response_serializer=greet__pb2.GreetManyTimesResponse.SerializeToString,
      ),
      'LongGreet': grpc.stream_unary_rpc_method_handler(
          servicer.LongGreet,
          request_deserializer=greet__pb2.LongGreetRequest.FromString,
          response_serializer=greet__pb2.LongGreetResponse.SerializeToString,
      ),
      'GreetEveryOne': grpc.stream_stream_rpc_method_handler(
          servicer.GreetEveryOne,
          request_deserializer=greet__pb2.GreetEveryOneRequest.FromString,
          response_serializer=greet__pb2.GreetEveryOneResponse.SerializeToString,
      ),
      'GreetWithDeadLines': grpc.unary_unary_rpc_method_handler(
          servicer.GreetWithDeadLines,
          request_deserializer=greet__pb2.WithDeadLineRequest.FromString,
          response_serializer=greet__pb2.WithDeadLineResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'greet.GreetService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))