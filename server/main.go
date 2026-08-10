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

