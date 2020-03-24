// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/recordpb/record.proto

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
	Length               int32    `protobuf:"varint,1,opt,name=length,proto3" json:"length,omitempty"`
	Attributes           int32    `protobuf:"varint,2,opt,name=attributes,proto3" json:"attributes,omitempty"`
	TimestampDelta       int32    `protobuf:"varint,3,opt,name=timestampDelta,proto3" json:"timestampDelta,omitempty"`
	OffsetDelta          int32    `protobuf:"varint,4,opt,name=offsetDelta,proto3" json:"offsetDelta,omitempty"`
	KeyLength            int32    `protobuf:"varint,5,opt,name=keyLength,proto3" json:"keyLength,omitempty"`
	Key                  []byte   `protobuf:"bytes,6,opt,name=key,proto3" json:"key,omitempty"`
	ValueLen             int32    `protobuf:"varint,7,opt,name=valueLen,proto3" json:"valueLen,omitempty"`
	Value                []byte   `protobuf:"bytes,8,opt,name=value,proto3" json:"value,omitempty"`
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

func (m *Record) GetKeyLength() int32 {
	if m != nil {
		return m.KeyLength
	}
	return 0
}

func (m *Record) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
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
	ProducerId           int64     `protobuf:"varint,10,opt,name=producerId,proto3" json:"producerId,omitempty"`
	ProducerEpoch        int32     `protobuf:"varint,11,opt,name=producerEpoch,proto3" json:"producerEpoch,omitempty"`
	BaseSequence         int32     `protobuf:"varint,12,opt,name=baseSequence,proto3" json:"baseSequence,omitempty"`
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

func (m *RecordBatch) GetProducerId() int64 {
	if m != nil {
		return m.ProducerId
	}
	return 0
}

func (m *RecordBatch) GetProducerEpoch() int32 {
	if m != nil {
		return m.ProducerEpoch
	}
	return 0
}

func (m *RecordBatch) GetBaseSequence() int32 {
	if m != nil {
		return m.BaseSequence
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
	proto.RegisterFile("proto/recordpb/record.proto", fileDescriptor_ea920a3f23369c2d)
}

var fileDescriptor_ea920a3f23369c2d = []byte{
	// 403 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0xdf, 0x8a, 0xd3, 0x40,
	0x14, 0xc6, 0xc9, 0xc6, 0x66, 0x77, 0x4f, 0xba, 0xab, 0x0c, 0xa5, 0x0c, 0x2a, 0x12, 0x8a, 0x48,
	0xae, 0xb2, 0xb2, 0x3e, 0xc1, 0x2e, 0x2a, 0x48, 0x0b, 0x85, 0xe8, 0x95, 0x77, 0x93, 0xc9, 0x69,
	0x1b, 0x9a, 0x7f, 0x4e, 0x26, 0x62, 0x1f, 0xd6, 0x37, 0xf0, 0x21, 0x64, 0xce, 0x24, 0x36, 0x89,
	0x7b, 0x95, 0x73, 0x7e, 0xe7, 0x23, 0xf9, 0xce, 0x37, 0x13, 0x78, 0x55, 0xab, 0x4a, 0x57, 0x77,
	0x0a, 0x65, 0xa5, 0xd2, 0x3a, 0xe9, 0x8a, 0x88, 0x28, 0xbb, 0xa5, 0x47, 0xd4, 0x0f, 0x57, 0x7f,
	0x1c, 0xf0, 0x62, 0x6a, 0xd8, 0x12, 0xbc, 0x1c, 0xcb, 0xbd, 0x3e, 0x70, 0x27, 0x70, 0xc2, 0x59,
	0xdc, 0x75, 0xec, 0x0d, 0x80, 0xd0, 0x5a, 0x65, 0x49, 0xab, 0xb1, 0xe1, 0x17, 0x34, 0x1b, 0x10,
	0xf6, 0x0e, 0x6e, 0x75, 0x56, 0x60, 0xa3, 0x45, 0x51, 0x7f, 0xc4, 0x5c, 0x0b, 0xee, 0x92, 0x66,
	0x42, 0x59, 0x00, 0x7e, 0xb5, 0xdb, 0x35, 0xa8, 0xad, 0xe8, 0x19, 0x89, 0x86, 0x88, 0xbd, 0x86,
	0xeb, 0x23, 0x9e, 0x36, 0xd6, 0xc4, 0x8c, 0xe6, 0x67, 0xc0, 0x5e, 0x80, 0x7b, 0xc4, 0x13, 0xf7,
	0x02, 0x27, 0x9c, 0xc7, 0xa6, 0x64, 0x2f, 0xe1, 0xea, 0xa7, 0xc8, 0x5b, 0xdc, 0x60, 0xc9, 0x2f,
	0x49, 0xfe, 0xaf, 0x67, 0x0b, 0x98, 0x51, 0xcd, 0xaf, 0x48, 0x6f, 0x9b, 0xd5, 0x6f, 0x17, 0x7c,
	0xbb, 0xee, 0xa3, 0xd0, 0x92, 0x76, 0x4b, 0x44, 0x83, 0x5b, 0x32, 0x41, 0x7b, 0xbb, 0xf1, 0x80,
	0x18, 0xcf, 0x89, 0x11, 0x76, 0x9e, 0x2e, 0xe8, 0x5d, 0x43, 0xc4, 0xee, 0x61, 0x51, 0x0b, 0xa5,
	0x33, 0x9d, 0x55, 0xe5, 0x06, 0x45, 0x8a, 0xea, 0x53, 0x5d, 0xc9, 0x43, 0x97, 0xc1, 0x93, 0x33,
	0xe3, 0xad, 0x10, 0xfb, 0x4c, 0x76, 0x19, 0xd8, 0xc6, 0xec, 0x27, 0x95, 0xec, 0xf6, 0x36, 0xe5,
	0x24, 0x79, 0xef, 0xbf, 0xe4, 0x43, 0x78, 0x9e, 0x8b, 0x46, 0x6f, 0x07, 0xa9, 0xda, 0x18, 0xa6,
	0xd8, 0x9c, 0xd1, 0x2e, 0x53, 0x8d, 0xfe, 0xd6, 0x1f, 0x09, 0xc5, 0xe2, 0xc6, 0x13, 0xca, 0x56,
	0x30, 0x2f, 0xc4, 0xaf, 0xb3, 0xea, 0x9a, 0x54, 0x23, 0x66, 0x5c, 0xd5, 0xaa, 0x4a, 0x5b, 0x89,
	0xea, 0x4b, 0xca, 0xc1, 0x66, 0x76, 0x26, 0xec, 0x2d, 0xdc, 0xf4, 0x9d, 0x8d, 0xc2, 0x27, 0x4f,
	0x63, 0x68, 0xbe, 0x64, 0x72, 0xfe, 0x8a, 0x3f, 0x5a, 0x2c, 0x25, 0xf2, 0x39, 0x89, 0x46, 0x8c,
	0xbd, 0x87, 0x4b, 0x7b, 0x51, 0x1b, 0x7e, 0x13, 0xb8, 0xa1, 0x7f, 0xbf, 0x8c, 0xc6, 0xd7, 0x37,
	0xb2, 0x67, 0x19, 0xf7, 0xb2, 0xc7, 0xe5, 0xf7, 0xc5, 0xc3, 0xfa, 0xf3, 0xc3, 0xfa, 0x6e, 0xfc,
	0x0f, 0x24, 0x1e, 0xf5, 0x1f, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x47, 0xae, 0x71, 0x0e, 0x1c,
	0x03, 0x00, 0x00,
}
