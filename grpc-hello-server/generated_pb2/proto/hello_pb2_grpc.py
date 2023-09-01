# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import hello_pb2 as proto_dot_hello__pb2


class HelloServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Hello = channel.unary_unary(
                '/hello.HelloService/Hello',
                request_serializer=proto_dot_hello__pb2.HelloRequest.SerializeToString,
                response_deserializer=proto_dot_hello__pb2.HelloResponse.FromString,
                )
        self.HelloClientStream = channel.stream_unary(
                '/hello.HelloService/HelloClientStream',
                request_serializer=proto_dot_hello__pb2.HelloClientStreamRequest.SerializeToString,
                response_deserializer=proto_dot_hello__pb2.HelloClientStreamResponse.FromString,
                )
        self.HelloServerStream = channel.unary_stream(
                '/hello.HelloService/HelloServerStream',
                request_serializer=proto_dot_hello__pb2.HelloServerStreamRequest.SerializeToString,
                response_deserializer=proto_dot_hello__pb2.HelloServerStreamResponse.FromString,
                )
        self.HelloBidirectional = channel.stream_stream(
                '/hello.HelloService/HelloBidirectional',
                request_serializer=proto_dot_hello__pb2.HelloBidirectionalRequest.SerializeToString,
                response_deserializer=proto_dot_hello__pb2.HelloBidirectionalResponse.FromString,
                )


class HelloServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Hello(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HelloClientStream(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HelloServerStream(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def HelloBidirectional(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_HelloServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Hello': grpc.unary_unary_rpc_method_handler(
                    servicer.Hello,
                    request_deserializer=proto_dot_hello__pb2.HelloRequest.FromString,
                    response_serializer=proto_dot_hello__pb2.HelloResponse.SerializeToString,
            ),
            'HelloClientStream': grpc.stream_unary_rpc_method_handler(
                    servicer.HelloClientStream,
                    request_deserializer=proto_dot_hello__pb2.HelloClientStreamRequest.FromString,
                    response_serializer=proto_dot_hello__pb2.HelloClientStreamResponse.SerializeToString,
            ),
            'HelloServerStream': grpc.unary_stream_rpc_method_handler(
                    servicer.HelloServerStream,
                    request_deserializer=proto_dot_hello__pb2.HelloServerStreamRequest.FromString,
                    response_serializer=proto_dot_hello__pb2.HelloServerStreamResponse.SerializeToString,
            ),
            'HelloBidirectional': grpc.stream_stream_rpc_method_handler(
                    servicer.HelloBidirectional,
                    request_deserializer=proto_dot_hello__pb2.HelloBidirectionalRequest.FromString,
                    response_serializer=proto_dot_hello__pb2.HelloBidirectionalResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'hello.HelloService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class HelloService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Hello(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/hello.HelloService/Hello',
            proto_dot_hello__pb2.HelloRequest.SerializeToString,
            proto_dot_hello__pb2.HelloResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HelloClientStream(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_unary(request_iterator, target, '/hello.HelloService/HelloClientStream',
            proto_dot_hello__pb2.HelloClientStreamRequest.SerializeToString,
            proto_dot_hello__pb2.HelloClientStreamResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HelloServerStream(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(request, target, '/hello.HelloService/HelloServerStream',
            proto_dot_hello__pb2.HelloServerStreamRequest.SerializeToString,
            proto_dot_hello__pb2.HelloServerStreamResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def HelloBidirectional(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(request_iterator, target, '/hello.HelloService/HelloBidirectional',
            proto_dot_hello__pb2.HelloBidirectionalRequest.SerializeToString,
            proto_dot_hello__pb2.HelloBidirectionalResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
