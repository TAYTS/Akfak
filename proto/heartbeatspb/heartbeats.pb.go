// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/heartbeatspb/heartbeats.proto

package heartbeatspb

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

type HeartbeatsRequest struct {
	BrokerID             int32    `protobuf:"varint,1,opt,name=brokerID,proto3" json:"brokerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatsRequest) Reset()         { *m = HeartbeatsRequest{} }
func (m *HeartbeatsRequest) String() string { return proto.CompactTextString(m) }
func (*HeartbeatsRequest) ProtoMessage()    {}
func (*HeartbeatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f00df0763f3e87f3, []int{0}
}

func (m *HeartbeatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatsRequest.Unmarshal(m, b)
}
func (m *HeartbeatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatsRequest.Marshal(b, m, deterministic)
}
func (m *HeartbeatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatsRequest.Merge(m, src)
}
func (m *HeartbeatsRequest) XXX_Size() int {
	return xxx_messageInfo_HeartbeatsRequest.Size(m)
}
func (m *HeartbeatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatsRequest proto.InternalMessageInfo

func (m *HeartbeatsRequest) GetBrokerID() int32 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

type HeartbeatsResponse struct {
	BrokerID             int32    `protobuf:"varint,1,opt,name=brokerID,proto3" json:"brokerID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeartbeatsResponse) Reset()         { *m = HeartbeatsResponse{} }
func (m *HeartbeatsResponse) String() string { return proto.CompactTextString(m) }
func (*HeartbeatsResponse) ProtoMessage()    {}
func (*HeartbeatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f00df0763f3e87f3, []int{1}
}

func (m *HeartbeatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeartbeatsResponse.Unmarshal(m, b)
}
func (m *HeartbeatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeartbeatsResponse.Marshal(b, m, deterministic)
}
func (m *HeartbeatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeartbeatsResponse.Merge(m, src)
}
func (m *HeartbeatsResponse) XXX_Size() int {
	return xxx_messageInfo_HeartbeatsResponse.Size(m)
}
func (m *HeartbeatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HeartbeatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HeartbeatsResponse proto.InternalMessageInfo

func (m *HeartbeatsResponse) GetBrokerID() int32 {
	if m != nil {
		return m.BrokerID
	}
	return 0
}

func init() {
	proto.RegisterType((*HeartbeatsRequest)(nil), "proto.heartbeatspb.HeartbeatsRequest")
	proto.RegisterType((*HeartbeatsResponse)(nil), "proto.heartbeatspb.HeartbeatsResponse")
}

func init() {
	proto.RegisterFile("proto/heartbeatspb/heartbeats.proto", fileDescriptor_f00df0763f3e87f3)
}

var fileDescriptor_f00df0763f3e87f3 = []byte{
	// 122 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x48, 0x4d, 0x2c, 0x2a, 0x49, 0x4a, 0x4d, 0x2c, 0x29, 0x2e, 0x48, 0x42, 0xe2,
	0xe8, 0x81, 0x65, 0x85, 0x84, 0xc0, 0x94, 0x1e, 0xb2, 0x22, 0x25, 0x7d, 0x2e, 0x41, 0x0f, 0x38,
	0x3f, 0x28, 0xb5, 0xb0, 0x34, 0xb5, 0xb8, 0x44, 0x48, 0x8a, 0x8b, 0x23, 0xa9, 0x28, 0x3f, 0x3b,
	0xb5, 0xc8, 0xd3, 0x45, 0x82, 0x51, 0x81, 0x51, 0x83, 0x35, 0x08, 0xce, 0x57, 0x32, 0xe0, 0x12,
	0x42, 0xd6, 0x50, 0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x8a, 0x4f, 0x87, 0x93, 0x54, 0x94, 0x84, 0xa3,
	0xb7, 0x9b, 0xa3, 0xb7, 0x3e, 0xa6, 0x1b, 0x93, 0xd8, 0xc0, 0x62, 0xc6, 0x80, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xfb, 0xcd, 0x76, 0x47, 0xc0, 0x00, 0x00, 0x00,
}