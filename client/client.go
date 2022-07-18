// copyright (c) 2022, tencent inc. all rights reserved.
// author: jimjqzhang (jimjqzhang@tencent.com)
// filename: client.go
// create at: 2022/7/18 17:11

// Package client : some package info
package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	pb "github/HermesNamco/GrpcTest/protocol"
)

func main() {
	serverAddr := "localhost:23333"
	conn, err := grpc.DialContext(context.TODO(), serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error connect, %v", err)
	}
	defer conn.Close()

	client := pb.NewHelloServiceClient(conn)
	rsp, err := client.Hello(context.TODO(), &pb.HelloReq{Code: 1, Msg: "Hello server"})
	fmt.Println(rsp)
}
