package main

import (
	"bufio"
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"grpc-hello-server-client/pb"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/peterbourgon/ff/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)



var client pb.HelloServiceClient
var reader *bufio.Reader

func main() {
	fs := flag.NewFlagSet("grpc-hello-client", flag.ContinueOnError)
	var (
		serverAddr = fs.String("server-addr", "127.0.0.1:50051", "The server port")
		rpc = fs.Int("rpc", 0, "0=Hello, 1=HelloServerStream, 2=HelloClientStream 3=HelloBidirectional")
	)
	err := ff.Parse(fs, os.Args[1:], ff.WithEnvVars())
	if err != nil {
		log.Fatal(err)
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(*serverAddr))
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{InsecureSkipVerify: false})))

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Lshortfile)
	logopts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
		// Add any other option (check functions starting with logging.With).
	}
	opts = append(
		opts,
		grpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(InterceptorLogger(logger), logopts...),
			// Add any other interceptor you want.
		),
		grpc.WithChainStreamInterceptor(
			logging.StreamClientInterceptor(InterceptorLogger(logger), logopts...),
			// Add any other interceptor you want.
		),
	)

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client = pb.NewHelloServiceClient(conn)
	reader = bufio.NewReader(os.Stdin)

	var rpcErr error = nil
	var clientErr error = nil
	switch r := *rpc; r {
	case 0: rpcErr = hello()
	case 1: rpcErr = helloServerStream()
	case 2: rpcErr = helloClientStream()
	case 3: rpcErr, clientErr = helloBidirectional()
	default: break
	}
	fmt.Println("rpcErr: ", rpcErr)
	fmt.Println("clientErr: ", clientErr)
}

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

func readName() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func getCtx() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	md := metadata.New(map[string]string{"Content-Type":"application/grpc"})
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx, cancel
}

func hello() error {
	fmt.Print("(Hello) Enter name: ")
	text := readName()
	ctx, cancel := getCtx()
	defer cancel()

	res, err := client.Hello(ctx, &pb.HelloRequest{Name: text})
	if err != nil {
		return err
	}
	println("Hello: ", res.GetHello())

	return nil
}

func helloServerStream() error {
	// Get user input
	fmt.Print("(HelloServerStream) Enter name: ")
	text := readName()
	ctx, cancel := getCtx()
	defer cancel()

	// Send name
	stream, err := client.HelloServerStream(ctx, &pb.HelloServerStreamRequest{Name: text})
	if err != nil {
		return err
	}

	// Listen server stream
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return err
		}
		if err != nil {
			return err
		}

		fmt.Println("HelloServerStream: ", res.GetHello())
	}
}

func helloClientStream() error {
	// Open stream
	ctx, cancel := getCtx()
	defer cancel()
	stream, err := client.HelloClientStream(ctx)
	if err != nil {
		return err
	}

	for {
		// Get user input
		fmt.Print("(HelloClientStream) Enter name (q for exit): ")
		text := readName()
		if text == "q" {
			break
		}

		// Send name
		err := stream.Send(&pb.HelloClientStreamRequest{Name: text})
		if err != nil {
			return err
		}
	}

	// Send close
	res, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	// Print response
	for _, hello := range res.GetHelloList() {
		fmt.Println("HelloClientStream: ", hello)
	}

	return nil
}

func helloBidirectional() (serverErr error, clientErr error) {
	// Open stream
	ctx, cancel := getCtx()
	defer cancel()
	stream, err := client.HelloBidirectional(ctx)
	if err != nil {
		return nil, err
	}

	// Listen server stream
	waitc := make(chan struct{})
	go func() {
		for {
			// Get response
			res, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("HelloBidirectional: server closed gracefully")
				close(waitc)
				serverErr = err
				return
			}
			if err != nil {
				serverErr = err
				return
			}

			// Print response
			fmt.Println("HelloBidirectional: ", res.GetHello())
		}
	} ()

	// Listen user input
	for {
		// Get user input
		fmt.Print("(HelloBidirectional) Enter name (q for exit): ")
		text := readName()
		if text == "q" {
			break
		}

		// Send name
		err := stream.Send(&pb.HelloBidirectionalRequest{Name: text})
		if err != nil {
			clientErr = err
		}
	}

	stream.CloseSend()
	<-waitc

	return
}