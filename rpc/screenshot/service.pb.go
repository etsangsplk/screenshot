// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rpc/screenshot/service.proto

package screenshot

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RequestImage struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestImage) Reset()         { *m = RequestImage{} }
func (m *RequestImage) String() string { return proto.CompactTextString(m) }
func (*RequestImage) ProtoMessage()    {}
func (*RequestImage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecff0a2cf2aeb787, []int{0}
}

func (m *RequestImage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestImage.Unmarshal(m, b)
}
func (m *RequestImage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestImage.Marshal(b, m, deterministic)
}
func (m *RequestImage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestImage.Merge(m, src)
}
func (m *RequestImage) XXX_Size() int {
	return xxx_messageInfo_RequestImage.Size(m)
}
func (m *RequestImage) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestImage.DiscardUnknown(m)
}

var xxx_messageInfo_RequestImage proto.InternalMessageInfo

func (m *RequestImage) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Resp struct {
	Resp                 []byte   `protobuf:"bytes,1,opt,name=resp,proto3" json:"resp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resp) Reset()         { *m = Resp{} }
func (m *Resp) String() string { return proto.CompactTextString(m) }
func (*Resp) ProtoMessage()    {}
func (*Resp) Descriptor() ([]byte, []int) {
	return fileDescriptor_ecff0a2cf2aeb787, []int{1}
}

func (m *Resp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resp.Unmarshal(m, b)
}
func (m *Resp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resp.Marshal(b, m, deterministic)
}
func (m *Resp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resp.Merge(m, src)
}
func (m *Resp) XXX_Size() int {
	return xxx_messageInfo_Resp.Size(m)
}
func (m *Resp) XXX_DiscardUnknown() {
	xxx_messageInfo_Resp.DiscardUnknown(m)
}

var xxx_messageInfo_Resp proto.InternalMessageInfo

func (m *Resp) GetResp() []byte {
	if m != nil {
		return m.Resp
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestImage)(nil), "pressly.RequestImage")
	proto.RegisterType((*Resp)(nil), "pressly.Resp")
}

func init() { proto.RegisterFile("rpc/screenshot/service.proto", fileDescriptor_ecff0a2cf2aeb787) }

var fileDescriptor_ecff0a2cf2aeb787 = []byte{
	// 164 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2a, 0x48, 0xd6,
	0x2f, 0x4e, 0x2e, 0x4a, 0x4d, 0xcd, 0x2b, 0xce, 0xc8, 0x2f, 0xd1, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb,
	0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2f, 0x28, 0x4a, 0x2d, 0x2e, 0xce,
	0xa9, 0x54, 0x52, 0xe0, 0xe2, 0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0xf1, 0xcc, 0x4d, 0x4c,
	0x4f, 0x15, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02,
	0x31, 0x95, 0xa4, 0xb8, 0x58, 0x82, 0x52, 0x8b, 0x0b, 0x84, 0x84, 0xb8, 0x58, 0x8a, 0x52, 0x8b,
	0x0b, 0xc0, 0x52, 0x3c, 0x41, 0x60, 0xb6, 0x51, 0x06, 0x17, 0x57, 0x30, 0xdc, 0x0a, 0x21, 0x5d,
	0x2e, 0x56, 0x88, 0x21, 0xa2, 0x7a, 0x50, 0xe3, 0xf5, 0x90, 0xcd, 0x96, 0xe2, 0x45, 0x12, 0x2e,
	0x2e, 0x10, 0xd2, 0xe6, 0x62, 0x0e, 0x70, 0x71, 0x23, 0x4e, 0xb1, 0x13, 0x4f, 0x14, 0x17, 0xc2,
	0x33, 0x49, 0x6c, 0x60, 0x5f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x1b, 0xda, 0xe2, 0x06,
	0xe5, 0x00, 0x00, 0x00,
}
