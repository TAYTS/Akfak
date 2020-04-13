// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/zkmessagepb/get_cluster_metadata.proto

package zkmessagepb

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

type GetClusterMetadataRequest struct {
	Broker               *MetadataBroker `protobuf:"bytes,1,opt,name=broker,proto3" json:"broker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetClusterMetadataRequest) Reset()         { *m = GetClusterMetadataRequest{} }
func (m *GetClusterMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*GetClusterMetadataRequest) ProtoMessage()    {}
func (*GetClusterMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37ecb3af4c8fcc3, []int{0}
}

func (m *GetClusterMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClusterMetadataRequest.Unmarshal(m, b)
}
func (m *GetClusterMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClusterMetadataRequest.Marshal(b, m, deterministic)
}
func (m *GetClusterMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterMetadataRequest.Merge(m, src)
}
func (m *GetClusterMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_GetClusterMetadataRequest.Size(m)
}
func (m *GetClusterMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterMetadataRequest proto.InternalMessageInfo

func (m *GetClusterMetadataRequest) GetBroker() *MetadataBroker {
	if m != nil {
		return m.Broker
	}
	return nil
}

type GetClusterMetadataResponse struct {
	ClusterInfo          *MetadataCluster `protobuf:"bytes,1,opt,name=clusterInfo,proto3" json:"clusterInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *GetClusterMetadataResponse) Reset()         { *m = GetClusterMetadataResponse{} }
func (m *GetClusterMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*GetClusterMetadataResponse) ProtoMessage()    {}
func (*GetClusterMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a37ecb3af4c8fcc3, []int{1}
}

func (m *GetClusterMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetClusterMetadataResponse.Unmarshal(m, b)
}
func (m *GetClusterMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetClusterMetadataResponse.Marshal(b, m, deterministic)
}
func (m *GetClusterMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetClusterMetadataResponse.Merge(m, src)
}
func (m *GetClusterMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_GetClusterMetadataResponse.Size(m)
}
func (m *GetClusterMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetClusterMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetClusterMetadataResponse proto.InternalMessageInfo

func (m *GetClusterMetadataResponse) GetClusterInfo() *MetadataCluster {
	if m != nil {
		return m.ClusterInfo
	}
	return nil
}

func init() {
	proto.RegisterType((*GetClusterMetadataRequest)(nil), "proto.zkmessagepb.GetClusterMetadataRequest")
	proto.RegisterType((*GetClusterMetadataResponse)(nil), "proto.zkmessagepb.GetClusterMetadataResponse")
}

func init() {
	proto.RegisterFile("proto/zkmessagepb/get_cluster_metadata.proto", fileDescriptor_a37ecb3af4c8fcc3)
}

var fileDescriptor_a37ecb3af4c8fcc3 = []byte{
	// 181 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xaf, 0xca, 0xce, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x48, 0xd2, 0x4f, 0x4f,
	0x2d, 0x89, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0x2d, 0x49, 0x4c, 0x49,
	0x2c, 0x49, 0xd4, 0x03, 0x2b, 0x13, 0x12, 0x04, 0x53, 0x7a, 0x48, 0xaa, 0xa5, 0x34, 0x30, 0x0d,
	0xc0, 0xae, 0x59, 0x29, 0x8c, 0x4b, 0xd2, 0x3d, 0xb5, 0xc4, 0x19, 0x22, 0xe9, 0x0b, 0x95, 0x0b,
	0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xb2, 0xe4, 0x62, 0x4b, 0x2a, 0xca, 0xcf, 0x4e, 0x2d,
	0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd4, 0xc3, 0xb0, 0x4a, 0x0f, 0xa6, 0xc7, 0x09,
	0xac, 0x30, 0x08, 0xaa, 0x41, 0x29, 0x89, 0x4b, 0x0a, 0x9b, 0xb9, 0xc5, 0x05, 0xf9, 0x79, 0xc5,
	0xa9, 0x42, 0x2e, 0x5c, 0xdc, 0x50, 0xf7, 0x78, 0xe6, 0xa5, 0xe5, 0x43, 0x4d, 0x57, 0xc2, 0x63,
	0x3a, 0xd4, 0xa0, 0x20, 0x64, 0x6d, 0x4e, 0x92, 0x51, 0xe2, 0x8e, 0xde, 0x6e, 0x8e, 0xde, 0xfa,
	0x18, 0xbe, 0x4d, 0x62, 0x03, 0x0b, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x09, 0x2d,
	0xbe, 0x4a, 0x01, 0x00, 0x00,
}
