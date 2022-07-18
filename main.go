// copyright (c) 2022, tencent inc. all rights reserved.
// author: jimjqzhang (jimjqzhang@tencent.com)
// filename: main.go
// create at: 2022/7/18 16:01

// Package GrpcTest : some package info
package main

import (
	"fmt"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	pb "github/HermesNamco/GrpcTest/protocol"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 23333))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterHelloServiceServer(grpcServer, &HelloImpl{})
	grpcServer.Serve(lis)
}