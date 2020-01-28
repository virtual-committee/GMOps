// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/proto/git-hook.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Hook struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Source               string   `protobuf:"bytes,3,opt,name=Source,proto3" json:"Source,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hook) Reset()         { *m = Hook{} }
func (m *Hook) String() string { return proto.CompactTextString(m) }
func (*Hook) ProtoMessage()    {}
func (*Hook) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7501c3528381486, []int{0}
}

func (m *Hook) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hook.Unmarshal(m, b)
}
func (m *Hook) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hook.Marshal(b, m, deterministic)
}
func (m *Hook) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hook.Merge(m, src)
}
func (m *Hook) XXX_Size() int {
	return xxx_messageInfo_Hook.Size(m)
}
func (m *Hook) XXX_DiscardUnknown() {
	xxx_messageInfo_Hook.DiscardUnknown(m)
}

var xxx_messageInfo_Hook proto.InternalMessageInfo

func (m *Hook) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Hook) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Hook) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

type Hooks struct {
	Type                 string   `protobuf:"bytes,1,opt,name=Type,proto3" json:"Type,omitempty"`
	Hooks                []*Hook  `protobuf:"bytes,2,rep,name=Hooks,proto3" json:"Hooks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hooks) Reset()         { *m = Hooks{} }
func (m *Hooks) String() string { return proto.CompactTextString(m) }
func (*Hooks) ProtoMessage()    {}
func (*Hooks) Descriptor() ([]byte, []int) {
	return fileDescriptor_e7501c3528381486, []int{1}
}

func (m *Hooks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hooks.Unmarshal(m, b)
}
func (m *Hooks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hooks.Marshal(b, m, deterministic)
}
func (m *Hooks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hooks.Merge(m, src)
}
func (m *Hooks) XXX_Size() int {
	return xxx_messageInfo_Hooks.Size(m)
}
func (m *Hooks) XXX_DiscardUnknown() {
	xxx_messageInfo_Hooks.DiscardUnknown(m)
}

var xxx_messageInfo_Hooks proto.InternalMessageInfo

func (m *Hooks) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Hooks) GetHooks() []*Hook {
	if m != nil {
		return m.Hooks
	}
	return nil
}

func init() {
	proto.RegisterType((*Hook)(nil), "proto.Hook")
	proto.RegisterType((*Hooks)(nil), "proto.Hooks")
}

func init() { proto.RegisterFile("src/proto/git-hook.proto", fileDescriptor_e7501c3528381486) }

var fileDescriptor_e7501c3528381486 = []byte{
	// 145 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0xcf, 0x2c, 0xd1, 0xcd, 0xc8, 0xcf, 0xcf, 0xd6, 0x03,
	0x73, 0x85, 0x58, 0xc1, 0x94, 0x92, 0x13, 0x17, 0x8b, 0x47, 0x7e, 0x7e, 0xb6, 0x10, 0x1f, 0x17,
	0x93, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x93, 0x67, 0x8a, 0x90, 0x10, 0x17,
	0x8b, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x04, 0xcc, 0x16, 0x12, 0xe3, 0x62, 0x0b, 0xce,
	0x2f, 0x2d, 0x4a, 0x4e, 0x95, 0x60, 0x06, 0x8b, 0x42, 0x79, 0x4a, 0x76, 0x5c, 0xac, 0x20, 0x33,
	0x8a, 0x41, 0x9a, 0x42, 0x2a, 0x0b, 0x52, 0xa1, 0xc6, 0x80, 0xd9, 0x42, 0x8a, 0x50, 0x49, 0x09,
	0x26, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x6e, 0x88, 0xf5, 0x7a, 0x20, 0xb1, 0x20, 0x88, 0x4c, 0x12,
	0x1b, 0x58, 0xc8, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x89, 0xfc, 0xdf, 0xad, 0x00, 0x00,
	0x00,
}