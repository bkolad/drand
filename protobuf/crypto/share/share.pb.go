// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto/share/share.proto

/*
Package share is a generated protocol buffer package.

It is generated from these files:
	crypto/share/share.proto

It has these top-level messages:
	PrivateShare
*/
package share

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import element "github.com/dedis/drand/protobuf/crypto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// PrivateShare holds a share that is private for the given participant at the
// given index.
type PrivateShare struct {
	// index of the participant
	Index uint32 `protobuf:"varint,1,opt,name=index" json:"index,omitempty"`
	// the share itself
	Share *element.Scalar `protobuf:"bytes,2,opt,name=share" json:"share,omitempty"`
}

func (m *PrivateShare) Reset()                    { *m = PrivateShare{} }
func (m *PrivateShare) String() string            { return proto.CompactTextString(m) }
func (*PrivateShare) ProtoMessage()               {}
func (*PrivateShare) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PrivateShare) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *PrivateShare) GetShare() *element.Scalar {
	if m != nil {
		return m.Share
	}
	return nil
}

func init() {
	proto.RegisterType((*PrivateShare)(nil), "share.PrivateShare")
}

func init() { proto.RegisterFile("crypto/share/share.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x48, 0x2e, 0xaa, 0x2c,
	0x28, 0xc9, 0xd7, 0x2f, 0xce, 0x48, 0x2c, 0x4a, 0x85, 0x90, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xac, 0x60, 0x8e, 0x94, 0x08, 0x54, 0x41, 0x6a, 0x4e, 0x6a, 0x6e, 0x6a, 0x5e, 0x09, 0x44,
	0x52, 0xc9, 0x9b, 0x8b, 0x27, 0xa0, 0x28, 0xb3, 0x2c, 0xb1, 0x24, 0x35, 0x18, 0xa4, 0x4a, 0x48,
	0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51, 0x83, 0x37, 0x08, 0xc2,
	0x11, 0x52, 0xe5, 0x82, 0x18, 0x22, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x6d, 0xc4, 0xaf, 0x07, 0x33,
	0x24, 0x38, 0x39, 0x31, 0x27, 0xb1, 0x28, 0x08, 0x22, 0xeb, 0xa4, 0x17, 0xa5, 0x93, 0x9e, 0x59,
	0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x92, 0x9a, 0x92, 0x59, 0xac, 0x9f, 0x52,
	0x94, 0x98, 0x97, 0xa2, 0x0f, 0xb6, 0x2c, 0xa9, 0x34, 0x4d, 0x1f, 0xd9, 0x95, 0x49, 0x6c, 0x60,
	0x61, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x27, 0x2b, 0x59, 0x48, 0xbc, 0x00, 0x00, 0x00,
}
