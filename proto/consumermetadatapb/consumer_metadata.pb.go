// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/consumermetadatapb/consumer_metadata.proto

package consumermetadatapb

import (
	clustermetadatapb "AKFAK/proto/clustermetadatapb"
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

type Assignment struct {
	TopicName            string                            `protobuf:"bytes,1,opt,name=topicName,proto3" json:"topicName,omitempty"`
	PartitionIndex       int32                             `protobuf:"varint,2,opt,name=partitionIndex,proto3" json:"partitionIndex,omitempty"`
	Offset               int32                             `protobuf:"varint,3,opt,name=offset,proto3" json:"offset,omitempty"`
	Broker               *clustermetadatapb.MetadataBroker `protobuf:"bytes,4,opt,name=broker,proto3" json:"broker,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Assignment) Reset()         { *m = Assignment{} }
func (m *Assignment) String() string { return proto.CompactTextString(m) }
func (*Assignment) ProtoMessage()    {}
func (*Assignment) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad10e0c905c64b8e, []int{0}
}

func (m *Assignment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Assignment.Unmarshal(m, b)
}
func (m *Assignment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Assignment.Marshal(b, m, deterministic)
}
func (m *Assignment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assignment.Merge(m, src)
}
func (m *Assignment) XXX_Size() int {
	return xxx_messageInfo_Assignment.Size(m)
}
func (m *Assignment) XXX_DiscardUnknown() {
	xxx_messageInfo_Assignment.DiscardUnknown(m)
}

var xxx_messageInfo_Assignment proto.InternalMessageInfo

func (m *Assignment) GetTopicName() string {
	if m != nil {
		return m.TopicName
	}
	return ""
}

func (m *Assignment) GetPartitionIndex() int32 {
	if m != nil {
		return m.PartitionIndex
	}
	return 0
}

func (m *Assignment) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Assignment) GetBroker() *clustermetadatapb.MetadataBroker {
	if m != nil {
		return m.Broker
	}
	return nil
}

type Consumer struct {
	ID                   int32         `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Assignments          []*Assignment `protobuf:"bytes,2,rep,name=assignments,proto3" json:"assignments,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Consumer) Reset()         { *m = Consumer{} }
func (m *Consumer) String() string { return proto.CompactTextString(m) }
func (*Consumer) ProtoMessage()    {}
func (*Consumer) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad10e0c905c64b8e, []int{1}
}

func (m *Consumer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consumer.Unmarshal(m, b)
}
func (m *Consumer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consumer.Marshal(b, m, deterministic)
}
func (m *Consumer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consumer.Merge(m, src)
}
func (m *Consumer) XXX_Size() int {
	return xxx_messageInfo_Consumer.Size(m)
}
func (m *Consumer) XXX_DiscardUnknown() {
	xxx_messageInfo_Consumer.DiscardUnknown(m)
}

var xxx_messageInfo_Consumer proto.InternalMessageInfo

func (m *Consumer) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Consumer) GetAssignments() []*Assignment {
	if m != nil {
		return m.Assignments
	}
	return nil
}

type ConsumerGroup struct {
	ID                   int32       `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Consumers            []*Consumer `protobuf:"bytes,2,rep,name=consumers,proto3" json:"consumers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *ConsumerGroup) Reset()         { *m = ConsumerGroup{} }
func (m *ConsumerGroup) String() string { return proto.CompactTextString(m) }
func (*ConsumerGroup) ProtoMessage()    {}
func (*ConsumerGroup) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad10e0c905c64b8e, []int{2}
}

func (m *ConsumerGroup) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConsumerGroup.Unmarshal(m, b)
}
func (m *ConsumerGroup) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConsumerGroup.Marshal(b, m, deterministic)
}
func (m *ConsumerGroup) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerGroup.Merge(m, src)
}
func (m *ConsumerGroup) XXX_Size() int {
	return xxx_messageInfo_ConsumerGroup.Size(m)
}
func (m *ConsumerGroup) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerGroup.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerGroup proto.InternalMessageInfo

func (m *ConsumerGroup) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ConsumerGroup) GetConsumers() []*Consumer {
	if m != nil {
		return m.Consumers
	}
	return nil
}

type MetadataConsumerState struct {
	ConsumerGroups       []*ConsumerGroup `protobuf:"bytes,1,rep,name=consumerGroups,proto3" json:"consumerGroups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MetadataConsumerState) Reset()         { *m = MetadataConsumerState{} }
func (m *MetadataConsumerState) String() string { return proto.CompactTextString(m) }
func (*MetadataConsumerState) ProtoMessage()    {}
func (*MetadataConsumerState) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad10e0c905c64b8e, []int{3}
}

func (m *MetadataConsumerState) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetadataConsumerState.Unmarshal(m, b)
}
func (m *MetadataConsumerState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetadataConsumerState.Marshal(b, m, deterministic)
}
func (m *MetadataConsumerState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetadataConsumerState.Merge(m, src)
}
func (m *MetadataConsumerState) XXX_Size() int {
	return xxx_messageInfo_MetadataConsumerState.Size(m)
}
func (m *MetadataConsumerState) XXX_DiscardUnknown() {
	xxx_messageInfo_MetadataConsumerState.DiscardUnknown(m)
}

