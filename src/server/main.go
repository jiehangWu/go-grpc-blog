package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/jiehangWu/go-grpc-blog/proto/genproto"

	"google.golang.org/grpc"
)

const (
	defaultPort = "5000"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	port := defaultPort
	grpcServer := grpc.NewServer()
	pb.RegisterBlogServiceServer(grpcServer, newServer())
	grpcServer.Serve(lis)
}
