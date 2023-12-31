// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: proto/hello.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	HelloService_Hello_FullMethodName              = "/hello.HelloService/Hello"
	HelloService_HelloClientStream_FullMethodName  = "/hello.HelloService/HelloClientStream"
	HelloService_HelloServerStream_FullMethodName  = "/hello.HelloService/HelloServerStream"
	HelloService_HelloBidirectional_FullMethodName = "/hello.HelloService/HelloBidirectional"
)

// HelloServiceClient is the client API for HelloService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HelloServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
	HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloClientStreamClient, error)
	HelloServerStream(ctx context.Context, in *HelloServerStreamRequest, opts ...grpc.CallOption) (HelloService_HelloServerStreamClient, error)
	HelloBidirectional(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloBidirectionalClient, error)
}

type helloServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHelloServiceClient(cc grpc.ClientConnInterface) HelloServiceClient {
	return &helloServiceClient{cc}
}

func (c *helloServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, HelloService_Hello_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *helloServiceClient) HelloClientStream(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[0], HelloService_HelloClientStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloClientStreamClient{stream}
	return x, nil
}

type HelloService_HelloClientStreamClient interface {
	Send(*HelloClientStreamRequest) error
	CloseAndRecv() (*HelloClientStreamResponse, error)
	grpc.ClientStream
}

type helloServiceHelloClientStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloClientStreamClient) Send(m *HelloClientStreamRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceHelloClientStreamClient) CloseAndRecv() (*HelloClientStreamResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(HelloClientStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) HelloServerStream(ctx context.Context, in *HelloServerStreamRequest, opts ...grpc.CallOption) (HelloService_HelloServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[1], HelloService_HelloServerStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type HelloService_HelloServerStreamClient interface {
	Recv() (*HelloServerStreamResponse, error)
	grpc.ClientStream
}

type helloServiceHelloServerStreamClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloServerStreamClient) Recv() (*HelloServerStreamResponse, error) {
	m := new(HelloServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *helloServiceClient) HelloBidirectional(ctx context.Context, opts ...grpc.CallOption) (HelloService_HelloBidirectionalClient, error) {
	stream, err := c.cc.NewStream(ctx, &HelloService_ServiceDesc.Streams[2], HelloService_HelloBidirectional_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &helloServiceHelloBidirectionalClient{stream}
	return x, nil
}

type HelloService_HelloBidirectionalClient interface {
	Send(*HelloBidirectionalRequest) error
	Recv() (*HelloBidirectionalResponse, error)
	grpc.ClientStream
}

type helloServiceHelloBidirectionalClient struct {
	grpc.ClientStream
}

func (x *helloServiceHelloBidirectionalClient) Send(m *HelloBidirectionalRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *helloServiceHelloBidirectionalClient) Recv() (*HelloBidirectionalResponse, error) {
	m := new(HelloBidirectionalResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloServiceServer is the server API for HelloService service.
// All implementations must embed UnimplementedHelloServiceServer
// for forward compatibility
type HelloServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
	HelloClientStream(HelloService_HelloClientStreamServer) error
	HelloServerStream(*HelloServerStreamRequest, HelloService_HelloServerStreamServer) error
	HelloBidirectional(HelloService_HelloBidirectionalServer) error
	mustEmbedUnimplementedHelloServiceServer()
}

// UnimplementedHelloServiceServer must be embedded to have forward compatible implementations.
type UnimplementedHelloServiceServer struct {
}

func (UnimplementedHelloServiceServer) Hello(context.Context, *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Hello not implemented")
}
func (UnimplementedHelloServiceServer) HelloClientStream(HelloService_HelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloClientStream not implemented")
}
func (UnimplementedHelloServiceServer) HelloServerStream(*HelloServerStreamRequest, HelloService_HelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloServerStream not implemented")
}
func (UnimplementedHelloServiceServer) HelloBidirectional(HelloService_HelloBidirectionalServer) error {
	return status.Errorf(codes.Unimplemented, "method HelloBidirectional not implemented")
}
func (UnimplementedHelloServiceServer) mustEmbedUnimplementedHelloServiceServer() {}

// UnsafeHelloServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HelloServiceServer will
// result in compilation errors.
type UnsafeHelloServiceServer interface {
	mustEmbedUnimplementedHelloServiceServer()
}

func RegisterHelloServiceServer(s grpc.ServiceRegistrar, srv HelloServiceServer) {
	s.RegisterService(&HelloService_ServiceDesc, srv)
}

func _HelloService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HelloService_Hello_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HelloService_HelloClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HelloClientStream(&helloServiceHelloClientStreamServer{stream})
}

type HelloService_HelloClientStreamServer interface {
	SendAndClose(*HelloClientStreamResponse) error
	Recv() (*HelloClientStreamRequest, error)
	grpc.ServerStream
}

type helloServiceHelloClientStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloClientStreamServer) SendAndClose(m *HelloClientStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceHelloClientStreamServer) Recv() (*HelloClientStreamRequest, error) {
	m := new(HelloClientStreamRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _HelloService_HelloServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(HelloServerStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(HelloServiceServer).HelloServerStream(m, &helloServiceHelloServerStreamServer{stream})
}

type HelloService_HelloServerStreamServer interface {
	Send(*HelloServerStreamResponse) error
	grpc.ServerStream
}

type helloServiceHelloServerStreamServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloServerStreamServer) Send(m *HelloServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _HelloService_HelloBidirectional_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(HelloServiceServer).HelloBidirectional(&helloServiceHelloBidirectionalServer{stream})
}

type HelloService_HelloBidirectionalServer interface {
	Send(*HelloBidirectionalResponse) error
	Recv() (*HelloBidirectionalRequest, error)
	grpc.ServerStream
}

type helloServiceHelloBidirectionalServer struct {
	grpc.ServerStream
}

func (x *helloServiceHelloBidirectionalServer) Send(m *HelloBidirectionalResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *helloServiceHelloBidirectionalServer) Recv() (*HelloBidirectionalRequest, error) {
	m := new(HelloBidirectionalRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// HelloService_ServiceDesc is the grpc.ServiceDesc for HelloService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HelloService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloService",
	HandlerType: (*HelloServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _HelloService_Hello_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "HelloClientStream",
			Handler:       _HelloService_HelloClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "HelloServerStream",
			Handler:       _HelloService_HelloServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "HelloBidirectional",
			Handler:       _HelloService_HelloBidirectional_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/hello.proto",
}
