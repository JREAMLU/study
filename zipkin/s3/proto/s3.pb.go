// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s3.proto

/*
Package s3 is a generated protocol buffer package.

It is generated from these files:
	s3.proto

It has these top-level messages:
	CHelloRequest
	CHelloResponse
*/
package s3

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CHelloRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *CHelloRequest) Reset()                    { *m = CHelloRequest{} }
func (m *CHelloRequest) String() string            { return proto.CompactTextString(m) }
func (*CHelloRequest) ProtoMessage()               {}
func (*CHelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CHelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CHelloResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *CHelloResponse) Reset()                    { *m = CHelloResponse{} }
func (m *CHelloResponse) String() string            { return proto.CompactTextString(m) }
func (*CHelloResponse) ProtoMessage()               {}
func (*CHelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CHelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*CHelloRequest)(nil), "CHelloRequest")
	proto.RegisterType((*CHelloResponse)(nil), "CHelloResponse")
}

func init() { proto.RegisterFile("s3.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 125 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x36, 0xd6, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x52, 0xe6, 0xe2, 0x75, 0xf6, 0x48, 0xcd, 0xc9, 0xc9, 0x0f, 0x4a,
	0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x54,
	0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x74, 0xb8, 0xf8, 0x60, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3,
	0x8a, 0x53, 0x85, 0xa4, 0xb8, 0x38, 0xd2, 0x8b, 0x52, 0x53, 0x4b, 0x32, 0xf3, 0xd2, 0xa1, 0x2a,
	0xe1, 0x7c, 0x23, 0x43, 0x2e, 0xa6, 0x60, 0x63, 0x21, 0x6d, 0x2e, 0x36, 0x88, 0x1e, 0x21, 0x3e,
	0x3d, 0x14, 0x1b, 0xa4, 0xf8, 0xf5, 0x50, 0x0d, 0x53, 0x62, 0x48, 0x62, 0x03, 0x3b, 0xc6, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0x46, 0x4d, 0x72, 0x07, 0x98, 0x00, 0x00, 0x00,
}
