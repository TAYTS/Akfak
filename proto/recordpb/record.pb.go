// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/recordpb/zk.proto

package recordpb

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

type Record struct {
	Length     int32 `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Attributes int32 `protobuf:"varint,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// attributes
	//bit 0~7: unused
	TimestampDelta       int32    `protobuf:"varint,3,opt,name=timestampDelta,proto3" json:"timestampDelta,omitempty"`
	OffsetDelta          int32    `protobuf:"varint,4,opt,name=offsetDelta,proto3" json:"offsetDelta,omitempty"`
	ValueLen             int32    `protobuf:"varint,5,opt,name=valueLen,proto3" json:"valueLen,omitempty"`
	Value                []byte   `protobuf:"bytes,6,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Record) Reset()         { *m = Record{} }
func (m *Record) String() string { return proto.CompactTextString(m) }
func (*Record) ProtoMessage()    {}
func (*Record) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea920a3f23369c2d, []int{0}
}

func (m *Record) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Record.Unmarshal(m, b)
}
func (m *Record) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Record.Marshal(b, m, deterministic)
}
func (m *Record) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Record.Merge(m, src)
}
func (m *Record) XXX_Size() int {
	return xxx_messageInfo_Record.Size(m)
}
func (m *Record) XXX_DiscardUnknown() {
	xxx_messageInfo_Record.DiscardUnknown(m)
}

var xxx_messageInfo_Record proto.InternalMessageInfo

func (m *Record) GetLength() int32 {
	if m != nil {
		return m.Length
	}
	return 0
}

func (m *Record) GetAttributes() int32 {
	if m != nil {
		return m.Attributes
	}
	return 0
}

func (m *Record) GetTimestampDelta() int32 {
	if m != nil {
		return m.TimestampDelta
	}
	return 0
}

func (m *Record) GetOffsetDelta() int32 {
	if m != nil {
		return m.OffsetDelta
	}
	return 0
}

func (m *Record) GetValueLen() int32 {
	if m != nil {
		return m.ValueLen
	}
	return 0
}

func (m *Record) GetValue() []byte {
	if m != nil {
		return m.Value
	}
	return nil
}

type RecordBatch struct {
	BaseOffset           int64  `protobuf:"varint,1,opt,name=baseOffset,proto3" json:"baseOffset,omitempty"`
	BatchLength          []byte `protobuf:"bytes,2,opt,name=batchLength,proto3" json:"batchLength,omitempty"`
	PartitionLeaderEpoch int32  `protobuf:"varint,3,opt,name=partitionLeaderEpoch,proto3" json:"partitionLeaderEpoch,omitempty"`
	Magic                int32  `protobuf:"varint,4,opt,name=magic,proto3" json:"magic,omitempty"`
	Crc                  int32  `protobuf:"varint,5,opt,name=crc,proto3" json:"crc,omitempty"`
	Attributes           int32  `protobuf:"varint,6,opt,name=attributes,proto3" json:"attributes,omitempty"`
	// attributes
	//bit 0~2:
	//0: no compression
	//1: gzip
	//2: snappy
	//3: lz4
	//4: zstd
	//bit 3: timestampType
	//bit 4: isTransactional (0 means not transactional)
	//bit 5: isControlBatch (0 means not a control batch)
	//bit 6~15: unused
	LastOffsetDelta      int32     `protobuf:"varint,7,opt,name=lastOffsetDelta,proto3" json:"lastOffsetDelta,omitempty"`
	FirstTimestamp       int64     `protobuf:"varint,8,opt,name=firstTimestamp,proto3" json:"firstTimestamp,omitempty"`
	MaxTimestamp         int64     `protobuf:"varint,9,opt,name=maxTimestamp,proto3" json:"maxTimestamp,omitempty"`
	Records              []*Record `protobuf:"bytes,13,rep,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RecordBatch) Reset()         { *m = RecordBatch{} }
func (m *RecordBatch) String() string { return proto.CompactTextString(m) }
func (*RecordBatch) ProtoMessage()    {}
func (*RecordBatch) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea920a3f23369c2d, []int{1}
}

func (m *RecordBatch) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordBatch.Unmarshal(m, b)
}
func (m *RecordBatch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordBatch.Marshal(b, m, deterministic)
}
func (m *RecordBatch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordBatch.Merge(m, src)
}
func (m *RecordBatch) XXX_Size() int {
	return xxx_messageInfo_RecordBatch.Size(m)
}
func (m *RecordBatch) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordBatch.DiscardUnknown(m)
}

var xxx_messageInfo_RecordBatch proto.InternalMessageInfo

func (m *RecordBatch) GetBaseOffset() int64 {
	if m != nil {
		return m.BaseOffset
	}
	return 0
}

func (m *RecordBatch) GetBatchLength() []byte {
	if m != nil {
		return m.BatchLength
	}
	return nil
}

func (m *RecordBatch) GetPartitionLeaderEpoch() int32 {
	if m != nil {
		return m.PartitionLeaderEpoch
	}
	return 0
}

func (m *RecordBatch) GetMagic() int32 {
	if m != nil {
		return m.Magic
	}
	return 0
}

func (m *RecordBatch) GetCrc() int32 {
	if m != nil {
		return m.Crc
	}
	return 0
}

func (m *RecordBatch) GetAttributes() int32 {
	if m != nil {
		return m.Attributes
	}
	return 0
}

