package main

import (
	"context"
	"github.com/mwitkow/grpc-proxy/proxy"
	"google.golang.org/grpc"
	pb "grpc-reverse-proxy/proto/helloworld"
	"log"
	"net"
)

// Dummy gRPC server implementation
type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("in server SayHello")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

// Proxy server to forward gRPC requests to the gRPC server
type proxyServer struct {
	client pb.GreeterClient
	pb.UnimplementedGreeterServer
}

func (s *proxyServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Println("in proxy SayHello")
	resp, err := s.client.SayHello(ctx, in)
	if err == nil {
		// Log the size of the response
		log.Printf("Response size: %d bytes", len(resp.Message))
	}
	return resp, err
}

func main() {
	//// Start gRPC server
	//listener, err := net.Listen("tcp", ":50052")
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//
	//grpcServer := grpc.NewServer()
	//pb.RegisterGreeterServer(grpcServer, &server{})
	//
	//go func() {
	//	log.Println("gRPC server is running on :50052")
	//	if err := grpcServer.Serve(listener); err != nil {
	//		log.Fatalf("failed to serve: %v", err)
	//	}
	//}()

	// connect to grpc server

	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
	}
	defer conn.Close()
	log.Println("connected to server :9090")
	// Register the proxy handler that forwards requests to the backend gRPC server
	proxyNew := proxy.NewProxy(conn)

	// Create a listener for the proxy server
	proxyListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Start the gRPC proxy server in a goroutine
	go func() {
		if err := proxyNew.Serve(proxyListener); err != nil {
			log.Fatalf("proxy server failed to serve: %v", err)
		}
	}()

	log.Println("Reverse proxy is running on :8080, forwarding requests to :9090")

	// Run indefinitely
	select {}
}
