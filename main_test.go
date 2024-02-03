package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-reverse-proxy/proto/helloworld"
	"log"
	"testing"
	"time"
)

func TestIntegration(t *testing.T) {

	go func() {
		main()
	}()
	time.Sleep(5 * time.Second)
	// Wait for both gRPC server and reverse proxy to be ready

	// Dial the gRPC proxy
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial gRPC server: %v", err)
	}
	log.Println("connected to reverse proxy")
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	// Send a test gRPC request
	req := &pb.HelloRequest{Name: "Alice"}
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify the response
	if resp.Message != "Hello Alice" {
		t.Errorf("unexpected response: %v", resp)
	}
}
