package main

import (
	"log"
	"net"

	pb "github.com/MehulxBuilds/grpc-basics/proto"
	"google.golang.org/grpc"
)

//define the port
const (
	port = ":8080"
)

type HelloServer struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &HelloServer{})
	log.Printf("Server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start: %v", err)
	}
}