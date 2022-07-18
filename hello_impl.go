// copyright (c) 2022, tencent inc. all rights reserved.
// author: jimjqzhang (jimjqzhang@tencent.com)
// filename: hello_impl.go
// create at: 2022/7/18 16:02

// Package GrpcTest : some package info
package main

import (
	"context"
	"fmt"

	pb "github/HermesNamco/GrpcTest/protocol"
)

type HelloImpl struct {
	// 不使用向前兼容，希望在新增rpc方法却没有显示实现时，编译器会报错
	// 如果使用向前兼容，则嵌入pb.UnimplementedHelloServiceServer，新增rpc方法时却没有显示实现时，该结构体提供了默认实现，不会报错
	pb.UnsafeHelloServiceServer
}

func(impl *HelloImpl)Hello(ctx context.Context, req *pb.HelloReq) (*pb.HelloRsp, error) {
	fmt.Println(req.Code, req.Msg)
	return &pb.HelloRsp{Msg: "Hello"}, nil
}
