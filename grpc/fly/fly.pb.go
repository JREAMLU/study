// Code generated by protoc-gen-go.
// source: fly.proto
// DO NOT EDIT!

/*
Package fly is a generated protocol buffer package.

It is generated from these files:
	fly.proto

It has these top-level messages:
	FlyRequest
	FlyReply
*/
package fly

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

type FlyRequest struct {
	Version string `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
}

func (m *FlyRequest) Reset()                    { *m = FlyRequest{} }
func (m *FlyRequest) String() string            { return proto.CompactTextString(m) }
func (*FlyRequest) ProtoMessage()               {}
func (*FlyRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *FlyRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

type FlyReply struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *FlyReply) Reset()                    { *m = FlyReply{} }
func (m *FlyReply) String() string            { return proto.CompactTextString(m) }
func (*FlyReply) ProtoMessage()               {}
func (*FlyReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *FlyReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*FlyRequest)(nil), "fly.FlyRequest")
	proto.RegisterType((*FlyReply)(nil), "fly.FlyReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for FlyToSky service

type FlyToSkyClient interface {
	DriverPlane(ctx context.Context, in *FlyRequest, opts ...grpc.CallOption) (*FlyReply, error)
}

type flyToSkyClient struct {
	cc *grpc.ClientConn
}

func NewFlyToSkyClient(cc *grpc.ClientConn) FlyToSkyClient {
	return &flyToSkyClient{cc}
}

func (c *flyToSkyClient) DriverPlane(ctx context.Context, in *FlyRequest, opts ...grpc.CallOption) (*FlyReply, error) {
	out := new(FlyReply)
	err := grpc.Invoke(ctx, "/fly.FlyToSky/DriverPlane", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for FlyToSky service

type FlyToSkyServer interface {
	DriverPlane(context.Context, *FlyRequest) (*FlyReply, error)
}

func RegisterFlyToSkyServer(s *grpc.Server, srv FlyToSkyServer) {
	s.RegisterService(&_FlyToSky_serviceDesc, srv)
}

func _FlyToSky_DriverPlane_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FlyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlyToSkyServer).DriverPlane(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/fly.FlyToSky/DriverPlane",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlyToSkyServer).DriverPlane(ctx, req.(*FlyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FlyToSky_serviceDesc = grpc.ServiceDesc{
	ServiceName: "fly.FlyToSky",
	HandlerType: (*FlyToSkyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DriverPlane",
			Handler:    _FlyToSky_DriverPlane_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fly.proto",
}

func init() { proto.RegisterFile("fly.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0xcb, 0xa9, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0xcb, 0xa9, 0x54, 0x52, 0xe3, 0xe2, 0x72, 0xcb,
	0xa9, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe0, 0x62, 0x2f, 0x4b, 0x2d, 0x2a,
	0xce, 0xcc, 0xcf, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x95, 0xe4, 0xb8, 0x38,
	0xc0, 0xea, 0x0a, 0x72, 0x2a, 0x85, 0x84, 0xb8, 0x58, 0xf2, 0x12, 0x73, 0x53, 0xa1, 0x4a, 0xc0,
	0x6c, 0x23, 0x6b, 0xb0, 0x7c, 0x48, 0x7e, 0x70, 0x76, 0xa5, 0x90, 0x3e, 0x17, 0xb7, 0x4b, 0x51,
	0x66, 0x59, 0x6a, 0x51, 0x40, 0x4e, 0x62, 0x5e, 0xaa, 0x10, 0xbf, 0x1e, 0xc8, 0x4e, 0x84, 0x2d,
	0x52, 0xbc, 0x08, 0x81, 0x82, 0x9c, 0x4a, 0x25, 0x06, 0x27, 0x03, 0x2e, 0xe9, 0xcc, 0x7c, 0xbd,
	0xf4, 0xa2, 0x82, 0x64, 0xbd, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c, 0xd4, 0x62, 0xbd, 0xa2, 0xfc,
	0xd2, 0x92, 0xd4, 0xf4, 0xd2, 0xcc, 0x94, 0x54, 0x27, 0xfe, 0x20, 0x10, 0xdb, 0x1d, 0xc4, 0x0e,
	0x00, 0xb9, 0x3c, 0x80, 0x31, 0x89, 0x0d, 0xec, 0x05, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x5b, 0x30, 0x1b, 0xa2, 0xcf, 0x00, 0x00, 0x00,
}