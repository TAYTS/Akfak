// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/zkmessagepb/update_cluster_metadata.proto

package zkmessagepb

import (
	clustermetadatapb "AKFAK/proto/clustermetadatapb"
	commonpb "AKFAK/proto/commonpb"
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

type UpdateClusterMetadataRequest struct {
	NewClusterInfo       *clustermetadatapb.MetadataCluster `protobuf:"bytes,1,opt,name=newClusterInfo,proto3" json:"newClusterInfo,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                           `json:"-"`
	XXX_unrecognized     []byte                             `json:"-"`
	XXX_sizecache        int32                              `json:"-"`
}

func (m *UpdateClusterMetadataRequest) Reset()         { *m = UpdateClusterMetadataRequest{} }
func (m *UpdateClusterMetadataRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterMetadataRequest) ProtoMessage()    {}
func (*UpdateClusterMetadataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_13240d69a2718054, []int{0}
}

func (m *UpdateClusterMetadataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateClusterMetadataRequest.Unmarshal(m, b)
}
func (m *UpdateClusterMetadataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateClusterMetadataRequest.Marshal(b, m, deterministic)
}
func (m *UpdateClusterMetadataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateClusterMetadataRequest.Merge(m, src)
}
func (m *UpdateClusterMetadataRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateClusterMetadataRequest.Size(m)
}
func (m *UpdateClusterMetadataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateClusterMetadataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateClusterMetadataRequest proto.InternalMessageInfo

func (m *UpdateClusterMetadataRequest) GetNewClusterInfo() *clustermetadatapb.MetadataCluster {
	if m != nil {
		return m.NewClusterInfo
	}
	return nil
}

type UpdateClusterMetadataResponse struct {
	Response             *commonpb.Response `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *UpdateClusterMetadataResponse) Reset()         { *m = UpdateClusterMetadataResponse{} }
func (m *UpdateClusterMetadataResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateClusterMetadataResponse) ProtoMessage()    {}
func (*UpdateClusterMetadataResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_13240d69a2718054, []int{1}
}

func (m *UpdateClusterMetadataResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateClusterMetadataResponse.Unmarshal(m, b)
}
func (m *UpdateClusterMetadataResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateClusterMetadataResponse.Marshal(b, m, deterministic)
}
func (m *UpdateClusterMetadataResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateClusterMetadataResponse.Merge(m, src)
}
func (m *UpdateClusterMetadataResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateClusterMetadataResponse.Size(m)
}
func (m *UpdateClusterMetadataResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateClusterMetadataResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateClusterMetadataResponse proto.InternalMessageInfo

func (m *UpdateClusterMetadataResponse) GetResponse() *commonpb.Response {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*UpdateClusterMetadataRequest)(nil), "proto.zkmessagepb.UpdateClusterMetadataRequest")
	proto.RegisterType((*UpdateClusterMetadataResponse)(nil), "proto.zkmessagepb.UpdateClusterMetadataResponse")
}

func init() {
	proto.RegisterFile("proto/zkmessagepb/update_cluster_metadata.proto", fileDescriptor_13240d69a2718054)
}

var fileDescriptor_13240d69a2718054 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2f, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xaf, 0xca, 0xce, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x48, 0xd2, 0x2f, 0x2d,
	0x48, 0x49, 0x2c, 0x49, 0x8d, 0x4f, 0xce, 0x29, 0x2d, 0x2e, 0x49, 0x2d, 0x8a, 0xcf, 0x4d, 0x2d,
	0x49, 0x4c, 0x49, 0x2c, 0x49, 0xd4, 0x03, 0xab, 0x14, 0x12, 0x04, 0x53, 0x7a, 0x48, 0x1a, 0xa4,
	0x64, 0x21, 0x66, 0x24, 0xe7, 0xe7, 0xe6, 0xe6, 0xe7, 0x15, 0x24, 0xe9, 0x17, 0xa5, 0x16, 0x17,
	0xe4, 0xe7, 0x15, 0xa7, 0x42, 0x74, 0x48, 0xe9, 0x41, 0xa5, 0x21, 0xe6, 0xc1, 0x8c, 0x2b, 0x48,
	0xd2, 0xc7, 0x6e, 0x83, 0x52, 0x01, 0x97, 0x4c, 0x28, 0xd8, 0x09, 0xce, 0x10, 0x79, 0x5f, 0xa8,
	0x74, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x50, 0x00, 0x17, 0x5f, 0x5e, 0x6a, 0x39, 0x54,
	0xd2, 0x33, 0x2f, 0x2d, 0x5f, 0x82, 0x51, 0x81, 0x51, 0x83, 0xdb, 0x48, 0x03, 0xa2, 0x5f, 0x0f,
	0xc3, 0x22, 0x3d, 0x98, 0x09, 0x50, 0x3d, 0x41, 0x68, 0xfa, 0x95, 0x42, 0xb9, 0x64, 0x71, 0xd8,
	0x08, 0xf1, 0x88, 0x90, 0x09, 0x17, 0x07, 0xcc, 0x53, 0x50, 0xcb, 0x24, 0x60, 0x96, 0x41, 0x3d,
	0xad, 0x07, 0x53, 0x1b, 0x04, 0x57, 0xe9, 0x24, 0x19, 0x25, 0xee, 0xe8, 0xed, 0xe6, 0xe8, 0x8d,
	0x19, 0xc6, 0x49, 0x6c, 0x60, 0x21, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xed, 0xad, 0xb1,
	0x34, 0x7f, 0x01, 0x00, 0x00,
}