func (m *RecordBatch) GetLastOffsetDelta() int32 {
	if m != nil {
		return m.LastOffsetDelta
	}
	return 0
}

func (m *RecordBatch) GetFirstTimestamp() int64 {
	if m != nil {
		return m.FirstTimestamp
	}
	return 0
}

func (m *RecordBatch) GetMaxTimestamp() int64 {
	if m != nil {
		return m.MaxTimestamp
	}
	return 0
}

func (m *RecordBatch) GetRecords() []*Record {
	if m != nil {
		return m.Records
	}
	return nil
}

func init() {
	proto.RegisterType((*Record)(nil), "proto.recordpb.Record")
	proto.RegisterType((*RecordBatch)(nil), "proto.recordpb.RecordBatch")
}

func init() {
	proto.RegisterFile("proto/recordpb/zk.proto", fileDescriptor_ea920a3f23369c2d)
}

var fileDescriptor_ea920a3f23369c2d = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0x41, 0x6b, 0xf2, 0x40,
	0x10, 0x25, 0xe6, 0x33, 0xfa, 0x8d, 0xd6, 0x96, 0x45, 0x64, 0x69, 0xa1, 0x04, 0x0f, 0x25, 0xa7,
	0x58, 0xec, 0x2f, 0x50, 0xda, 0x5e, 0x14, 0x84, 0xd0, 0x53, 0x6f, 0x9b, 0xb8, 0x6a, 0x20, 0x71,
	0xc3, 0xee, 0x58, 0xfa, 0xe3, 0xfa, 0xaf, 0xfa, 0x07, 0x4a, 0x66, 0x63, 0x9b, 0xa4, 0x3d, 0xed,
	0xbc, 0x37, 0x8f, 0xe5, 0xbd, 0xc7, 0xc0, 0x4d, 0xa1, 0x15, 0xaa, 0x99, 0x96, 0x89, 0xd2, 0xdb,
	0x22, 0xae, 0x86, 0x90, 0x58, 0x36, 0xa2, 0x27, 0x3c, 0x2f, 0xa7, 0x1f, 0x0e, 0x78, 0x11, 0x01,
	0x36, 0x01, 0x2f, 0x93, 0xc7, 0x3d, 0x1e, 0xb8, 0xe3, 0x3b, 0x41, 0x37, 0xaa, 0x10, 0xbb, 0x05,
	0x10, 0x88, 0x3a, 0x8d, 0x4f, 0x28, 0x0d, 0xef, 0xd0, 0xae, 0xc6, 0xb0, 0x3b, 0x18, 0x61, 0x9a,
	0x4b, 0x83, 0x22, 0x2f, 0x1e, 0x65, 0x86, 0x82, 0xbb, 0xa4, 0x69, 0xb1, 0xcc, 0x87, 0x81, 0xda,
	0xed, 0x8c, 0x44, 0x2b, 0xfa, 0x47, 0xa2, 0x3a, 0xc5, 0xae, 0xa1, 0xff, 0x26, 0xb2, 0x93, 0x5c,
	0xcb, 0x23, 0xef, 0xd2, 0xfa, 0x1b, 0xb3, 0x31, 0x74, 0x69, 0xe6, 0x9e, 0xef, 0x04, 0xc3, 0xc8,
	0x82, 0xe9, 0x67, 0x07, 0x06, 0xd6, 0xfe, 0x52, 0x60, 0x42, 0x5e, 0x63, 0x61, 0xe4, 0x86, 0x3e,
	0xa5, 0x1c, 0x6e, 0x54, 0x63, 0x4a, 0x0f, 0x71, 0x29, 0x5c, 0xdb, 0xa0, 0x1d, 0xfa, 0xab, 0x4e,
	0xb1, 0x39, 0x8c, 0x0b, 0xa1, 0x31, 0xc5, 0x54, 0x1d, 0xd7, 0x52, 0x6c, 0xa5, 0x7e, 0x2a, 0x54,
	0x72, 0xa8, 0x32, 0xfd, 0xb9, 0x2b, 0xbd, 0xe5, 0x62, 0x9f, 0x26, 0x55, 0x26, 0x0b, 0xd8, 0x15,
	0xb8, 0x89, 0x4e, 0xaa, 0x20, 0xe5, 0xd8, 0x6a, 0xd2, 0xfb, 0xd5, 0x64, 0x00, 0x97, 0x99, 0x30,
	0xb8, 0xa9, 0xb5, 0xd4, 0x23, 0x51, 0x9b, 0x2e, 0x3b, 0xdf, 0xa5, 0xda, 0xe0, 0xcb, 0xb9, 0x62,
	0xde, 0xa7, 0xac, 0x2d, 0x96, 0x4d, 0x61, 0x98, 0x8b, 0xf7, 0x1f, 0xd5, 0x7f, 0x52, 0x35, 0x38,
	0x76, 0x0f, 0x3d, 0x7b, 0x0e, 0x86, 0x5f, 0xf8, 0x6e, 0x30, 0x98, 0x4f, 0xc2, 0xe6, 0x91, 0x84,
	0xb6, 0xe1, 0xe8, 0x2c, 0x5b, 0x4e, 0x5e, 0xc7, 0x8b, 0xd5, 0xf3, 0x62, 0x35, 0x6b, 0x5e, 0x5a,
	0xec, 0x11, 0x7e, 0xf8, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xef, 0x9a, 0xae, 0x2e, 0x82, 0x02, 0x00,
	0x00,
}
