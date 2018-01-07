// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	sample.proto

It has these top-level messages:
	StringMessage
*/
package pb

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

type StringMessage struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *StringMessage) Reset()                    { *m = StringMessage{} }
func (m *StringMessage) String() string            { return proto.CompactTextString(m) }
func (*StringMessage) ProtoMessage()               {}
func (*StringMessage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StringMessage) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*StringMessage)(nil), "StringMessage")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Sample service

type SampleClient interface {
	Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error)
}

type sampleClient struct {
	cc *grpc.ClientConn
}

func NewSampleClient(cc *grpc.ClientConn) SampleClient {
	return &sampleClient{cc}
}

func (c *sampleClient) Echo(ctx context.Context, in *StringMessage, opts ...grpc.CallOption) (*StringMessage, error) {
	out := new(StringMessage)
	err := grpc.Invoke(ctx, "/Sample/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Sample service

type SampleServer interface {
	Echo(context.Context, *StringMessage) (*StringMessage, error)
}

func RegisterSampleServer(s *grpc.Server, srv SampleServer) {
	s.RegisterService(&_Sample_serviceDesc, srv)
}

func _Sample_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StringMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SampleServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Sample/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SampleServer).Echo(ctx, req.(*StringMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Sample_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Sample",
	HandlerType: (*SampleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Sample_Echo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample.proto",
}

func init() { proto.RegisterFile("sample.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x97, 0x92, 0x49, 0xcf, 0xcf, 0x4f, 0xcf, 0x49,
	0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x4f, 0xcc, 0xcb, 0xcb, 0x2f, 0x49, 0x2c, 0xc9, 0xcc, 0xcf, 0x2b,
	0x86, 0xc8, 0x2a, 0xa9, 0x72, 0xf1, 0x06, 0x97, 0x14, 0x65, 0xe6, 0xa5, 0xfb, 0xa6, 0x16, 0x17,
	0x27, 0xa6, 0xa7, 0x0a, 0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0x41, 0x38, 0x46, 0x4e, 0x5c, 0x6c, 0xc1, 0x60, 0x43, 0x85, 0x2c, 0xb8, 0x58,
	0x5c, 0x93, 0x33, 0xf2, 0x85, 0xf8, 0xf4, 0x50, 0xf4, 0x49, 0xa1, 0xf1, 0x95, 0x04, 0x9a, 0x2e,
	0x3f, 0x99, 0xcc, 0xc4, 0xa5, 0xc4, 0xaa, 0x9f, 0x9a, 0x9c, 0x91, 0x6f, 0xc5, 0xa8, 0xe5, 0xc4,
	0x12, 0xc5, 0x54, 0x90, 0x94, 0xc4, 0x06, 0xb6, 0xd7, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x92, 0xca, 0x58, 0xa5, 0x00, 0x00, 0x00,
}
