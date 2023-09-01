package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"grpc-hello-server/pb"
	"grpc-hello-server/server"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/peterbourgon/ff/v3"
	"google.golang.org/grpc"
)

func main() {
	fs := flag.NewFlagSet("grpc-hello-server", flag.ContinueOnError)
	var (
		port = fs.String("port", "50051", "The server port")
	)
	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars())
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(InterceptorLogger(logger), opts...),
		),
		grpc.ChainStreamInterceptor(
			logging.StreamServerInterceptor(InterceptorLogger(logger), opts...),
		),
	)
	pb.RegisterHelloServiceServer(grpcServer, server.NewServer())
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// InterceptorLogger adapts standard Go logger to interceptor logger.
// This code is simple enough to be copied and not imported.
func InterceptorLogger(l *log.Logger) logging.Logger {
	return logging.LoggerFunc(func(_ context.Context, lvl logging.Level, msg string, fields ...any) {
		switch lvl {
		case logging.LevelDebug:
			msg = fmt.Sprintf("DEBUG :%v", msg)
		case logging.LevelInfo:
			msg = fmt.Sprintf("INFO :%v", msg)
		case logging.LevelWarn:
			msg = fmt.Sprintf("WARN :%v", msg)
		case logging.LevelError:
			msg = fmt.Sprintf("ERROR :%v", msg)
		default:
			panic(fmt.Sprintf("unknown level %v", lvl))
		}
		l.Println(append([]any{"msg", msg}, fields...))
	})
}