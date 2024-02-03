# gRPC Reverse Proxy with Go

This is a sample Go program that demonstrates a simple gRPC reverse proxy setup. The code consists of two main components: a gRPC server and a gRPC reverse proxy server.
Components
1. gRPC Server

The server struct in the code represents a simple gRPC server with a SayHello method. This method responds to a client with a greeting. The server is intended to be running on port :50052, but it is currently commented out in the code.
2. gRPC Proxy Server

The proxyServer struct acts as a reverse proxy for the gRPC server. It forwards incoming gRPC requests to the actual gRPC server running on localhost:9090. The proxy server listens on port :8080.
How to Run

    Uncomment the code related to the gRPC server if you want to run it separately on port :50052.
    Connect to the gRPC server by dialing its address (localhost:9090).
    Create a gRPC proxy using the proxy.NewProxy function, passing the connection to the gRPC server.
    Start the proxy server, listening on port :8080.
    The reverse proxy is now running, forwarding requests from port :8080 to the gRPC server on localhost:9090.

Dependencies

The code uses the following external dependencies:

    github.com/mwitkow/grpc-proxy/proxy: A package for creating a gRPC proxy server.
    google.golang.org/grpc: The official gRPC Go implementation.

Run the Code

To run the code, execute the following steps:

bash

# Clone the repository
git clone https://github.com/your-username/your-repository.git
cd your-repository

# Install dependencies
go get -u github.com/mwitkow/grpc-proxy/proxy
go get -u google.golang.org/grpc

# Run the code
go run main.go

Ensure that you have Go installed on your machine.
Access the gRPC Server and Proxy

    The gRPC server (if uncommented) is accessible at localhost:50052.
    The gRPC reverse proxy is accessible at localhost:8080. It forwards requests to the gRPC server.