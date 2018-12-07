// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service/service.proto

/*
Package service is a generated protocol buffer package.

It is generated from these files:
	service/service.proto

It has these top-level messages:
	Request
	Response
*/
package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type Request struct {
	Duration string `protobuf:"bytes,1,opt,name=duration" json:"duration,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

type Response struct {
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func init() {
	proto.RegisterType((*Request)(nil), "service.Request")
	proto.RegisterType((*Response)(nil), "service.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Time service

type TimeClient interface {
	Sleep(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type timeClient struct {
	cc *grpc.ClientConn
}

func NewTimeClient(cc *grpc.ClientConn) TimeClient {
	return &timeClient{cc}
}

func (c *timeClient) Sleep(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/service.Time/Sleep", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Time service

type TimeServer interface {
	Sleep(context.Context, *Request) (*Response, error)
}

func RegisterTimeServer(s *grpc.Server, srv TimeServer) {
	s.RegisterService(&_Time_serviceDesc, srv)
}

func _Time_Sleep_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TimeServer).Sleep(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.Time/Sleep",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TimeServer).Sleep(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Time_serviceDesc = grpc.ServiceDesc{
	ServiceName: "service.Time",
	HandlerType: (*TimeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sleep",
			Handler:    _Time_Sleep_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service/service.proto",
}

func init() { proto.RegisterFile("service/service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x92, 0x2a, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x14, 0x17, 0x47, 0x4a,
	0x69, 0x51, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x9c, 0xaf,
	0xc4, 0xc5, 0xc5, 0x11, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x6a, 0x64, 0xc2, 0xc5, 0x12,
	0x92, 0x99, 0x9b, 0x2a, 0xa4, 0xc3, 0xc5, 0x1a, 0x9c, 0x93, 0x9a, 0x5a, 0x20, 0x24, 0xa0, 0x07,
	0x33, 0x1c, 0x6a, 0x94, 0x94, 0x20, 0x92, 0x08, 0x44, 0x57, 0x12, 0x1b, 0xd8, 0x62, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x67, 0x7a, 0xc8, 0x1c, 0x91, 0x00, 0x00, 0x00,
}