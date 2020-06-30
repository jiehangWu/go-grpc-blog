package main

import (
	"fmt"
	"log"
	"net"

	pb "go-grpc-blog/proto/genproto"

	"google.golang.org/grpc"
)

const (
	defaultPort = ":3000"
)

type blogServiceServer struct {
	pb.UnimplementedBlogServiceServer
}

func main() {
	port := defaultPort

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	srv := &blogServiceServer{}
	pb.RegisterBlogServiceServer(grpcServer, srv)

	fmt.Println("Port is listening ...")

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
}
