// Code generated by protoc-gen-go. DO NOT EDIT.
// source: s2.proto

/*
Package s2 is a generated protocol buffer package.

It is generated from these files:
	s2.proto

It has these top-level messages:
	BHelloRequest
	BHelloResponse
*/
package s2

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

type BHelloRequest struct {
	Name   string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Rid    []string          `protobuf:"bytes,2,rep,name=rid" json:"rid,omitempty"`
	Extras map[string]string `protobuf:"bytes,3,rep,name=extras" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *BHelloRequest) Reset()                    { *m = BHelloRequest{} }
func (m *BHelloRequest) String() string            { return proto.CompactTextString(m) }
func (*BHelloRequest) ProtoMessage()               {}
func (*BHelloRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BHelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BHelloRequest) GetRid() []string {
	if m != nil {
		return m.Rid
	}
	return nil
}

func (m *BHelloRequest) GetExtras() map[string]string {
	if m != nil {
		return m.Extras
	}
	return nil
}

type BHelloResponse struct {
	Greeting string `protobuf:"bytes,1,opt,name=greeting" json:"greeting,omitempty"`
}

func (m *BHelloResponse) Reset()                    { *m = BHelloResponse{} }
func (m *BHelloResponse) String() string            { return proto.CompactTextString(m) }
func (*BHelloResponse) ProtoMessage()               {}
func (*BHelloResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *BHelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*BHelloRequest)(nil), "BHelloRequest")
	proto.RegisterType((*BHelloResponse)(nil), "BHelloResponse")
}

func init() { proto.RegisterFile("s2.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 208 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0x36, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x57, 0x5a, 0xc2, 0xc8, 0xc5, 0xeb, 0xe4, 0x91, 0x9a, 0x93, 0x93, 0x1f,
	0x94, 0x5a, 0x58, 0x9a, 0x5a, 0x5c, 0x22, 0x24, 0xc4, 0xc5, 0x92, 0x97, 0x98, 0x9b, 0x2a, 0xc1,
	0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0x66, 0x0b, 0x09, 0x70, 0x31, 0x17, 0x65, 0xa6, 0x48, 0x30,
	0x29, 0x30, 0x6b, 0x70, 0x06, 0x81, 0x98, 0x42, 0x46, 0x5c, 0x6c, 0xa9, 0x15, 0x25, 0x45, 0x89,
	0xc5, 0x12, 0xcc, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x28, 0xa6, 0xe8, 0xb9, 0x82, 0x25,
	0x5d, 0xf3, 0x4a, 0x8a, 0x2a, 0x83, 0xa0, 0x2a, 0xa5, 0x2c, 0xb9, 0xb8, 0x91, 0x84, 0x41, 0x86,
	0x66, 0xa7, 0x56, 0x42, 0xed, 0x01, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x12, 0x73, 0x4a, 0x53,
	0x25, 0x98, 0xc0, 0x62, 0x10, 0x8e, 0x15, 0x93, 0x05, 0xa3, 0x92, 0x0e, 0x17, 0x1f, 0xcc, 0xfc,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x29, 0x2e, 0x8e, 0xf4, 0xa2, 0xd4, 0xd4, 0x92, 0xcc,
	0xbc, 0x74, 0xa8, 0x11, 0x70, 0xbe, 0x91, 0x21, 0x17, 0x53, 0xb0, 0x91, 0x90, 0x36, 0x17, 0x1b,
	0x44, 0x8f, 0x10, 0x1f, 0xaa, 0xe3, 0xa4, 0xf8, 0xf5, 0x50, 0x0d, 0x53, 0x62, 0x48, 0x62, 0x03,
	0x07, 0x87, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x61, 0xa2, 0x53, 0x8d, 0x1a, 0x01, 0x00, 0x00,
}