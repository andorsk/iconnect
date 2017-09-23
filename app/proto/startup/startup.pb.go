// Code generated by protoc-gen-go.
// source: startup/startup.proto
// DO NOT EDIT!

/*
Package startup is a generated protocol buffer package.

It is generated from these files:
	startup/startup.proto

It has these top-level messages:
	Startup
*/
package startup

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import location "location"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Startup struct {
	Name        string             `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Id          string             `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Email       string             `protobuf:"bytes,3,opt,name=email" json:"email,omitempty"`
	Description string             `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	Shortdesc   string             `protobuf:"bytes,5,opt,name=shortdesc" json:"shortdesc,omitempty"`
	Location    *location.Location `protobuf:"bytes,6,opt,name=location" json:"location,omitempty"`
}

func (m *Startup) Reset()                    { *m = Startup{} }
func (m *Startup) String() string            { return proto.CompactTextString(m) }
func (*Startup) ProtoMessage()               {}
func (*Startup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Startup) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Startup) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Startup) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Startup) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Startup) GetShortdesc() string {
	if m != nil {
		return m.Shortdesc
	}
	return ""
}

func (m *Startup) GetLocation() *location.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*Startup)(nil), "startup.Startup")
}

func init() { proto.RegisterFile("startup/startup.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x2e, 0x49, 0x2c,
	0x2a, 0x29, 0x2d, 0xd0, 0x87, 0xd2, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae,
	0x94, 0x78, 0x4e, 0x7e, 0x72, 0x62, 0x49, 0x66, 0x7e, 0x9e, 0x3e, 0x8c, 0x01, 0x51, 0xa1, 0xb4,
	0x99, 0x91, 0x8b, 0x3d, 0x18, 0xa2, 0x48, 0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82,
	0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xcc, 0x16, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x02,
	0x8b, 0x30, 0x65, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0xa6, 0xe6, 0x26, 0x66, 0xe6, 0x48, 0x30, 0x83,
	0x85, 0x20, 0x1c, 0x21, 0x05, 0x2e, 0xee, 0x94, 0xd4, 0xe2, 0xe4, 0xa2, 0xcc, 0x02, 0x90, 0xd1,
	0x12, 0x2c, 0x60, 0x39, 0x64, 0x21, 0x21, 0x19, 0x2e, 0xce, 0xe2, 0x8c, 0xfc, 0xa2, 0x12, 0x90,
	0x98, 0x04, 0x2b, 0x58, 0x1e, 0x21, 0x20, 0xa4, 0xc7, 0xc5, 0x01, 0x73, 0x97, 0x04, 0x9b, 0x02,
	0xa3, 0x06, 0xb7, 0x91, 0x90, 0x1e, 0xdc, 0xa1, 0x3e, 0x50, 0x46, 0x10, 0x5c, 0x4d, 0x12, 0x1b,
	0xd8, 0xf1, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x13, 0x28, 0x2a, 0xb8, 0xf7, 0x00, 0x00,
	0x00,
}
