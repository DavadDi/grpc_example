package main

import (
	"fmt"
	"net"

	pb "github.com/DavadDi/grpc_example/grpc_gateway_ex/proto/hello"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	Address = "127.0.0.1:50001"
)

type helloService struct{}

var HelloService = helloService{}

func (h helloService) Echo(ctx context.Context, in *pb.HelloReq) (*pb.HelloResp, error) {
	resp := new(pb.HelloResp)
	resp.Msg = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterHelloSrvServer(s, HelloService)

	grpclog.Info("Listen on " + Address)
	s.Serve(listen)
}
