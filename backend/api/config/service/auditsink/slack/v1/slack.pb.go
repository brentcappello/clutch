// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config/service/auditsink/slack/v1/slack.proto

package slackv1

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	v1 "github.com/lyft/clutch/backend/api/config/service/audit/v1"
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

type SlackConfig struct {
	Token                string     `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Channel              string     `protobuf:"bytes,2,opt,name=channel,proto3" json:"channel,omitempty"`
	Filter               *v1.Filter `protobuf:"bytes,3,opt,name=filter,proto3" json:"filter,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SlackConfig) Reset()         { *m = SlackConfig{} }
func (m *SlackConfig) String() string { return proto.CompactTextString(m) }
func (*SlackConfig) ProtoMessage()    {}
func (*SlackConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_14ed03503fa9dce3, []int{0}
}

func (m *SlackConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SlackConfig.Unmarshal(m, b)
}
func (m *SlackConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SlackConfig.Marshal(b, m, deterministic)
}
func (m *SlackConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SlackConfig.Merge(m, src)
}
func (m *SlackConfig) XXX_Size() int {
	return xxx_messageInfo_SlackConfig.Size(m)
}
func (m *SlackConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_SlackConfig.DiscardUnknown(m)
}

var xxx_messageInfo_SlackConfig proto.InternalMessageInfo

func (m *SlackConfig) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *SlackConfig) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

func (m *SlackConfig) GetFilter() *v1.Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func init() {
	proto.RegisterType((*SlackConfig)(nil), "clutch.config.service.auditsink.slack.v1.SlackConfig")
}

func init() {
	proto.RegisterFile("config/service/auditsink/slack/v1/slack.proto", fileDescriptor_14ed03503fa9dce3)
}

var fileDescriptor_14ed03503fa9dce3 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4d, 0xce, 0xcf, 0x4b,
	0xcb, 0x4c, 0xd7, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x4f, 0x2c, 0x4d, 0xc9, 0x2c,
	0x29, 0xce, 0xcc, 0xcb, 0xd6, 0x2f, 0xce, 0x49, 0x4c, 0xce, 0xd6, 0x2f, 0x33, 0x84, 0x30, 0xf4,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x34, 0x92, 0x73, 0x4a, 0x4b, 0x92, 0x33, 0xf4, 0x20, 0xba,
	0xf4, 0xa0, 0xba, 0xf4, 0xe0, 0xba, 0xf4, 0x20, 0x8a, 0xcb, 0x0c, 0xa5, 0x94, 0xb1, 0x19, 0x0c,
	0x32, 0x0e, 0xcc, 0x80, 0x18, 0x27, 0x25, 0x5e, 0x96, 0x98, 0x93, 0x99, 0x92, 0x58, 0x92, 0xaa,
	0x0f, 0x63, 0x40, 0x24, 0x94, 0xfa, 0x19, 0xb9, 0xb8, 0x83, 0x41, 0x46, 0x39, 0x83, 0x4d, 0x11,
	0x92, 0xe5, 0x62, 0x2d, 0xc9, 0xcf, 0x4e, 0xcd, 0x93, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x74, 0x62,
	0xff, 0xe5, 0xc4, 0x52, 0xc4, 0xa4, 0xc0, 0x18, 0x04, 0x11, 0x15, 0x52, 0xe4, 0x62, 0x4f, 0xce,
	0x48, 0xcc, 0xcb, 0x4b, 0xcd, 0x91, 0x60, 0x42, 0x55, 0x00, 0x13, 0x17, 0xb2, 0xe3, 0x62, 0x4b,
	0xcb, 0xcc, 0x29, 0x49, 0x2d, 0x92, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd3, 0xc3, 0xe3,
	0x15, 0xbd, 0x32, 0x43, 0x3d, 0x37, 0xb0, 0xea, 0x20, 0xa8, 0x2e, 0x27, 0xce, 0x28, 0x76, 0xb0,
	0xdf, 0xca, 0x0c, 0x93, 0xd8, 0xc0, 0x6e, 0x34, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xc0, 0xc0,
	0xc0, 0x39, 0x3c, 0x01, 0x00, 0x00,
}