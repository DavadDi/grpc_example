// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service.proto

/*
Package hello is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	HelloReq
	HelloResp
*/
package hello

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HelloReq struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *HelloReq) Reset()                    { *m = HelloReq{} }
func (m *HelloReq) String() string            { return proto.CompactTextString(m) }
func (*HelloReq) ProtoMessage()               {}
func (*HelloReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HelloReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResp struct {
	Msg string `protobuf:"bytes,1,opt,name=msg" json:"msg,omitempty"`
}

func (m *HelloResp) Reset()                    { *m = HelloResp{} }
func (m *HelloResp) String() string            { return proto.CompactTextString(m) }
func (*HelloResp) ProtoMessage()               {}
func (*HelloResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *HelloResp) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloReq)(nil), "hello.HelloReq")
	proto.RegisterType((*HelloResp)(nil), "hello.HelloResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for HelloSrv service

type HelloSrvClient interface {
	Echo(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error)
}

type helloSrvClient struct {
	cc *grpc.ClientConn
}

func NewHelloSrvClient(cc *grpc.ClientConn) HelloSrvClient {
	return &helloSrvClient{cc}
}

func (c *helloSrvClient) Echo(ctx context.Context, in *HelloReq, opts ...grpc.CallOption) (*HelloResp, error) {
	out := new(HelloResp)
	err := grpc.Invoke(ctx, "/hello.HelloSrv/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HelloSrv service

type HelloSrvServer interface {
	Echo(context.Context, *HelloReq) (*HelloResp, error)
}

func RegisterHelloSrvServer(s *grpc.Server, srv HelloSrvServer) {
	s.RegisterService(&_HelloSrv_serviceDesc, srv)
}

func _HelloSrv_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HelloSrvServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.HelloSrv/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HelloSrvServer).Echo(ctx, req.(*HelloReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _HelloSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.HelloSrv",
	HandlerType: (*HelloSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _HelloSrv_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}

func init() { proto.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcd, 0x48, 0xcd, 0xc9, 0xc9,
	0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb,
	0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b, 0x86, 0x28, 0x52, 0x92, 0xe3, 0xe2, 0xf0, 0x00,
	0x29, 0x0b, 0x4a, 0x2d, 0x14, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x64, 0xb9, 0x38, 0xa1, 0xf2, 0xc5, 0x05, 0x42, 0x02, 0x5c,
	0xcc, 0xb9, 0xc5, 0xe9, 0x50, 0x79, 0x10, 0xd3, 0x28, 0x00, 0xaa, 0x3d, 0xb8, 0xa8, 0x4c, 0xc8,
	0x85, 0x8b, 0xc5, 0x35, 0x39, 0x23, 0x5f, 0x88, 0x5f, 0x0f, 0x6c, 0xb1, 0x1e, 0xcc, 0x5c, 0x29,
	0x01, 0x54, 0x81, 0xe2, 0x02, 0x25, 0xc9, 0xa6, 0xcb, 0x4f, 0x26, 0x33, 0x09, 0x2b, 0xf1, 0xe9,
	0x97, 0x19, 0xea, 0x83, 0x25, 0xf5, 0x53, 0x93, 0x33, 0xf2, 0xad, 0x18, 0xb5, 0x92, 0xd8, 0xc0,
	0xee, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x6d, 0x96, 0xef, 0x3e, 0xcd, 0x00, 0x00, 0x00,
}