var xxx_messageInfo_MetadataConsumerState proto.InternalMessageInfo

func (m *MetadataConsumerState) GetConsumerGroups() []*ConsumerGroup {
	if m != nil {
		return m.ConsumerGroups
	}
	return nil
}

func init() {
	proto.RegisterType((*Assignment)(nil), "proto.clustermetadatapb.Assignment")
	proto.RegisterType((*Consumer)(nil), "proto.clustermetadatapb.Consumer")
	proto.RegisterType((*ConsumerGroup)(nil), "proto.clustermetadatapb.ConsumerGroup")
	proto.RegisterType((*MetadataConsumerState)(nil), "proto.clustermetadatapb.MetadataConsumerState")
}

func init() {
	proto.RegisterFile("proto/consumermetadatapb/consumer_metadata.proto", fileDescriptor_ad10e0c905c64b8e)
}

var fileDescriptor_ad10e0c905c64b8e = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x25, 0x9d, 0x1b, 0xee, 0x1b, 0xf6, 0x10, 0x50, 0x83, 0x88, 0xd4, 0x0a, 0xb3, 0xa7, 0x4e,
	0xe6, 0x0f, 0x18, 0x9d, 0x53, 0x29, 0xc3, 0x1d, 0xe2, 0xcd, 0x8b, 0xa6, 0x5d, 0x36, 0x8a, 0xb6,
	0x29, 0x49, 0x0a, 0xfe, 0x27, 0xff, 0xa4, 0x10, 0x9a, 0xb5, 0x54, 0x8a, 0xa7, 0x36, 0xef, 0x7d,
	0xef, 0xbd, 0xbc, 0x24, 0x70, 0x57, 0x4a, 0xa1, 0xc5, 0x2c, 0x15, 0x85, 0xaa, 0x72, 0x2e, 0x73,
	0xae, 0xd9, 0x96, 0x69, 0x56, 0x26, 0x07, 0xe8, 0xdd, 0x62, 0xa1, 0x19, 0xc5, 0xe7, 0xe6, 0x13,
	0xa6, 0x5f, 0x95, 0xd2, 0x6d, 0xc1, 0x45, 0x58, 0x5b, 0x75, 0x09, 0x8b, 0x74, 0x8c, 0xfc, 0x1f,
	0x04, 0x10, 0x29, 0x95, 0xed, 0x8b, 0x9c, 0x17, 0x1a, 0x5f, 0xc2, 0x58, 0x8b, 0x32, 0x4b, 0x37,
	0x2c, 0xe7, 0x04, 0x79, 0x28, 0x18, 0xd3, 0x06, 0xc0, 0x53, 0x70, 0x4b, 0x26, 0x75, 0xa6, 0x33,
	0x51, 0xc4, 0xc5, 0x96, 0x7f, 0x13, 0xc7, 0x43, 0xc1, 0x90, 0x76, 0x50, 0x7c, 0x06, 0x23, 0xb1,
	0xdb, 0x29, 0xae, 0xc9, 0xc0, 0xf0, 0xf5, 0x0a, 0x2f, 0x60, 0x94, 0x48, 0xf1, 0xc9, 0x25, 0x39,
	0xf2, 0x50, 0x30, 0x99, 0xdf, 0x86, 0x3d, 0x35, 0xc2, 0x97, 0xfa, 0x77, 0x69, 0xc6, 0x69, 0x2d,
	0xf3, 0x19, 0x1c, 0x3f, 0xd4, 0x27, 0x82, 0x5d, 0x70, 0xe2, 0x95, 0xd9, 0xe3, 0x90, 0x3a, 0xf1,
	0x0a, 0x3f, 0xc2, 0x84, 0x1d, 0x8a, 0x28, 0xe2, 0x78, 0x83, 0x60, 0x32, 0xbf, 0xe9, 0x4d, 0x68,
	0x4a, 0xd3, 0xb6, 0xce, 0xff, 0x80, 0x13, 0x1b, 0xf1, 0x2c, 0x45, 0x55, 0xfe, 0xc9, 0x59, 0xc0,
	0xd8, 0xde, 0x8a, 0x4d, 0xb9, 0xee, 0x4d, 0xb1, 0x56, 0xb4, 0xd1, 0xf8, 0x7b, 0x38, 0xb5, 0xf5,
	0x2c, 0xfd, 0xaa, 0x99, 0xe6, 0x78, 0x03, 0x6e, 0xda, 0x8e, 0x56, 0x04, 0x19, 0xfb, 0xe9, 0xbf,
	0xf6, 0x66, 0x9c, 0x76, 0xd4, 0x4b, 0xef, 0xed, 0x2a, 0x5a, 0x3f, 0x45, 0xeb, 0x59, 0xdf, 0xf3,
	0x4a, 0x46, 0x86, 0xb9, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x1c, 0xe9, 0x7f, 0x9b, 0x81, 0x02,
	0x00, 0x00,
}
