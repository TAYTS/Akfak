// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/adminclientpb/sync_messages.proto

package adminclientpb

import (
	recordpb "AKFAK/proto/recordpb"
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

//* REQUEST *
type SyncPartition struct {
	Partition            int32    `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	LogStartOffset       int64    `protobuf:"varint,2,opt,name=logStartOffset,proto3" json:"logStartOffset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncPartition) Reset()         { *m = SyncPartition{} }
func (m *SyncPartition) String() string { return proto.CompactTextString(m) }
func (*SyncPartition) ProtoMessage()    {}
func (*SyncPartition) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{0}
}

func (m *SyncPartition) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPartition.Unmarshal(m, b)
}
func (m *SyncPartition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPartition.Marshal(b, m, deterministic)
}
func (m *SyncPartition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPartition.Merge(m, src)
}
func (m *SyncPartition) XXX_Size() int {
	return xxx_messageInfo_SyncPartition.Size(m)
}
func (m *SyncPartition) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPartition.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPartition proto.InternalMessageInfo

func (m *SyncPartition) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *SyncPartition) GetLogStartOffset() int64 {
	if m != nil {
		return m.LogStartOffset
	}
	return 0
}

// SyncForgottenTopicsData is used to remove the topic from the fetch session
type SyncForgottenTopicsData struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Partitions           []int32  `protobuf:"varint,2,rep,packed,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncForgottenTopicsData) Reset()         { *m = SyncForgottenTopicsData{} }
func (m *SyncForgottenTopicsData) String() string { return proto.CompactTextString(m) }
func (*SyncForgottenTopicsData) ProtoMessage()    {}
func (*SyncForgottenTopicsData) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{1}
}

func (m *SyncForgottenTopicsData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncForgottenTopicsData.Unmarshal(m, b)
}
func (m *SyncForgottenTopicsData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncForgottenTopicsData.Marshal(b, m, deterministic)
}
func (m *SyncForgottenTopicsData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncForgottenTopicsData.Merge(m, src)
}
func (m *SyncForgottenTopicsData) XXX_Size() int {
	return xxx_messageInfo_SyncForgottenTopicsData.Size(m)
}
func (m *SyncForgottenTopicsData) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncForgottenTopicsData.DiscardUnknown(m)
}

var xxx_messageInfo_SyncForgottenTopicsData proto.InternalMessageInfo

func (m *SyncForgottenTopicsData) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SyncForgottenTopicsData) GetPartitions() []int32 {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type SyncTopic struct {
	Topic                string           `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	Partitions           []*SyncPartition `protobuf:"bytes,2,rep,name=partitions,proto3" json:"partitions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SyncTopic) Reset()         { *m = SyncTopic{} }
func (m *SyncTopic) String() string { return proto.CompactTextString(m) }
func (*SyncTopic) ProtoMessage()    {}
func (*SyncTopic) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{2}
}

func (m *SyncTopic) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncTopic.Unmarshal(m, b)
}
func (m *SyncTopic) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncTopic.Marshal(b, m, deterministic)
}
func (m *SyncTopic) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncTopic.Merge(m, src)
}
func (m *SyncTopic) XXX_Size() int {
	return xxx_messageInfo_SyncTopic.Size(m)
}
func (m *SyncTopic) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncTopic.DiscardUnknown(m)
}

var xxx_messageInfo_SyncTopic proto.InternalMessageInfo

func (m *SyncTopic) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SyncTopic) GetPartitions() []*SyncPartition {
	if m != nil {
		return m.Partitions
	}
	return nil
}

type SyncMessagesRequest struct {
	ReplicaID            int32                      `protobuf:"varint,1,opt,name=replicaID,proto3" json:"replicaID,omitempty"`
	Topics               []*SyncTopic               `protobuf:"bytes,2,rep,name=topics,proto3" json:"topics,omitempty"`
	ForgottenTopics      []*SyncForgottenTopicsData `protobuf:"bytes,3,rep,name=forgottenTopics,proto3" json:"forgottenTopics,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *SyncMessagesRequest) Reset()         { *m = SyncMessagesRequest{} }
