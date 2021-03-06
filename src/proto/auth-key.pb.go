// Code generated by protoc-gen-go. DO NOT EDIT.
// source: src/proto/auth-key.proto

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

type AuthKey struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	Key                  string   `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Available            bool     `protobuf:"varint,4,opt,name=Available,proto3" json:"Available,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AuthKey) Reset()         { *m = AuthKey{} }
func (m *AuthKey) String() string { return proto.CompactTextString(m) }
func (*AuthKey) ProtoMessage()    {}
func (*AuthKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7165c3c139066a, []int{0}
}

func (m *AuthKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AuthKey.Unmarshal(m, b)
}
func (m *AuthKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AuthKey.Marshal(b, m, deterministic)
}
func (m *AuthKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthKey.Merge(m, src)
}
func (m *AuthKey) XXX_Size() int {
	return xxx_messageInfo_AuthKey.Size(m)
}
func (m *AuthKey) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthKey.DiscardUnknown(m)
}

var xxx_messageInfo_AuthKey proto.InternalMessageInfo

func (m *AuthKey) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *AuthKey) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *AuthKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *AuthKey) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

type UserAuthKeys struct {
	UserId               string     `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	Keys                 []*AuthKey `protobuf:"bytes,2,rep,name=Keys,proto3" json:"Keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UserAuthKeys) Reset()         { *m = UserAuthKeys{} }
func (m *UserAuthKeys) String() string { return proto.CompactTextString(m) }
func (*UserAuthKeys) ProtoMessage()    {}
func (*UserAuthKeys) Descriptor() ([]byte, []int) {
	return fileDescriptor_2d7165c3c139066a, []int{1}
}

func (m *UserAuthKeys) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAuthKeys.Unmarshal(m, b)
}
func (m *UserAuthKeys) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAuthKeys.Marshal(b, m, deterministic)
}
func (m *UserAuthKeys) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAuthKeys.Merge(m, src)
}
func (m *UserAuthKeys) XXX_Size() int {
	return xxx_messageInfo_UserAuthKeys.Size(m)
}
func (m *UserAuthKeys) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAuthKeys.DiscardUnknown(m)
}

var xxx_messageInfo_UserAuthKeys proto.InternalMessageInfo

func (m *UserAuthKeys) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *UserAuthKeys) GetKeys() []*AuthKey {
	if m != nil {
		return m.Keys
	}
	return nil
}

func init() {
	proto.RegisterType((*AuthKey)(nil), "proto.AuthKey")
	proto.RegisterType((*UserAuthKeys)(nil), "proto.UserAuthKeys")
}

func init() { proto.RegisterFile("src/proto/auth-key.proto", fileDescriptor_2d7165c3c139066a) }

var fileDescriptor_2d7165c3c139066a = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x2e, 0x4a, 0xd6,
	0x2f, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0xcd, 0x4e, 0xad, 0xd4, 0x03,
	0x73, 0x85, 0x58, 0xc1, 0x94, 0x52, 0x3c, 0x17, 0xbb, 0x63, 0x69, 0x49, 0x86, 0x77, 0x6a, 0xa5,
	0x10, 0x1f, 0x17, 0x93, 0x67, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x93, 0x67, 0x8a,
	0x90, 0x08, 0x17, 0x6b, 0x48, 0x66, 0x49, 0x4e, 0xaa, 0x04, 0x13, 0x58, 0x08, 0xc2, 0x11, 0x12,
	0xe0, 0x62, 0xf6, 0x4e, 0xad, 0x94, 0x60, 0x06, 0x8b, 0x81, 0x98, 0x42, 0x32, 0x5c, 0x9c, 0x8e,
	0x65, 0x89, 0x99, 0x39, 0x89, 0x49, 0x39, 0xa9, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x1c, 0x41, 0x08,
	0x01, 0x25, 0x2f, 0x2e, 0x9e, 0xd0, 0xe2, 0xd4, 0x22, 0xa8, 0x25, 0xc5, 0x42, 0x62, 0x5c, 0x6c,
	0x20, 0x3e, 0xdc, 0x26, 0x28, 0x4f, 0x48, 0x89, 0x8b, 0x05, 0x24, 0x2f, 0xc1, 0xa4, 0xc0, 0xac,
	0xc1, 0x6d, 0xc4, 0x07, 0x71, 0xa5, 0x1e, 0x54, 0x5b, 0x10, 0x58, 0x2e, 0x89, 0x0d, 0x2c, 0x68,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x37, 0x23, 0x8e, 0xa9, 0xd6, 0x00, 0x00, 0x00,
}
