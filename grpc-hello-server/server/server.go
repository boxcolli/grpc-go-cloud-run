package server

import (
	"context"
	"io"
	"grpc-hello-server/pb"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	history []string
	pb.UnimplementedHelloServiceServer
}

func NewServer() pb.HelloServiceServer {
	return &server{}
}

func getHelloSentence(name string) string {
	return "Hello, " + strings.TrimSpace(name)
}

// Hello implements pb.HelloServiceServer.
func (s *server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	// Check name argument
	name := req.GetName()
	if name == "" {
		return nil, status.Error(codes.InvalidArgument, "Name argument is empty")
	}

	// Process new name
	hello := getHelloSentence(name)
	s.history = append(s.history, hello)
	return &pb.HelloResponse{
		Hello: hello,
	}, nil
}

// HelloClientStream implements pb.HelloServiceServer.
func (s *server) HelloClientStream(stream pb.HelloService_HelloClientStreamServer) error {
	// Collect names
	names := make([]string, 1)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		name := req.GetName()
		if name == "" {
			return status.Error(codes.InvalidArgument, "There is empty name in the argument")
		}
		names = append(names, req.GetName())
	}

	// Process new names
	for _, name := range names {
		s.history = append(s.history, getHelloSentence(name))
	}

	// Send all hello
	stream.SendAndClose(&pb.HelloClientStreamResponse{HelloList: s.history})

	return nil
}

// HelloServerStream implements pb.HelloServiceServer.
func (s *server) HelloServerStream(req *pb.HelloServerStreamRequest, stream pb.HelloService_HelloServerStreamServer) error {
	// Check name argument
	name := req.GetName()
	if name == "" {
		return status.Error(codes.InvalidArgument, "Name argument is empty")
	}

	// Append new hello
	s.history = append(s.history, getHelloSentence(name))

	// Send all hello
	for _, hello := range s.history{
		if err := stream.Send(&pb.HelloServerStreamResponse{Hello: hello}); err != nil {
			return err
		}
	}

	return nil
}

// HelloBidirectional implements pb.HelloServiceServer.
func (s *server) HelloBidirectional(stream pb.HelloService_HelloBidirectionalServer) error {
	for {
		// Catch new name
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		// Send back hello
		hello := getHelloSentence(req.GetName())
		stream.Send(&pb.HelloBidirectionalResponse{Hello: hello})

		// Keep the new hello
		s.history = append(s.history, hello)
	}
}
