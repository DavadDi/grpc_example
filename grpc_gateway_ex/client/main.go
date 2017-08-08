package main

import (
	pb "github.com/DavadDi/grpc_example/grpc_gateway_ex/proto/hello" // 引入proto包

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50001"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloSrvClient(conn)

	// 调用方法
	reqBody := new(pb.HelloReq)
	reqBody.Name = "Dave"
	r, err := c.Echo(context.Background(), reqBody)
	if err != nil {
		grpclog.Fatalln(err)
	}

	grpclog.Info(r.Msg)
}