func (m *SyncMessagesRequest) String() string { return proto.CompactTextString(m) }
func (*SyncMessagesRequest) ProtoMessage()    {}
func (*SyncMessagesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{3}
}

func (m *SyncMessagesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncMessagesRequest.Unmarshal(m, b)
}
func (m *SyncMessagesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncMessagesRequest.Marshal(b, m, deterministic)
}
func (m *SyncMessagesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncMessagesRequest.Merge(m, src)
}
func (m *SyncMessagesRequest) XXX_Size() int {
	return xxx_messageInfo_SyncMessagesRequest.Size(m)
}
func (m *SyncMessagesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncMessagesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SyncMessagesRequest proto.InternalMessageInfo

func (m *SyncMessagesRequest) GetReplicaID() int32 {
	if m != nil {
		return m.ReplicaID
	}
	return 0
}

func (m *SyncMessagesRequest) GetTopics() []*SyncTopic {
	if m != nil {
		return m.Topics
	}
	return nil
}

func (m *SyncMessagesRequest) GetForgottenTopics() []*SyncForgottenTopicsData {
	if m != nil {
		return m.ForgottenTopics
	}
	return nil
}

//* RESPONSE *
type SyncPartitionResponse struct {
	Partition            int32                   `protobuf:"varint,1,opt,name=partition,proto3" json:"partition,omitempty"`
	LogStartOffset       int64                   `protobuf:"varint,2,opt,name=logStartOffset,proto3" json:"logStartOffset,omitempty"`
	RecordSets           []*recordpb.RecordBatch `protobuf:"bytes,4,rep,name=recordSets,proto3" json:"recordSets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *SyncPartitionResponse) Reset()         { *m = SyncPartitionResponse{} }
func (m *SyncPartitionResponse) String() string { return proto.CompactTextString(m) }
func (*SyncPartitionResponse) ProtoMessage()    {}
func (*SyncPartitionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{4}
}

func (m *SyncPartitionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncPartitionResponse.Unmarshal(m, b)
}
func (m *SyncPartitionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncPartitionResponse.Marshal(b, m, deterministic)
}
func (m *SyncPartitionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncPartitionResponse.Merge(m, src)
}
func (m *SyncPartitionResponse) XXX_Size() int {
	return xxx_messageInfo_SyncPartitionResponse.Size(m)
}
func (m *SyncPartitionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncPartitionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncPartitionResponse proto.InternalMessageInfo

func (m *SyncPartitionResponse) GetPartition() int32 {
	if m != nil {
		return m.Partition
	}
	return 0
}

func (m *SyncPartitionResponse) GetLogStartOffset() int64 {
	if m != nil {
		return m.LogStartOffset
	}
	return 0
}

func (m *SyncPartitionResponse) GetRecordSets() []*recordpb.RecordBatch {
	if m != nil {
		return m.RecordSets
	}
	return nil
}

type SyncTopicResponse struct {
	Topic                string                   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	PartitionResponses   []*SyncPartitionResponse `protobuf:"bytes,2,rep,name=partitionResponses,proto3" json:"partitionResponses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *SyncTopicResponse) Reset()         { *m = SyncTopicResponse{} }
func (m *SyncTopicResponse) String() string { return proto.CompactTextString(m) }
func (*SyncTopicResponse) ProtoMessage()    {}
func (*SyncTopicResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{5}
}

func (m *SyncTopicResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncTopicResponse.Unmarshal(m, b)
}
func (m *SyncTopicResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncTopicResponse.Marshal(b, m, deterministic)
}
func (m *SyncTopicResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncTopicResponse.Merge(m, src)
}
func (m *SyncTopicResponse) XXX_Size() int {
	return xxx_messageInfo_SyncTopicResponse.Size(m)
}
func (m *SyncTopicResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncTopicResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncTopicResponse proto.InternalMessageInfo

func (m *SyncTopicResponse) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *SyncTopicResponse) GetPartitionResponses() []*SyncPartitionResponse {
	if m != nil {
		return m.PartitionResponses
	}
	return nil
}

type SyncMessagesResponse struct {
	Responses            []*SyncTopicResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SyncMessagesResponse) Reset()         { *m = SyncMessagesResponse{} }
func (m *SyncMessagesResponse) String() string { return proto.CompactTextString(m) }
func (*SyncMessagesResponse) ProtoMessage()    {}
func (*SyncMessagesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_58188d0c01e01550, []int{6}
}

func (m *SyncMessagesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncMessagesResponse.Unmarshal(m, b)
}
func (m *SyncMessagesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncMessagesResponse.Marshal(b, m, deterministic)
}
func (m *SyncMessagesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncMessagesResponse.Merge(m, src)
}
func (m *SyncMessagesResponse) XXX_Size() int {
	return xxx_messageInfo_SyncMessagesResponse.Size(m)
}
func (m *SyncMessagesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncMessagesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SyncMessagesResponse proto.InternalMessageInfo

func (m *SyncMessagesResponse) GetResponses() []*SyncTopicResponse {
	if m != nil {
		return m.Responses
	}
	return nil
}

func init() {
	proto.RegisterType((*SyncPartition)(nil), "proto.adminclientpb.SyncPartition")
	proto.RegisterType((*SyncForgottenTopicsData)(nil), "proto.adminclientpb.SyncForgottenTopicsData")
	proto.RegisterType((*SyncTopic)(nil), "proto.adminclientpb.SyncTopic")
	proto.RegisterType((*SyncMessagesRequest)(nil), "proto.adminclientpb.SyncMessagesRequest")
	proto.RegisterType((*SyncPartitionResponse)(nil), "proto.adminclientpb.SyncPartitionResponse")
	proto.RegisterType((*SyncTopicResponse)(nil), "proto.adminclientpb.SyncTopicResponse")
	proto.RegisterType((*SyncMessagesResponse)(nil), "proto.adminclientpb.SyncMessagesResponse")
}

func init() {
	proto.RegisterFile("proto/adminclientpb/sync_messages.proto", fileDescriptor_58188d0c01e01550)
}

var fileDescriptor_58188d0c01e01550 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0x51, 0x4b, 0x2a, 0x41,
	0x14, 0x66, 0xf5, 0x2a, 0x78, 0x2e, 0xf7, 0x5e, 0xee, 0x68, 0xb4, 0x65, 0x88, 0xec, 0x83, 0x49,
	0xc4, 0x0a, 0x05, 0xbd, 0xf4, 0xa4, 0x88, 0x10, 0x12, 0xc6, 0x58, 0x3d, 0x48, 0x10, 0xe3, 0x3a,
	0xda, 0x82, 0xee, 0x4c, 0x3b, 0xa7, 0x07, 0x7f, 0x40, 0x7f, 0xa2, 0xbf, 0xd3, 0x1f, 0x8b, 0x9d,
	0xdd, 0x51, 0xd7, 0x56, 0xea, 0xa1, 0x27, 0x67, 0x8e, 0xe7, 0x7c, 0xdf, 0x77, 0xbe, 0x6f, 0x16,
	0x8e, 0x65, 0x28, 0x50, 0xb4, 0xd8, 0x64, 0xe1, 0x07, 0xde, 0xdc, 0xe7, 0x01, 0xca, 0x71, 0x4b,
	0x2d, 0x03, 0xef, 0x71, 0xc1, 0x95, 0x62, 0x33, 0xae, 0x5c, 0xdd, 0x41, 0xca, 0xfa, 0xc7, 0x4d,
	0x35, 0x1e, 0x56, 0xe3, 0xe9, 0x90, 0x7b, 0x22, 0x9c, 0xc8, 0x71, 0x72, 0x88, 0x27, 0x9c, 0x3b,
	0xf8, 0x33, 0x5c, 0x06, 0xde, 0x0d, 0x0b, 0xd1, 0x47, 0x5f, 0x04, 0xe4, 0x08, 0x4a, 0xd2, 0x5c,
	0x6c, 0xab, 0x6e, 0x35, 0x0b, 0x74, 0x5d, 0x20, 0x0d, 0xf8, 0x3b, 0x17, 0xb3, 0x21, 0xb2, 0x10,
	0x07, 0xd3, 0xa9, 0xe2, 0x68, 0xe7, 0xea, 0x56, 0x33, 0x4f, 0xb7, 0xaa, 0xce, 0x00, 0xf6, 0x23,
	0xd8, 0x9e, 0x08, 0x67, 0x02, 0x91, 0x07, 0xb7, 0x42, 0xfa, 0x9e, 0xea, 0x32, 0x64, 0xa4, 0x02,
	0x05, 0x8c, 0x6e, 0x1a, 0xbc, 0x44, 0xe3, 0x0b, 0xa9, 0x01, 0xac, 0x58, 0x94, 0x9d, 0xab, 0xe7,
	0x9b, 0x05, 0xba, 0x51, 0x71, 0x38, 0x94, 0x22, 0x40, 0x8d, 0xb3, 0x03, 0xa2, 0xf3, 0x09, 0xe2,
	0xf7, 0x99, 0xe3, 0x66, 0x38, 0xe2, 0xa6, 0x36, 0x4e, 0xd1, 0xbc, 0x5b, 0x50, 0x8e, 0xfe, 0xbd,
	0x4e, 0x7c, 0xa5, 0xfc, 0xf9, 0x85, 0x2b, 0x8c, 0x5c, 0x09, 0xb9, 0x9c, 0xfb, 0x1e, 0xbb, 0xea,
	0x1a, 0x57, 0x56, 0x05, 0x72, 0x01, 0x45, 0x2d, 0xc1, 0xb0, 0xd6, 0x76, 0xb2, 0x6a, 0xfd, 0x34,
	0xe9, 0x26, 0xf7, 0xf0, 0x6f, 0x9a, 0x76, 0xc8, 0xce, 0x6b, 0x80, 0xd3, 0x9d, 0x00, 0x19, 0x8e,
	0xd2, 0x6d, 0x10, 0xe7, 0xcd, 0x82, 0xbd, 0xf4, 0x8e, 0x5c, 0x49, 0x11, 0x28, 0xfe, 0x33, 0xe9,
	0x92, 0x4b, 0x80, 0xf8, 0x11, 0x0d, 0x39, 0x2a, 0xfb, 0x97, 0x96, 0x5c, 0x4d, 0x24, 0x9b, 0x67,
	0xe6, 0x52, 0x7d, 0xe8, 0x30, 0xf4, 0x9e, 0xe8, 0x46, 0xbb, 0xf3, 0x6a, 0xc1, 0xff, 0xb5, 0x15,
	0x46, 0x58, 0x76, 0xa4, 0x23, 0x20, 0x72, 0x7b, 0x07, 0x63, 0xf2, 0xc9, 0x37, 0xa2, 0x4d, 0x46,
	0x68, 0x06, 0x8a, 0xf3, 0x00, 0x95, 0x74, 0xd2, 0x89, 0x92, 0x6e, 0x14, 0xb5, 0xa1, 0xb2, 0x34,
	0x55, 0xe3, 0x8b, 0x3c, 0x0d, 0xcd, 0x7a, 0xb0, 0x53, 0x1d, 0x1d, 0xb4, 0xfb, 0xbd, 0x76, 0xbf,
	0x95, 0xf1, 0xe9, 0x8e, 0x8b, 0xba, 0x78, 0xfe, 0x11, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x03, 0x6e,
	0xbd, 0xd8, 0x03, 0x00, 0x00,
}
