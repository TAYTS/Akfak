// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/messagepb/message.proto

package messagepb

import (
	recordpb "AKFAK/proto/recordpb"
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type MessageRequest struct {
	Record               *recordpb.Record `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *MessageRequest) Reset()         { *m = MessageRequest{} }
func (m *MessageRequest) String() string { return proto.CompactTextString(m) }
func (*MessageRequest) ProtoMessage()    {}
func (*MessageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afea1b987f254e01, []int{0}
}

func (m *MessageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageRequest.Unmarshal(m, b)
}
func (m *MessageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageRequest.Marshal(b, m, deterministic)
}
func (m *MessageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageRequest.Merge(m, src)
}
func (m *MessageRequest) XXX_Size() int {
	return xxx_messageInfo_MessageRequest.Size(m)
}
func (m *MessageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageRequest proto.InternalMessageInfo

func (m *MessageRequest) GetRecord() *recordpb.Record {
	if m != nil {
		return m.Record
	}
	return nil
}

type MessageResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageResponse) Reset()         { *m = MessageResponse{} }
func (m *MessageResponse) String() string { return proto.CompactTextString(m) }
func (*MessageResponse) ProtoMessage()    {}
func (*MessageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afea1b987f254e01, []int{1}
}

func (m *MessageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageResponse.Unmarshal(m, b)
}
func (m *MessageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageResponse.Marshal(b, m, deterministic)
}
func (m *MessageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageResponse.Merge(m, src)
}
func (m *MessageResponse) XXX_Size() int {
	return xxx_messageInfo_MessageResponse.Size(m)
}
func (m *MessageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageResponse proto.InternalMessageInfo

func (m *MessageResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type MessageBatchRequest struct {
	Records              *recordpb.RecordBatch `protobuf:"bytes,1,opt,name=records,proto3" json:"records,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *MessageBatchRequest) Reset()         { *m = MessageBatchRequest{} }
func (m *MessageBatchRequest) String() string { return proto.CompactTextString(m) }
func (*MessageBatchRequest) ProtoMessage()    {}
func (*MessageBatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_afea1b987f254e01, []int{2}
}

func (m *MessageBatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageBatchRequest.Unmarshal(m, b)
}
func (m *MessageBatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageBatchRequest.Marshal(b, m, deterministic)
}
func (m *MessageBatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageBatchRequest.Merge(m, src)
}
func (m *MessageBatchRequest) XXX_Size() int {
	return xxx_messageInfo_MessageBatchRequest.Size(m)
}
func (m *MessageBatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageBatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MessageBatchRequest proto.InternalMessageInfo

func (m *MessageBatchRequest) GetRecords() *recordpb.RecordBatch {
	if m != nil {
		return m.Records
	}
	return nil
}

type MessageBatchResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MessageBatchResponse) Reset()         { *m = MessageBatchResponse{} }
func (m *MessageBatchResponse) String() string { return proto.CompactTextString(m) }
func (*MessageBatchResponse) ProtoMessage()    {}
func (*MessageBatchResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_afea1b987f254e01, []int{3}
}

func (m *MessageBatchResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageBatchResponse.Unmarshal(m, b)
}
func (m *MessageBatchResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageBatchResponse.Marshal(b, m, deterministic)
}
func (m *MessageBatchResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageBatchResponse.Merge(m, src)
}
func (m *MessageBatchResponse) XXX_Size() int {
	return xxx_messageInfo_MessageBatchResponse.Size(m)
}
func (m *MessageBatchResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageBatchResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MessageBatchResponse proto.InternalMessageInfo

func (m *MessageBatchResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*MessageRequest)(nil), "proto.messagepb.MessageRequest")
	proto.RegisterType((*MessageResponse)(nil), "proto.messagepb.MessageResponse")
	proto.RegisterType((*MessageBatchRequest)(nil), "proto.messagepb.MessageBatchRequest")
	proto.RegisterType((*MessageBatchResponse)(nil), "proto.messagepb.MessageBatchResponse")
}

func init() {
	proto.RegisterFile("proto/messagepb/message.proto", fileDescriptor_afea1b987f254e01)
}

var fileDescriptor_afea1b987f254e01 = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0x48, 0x82, 0xb1, 0xf4, 0xc0, 0xe2,
	0x42, 0xfc, 0x60, 0x4a, 0x0f, 0x2e, 0x2d, 0x25, 0x0d, 0x51, 0x5f, 0x94, 0x9a, 0x9c, 0x5f, 0x94,
	0x52, 0x90, 0x04, 0x65, 0x40, 0x54, 0x2b, 0x39, 0x70, 0xf1, 0xf9, 0x42, 0x54, 0x06, 0xa5, 0x16,
	0x96, 0xa6, 0x16, 0x97, 0x08, 0xe9, 0x71, 0xb1, 0x41, 0x54, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70,
	0x1b, 0x89, 0x41, 0x54, 0xea, 0xc1, 0xf4, 0xeb, 0x05, 0x81, 0x19, 0x41, 0x50, 0x55, 0x4a, 0x9a,
	0x5c, 0xfc, 0x70, 0x13, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0xc4, 0x40, 0x46, 0x14, 0x97,
	0xe6, 0x94, 0x80, 0x8d, 0xe0, 0x0c, 0x82, 0xf2, 0x94, 0x7c, 0xb8, 0x84, 0xa1, 0x4a, 0x9d, 0x12,
	0x4b, 0x92, 0x33, 0x60, 0x36, 0x9a, 0x72, 0xb1, 0x43, 0xcc, 0x2a, 0x86, 0x5a, 0x29, 0x8d, 0xdd,
	0x4a, 0x88, 0x26, 0x98, 0x5a, 0x25, 0x3d, 0x2e, 0x11, 0x54, 0xd3, 0xf0, 0xdb, 0x6e, 0x74, 0x9c,
	0x11, 0xee, 0xd7, 0xe0, 0xd4, 0xa2, 0xb2, 0xcc, 0xe4, 0x54, 0xa1, 0x20, 0x2e, 0x76, 0xa8, 0x88,
	0x90, 0xbc, 0x1e, 0x5a, 0xb8, 0xe9, 0xa1, 0x86, 0x8b, 0x94, 0x02, 0x6e, 0x05, 0x10, 0x8b, 0x95,
	0x18, 0x34, 0x18, 0x0d, 0x18, 0x85, 0x12, 0xb9, 0x78, 0x90, 0x9d, 0x25, 0xa4, 0x82, 0x4b, 0x1f,
	0x72, 0x18, 0x48, 0xa9, 0x12, 0x50, 0x85, 0x6c, 0x85, 0x93, 0x78, 0x94, 0xa8, 0xa3, 0xb7, 0x9b,
	0xa3, 0xb7, 0x3e, 0x5a, 0x4a, 0x48, 0x62, 0x03, 0x0b, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0xb9, 0x55, 0x0c, 0xe2, 0x23, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// MessageServiceClient is the client API for MessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MessageServiceClient interface {
	Message(ctx context.Context, opts ...grpc.CallOption) (MessageService_MessageClient, error)
	MessageBatch(ctx context.Context, opts ...grpc.CallOption) (MessageService_MessageBatchClient, error)
}

type messageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMessageServiceClient(cc grpc.ClientConnInterface) MessageServiceClient {
	return &messageServiceClient{cc}
}

func (c *messageServiceClient) Message(ctx context.Context, opts ...grpc.CallOption) (MessageService_MessageClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[0], "/proto.messagepb.MessageService/Message", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceMessageClient{stream}
	return x, nil
}

type MessageService_MessageClient interface {
	Send(*MessageRequest) error
	Recv() (*MessageResponse, error)
	grpc.ClientStream
}

type messageServiceMessageClient struct {
	grpc.ClientStream
}

func (x *messageServiceMessageClient) Send(m *MessageRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceMessageClient) Recv() (*MessageResponse, error) {
	m := new(MessageResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messageServiceClient) MessageBatch(ctx context.Context, opts ...grpc.CallOption) (MessageService_MessageBatchClient, error) {
	stream, err := c.cc.NewStream(ctx, &_MessageService_serviceDesc.Streams[1], "/proto.messagepb.MessageService/MessageBatch", opts...)
	if err != nil {
		return nil, err
	}
	x := &messageServiceMessageBatchClient{stream}
	return x, nil
}

type MessageService_MessageBatchClient interface {
	Send(*MessageBatchRequest) error
	Recv() (*MessageBatchResponse, error)
	grpc.ClientStream
}

type messageServiceMessageBatchClient struct {
	grpc.ClientStream
}

func (x *messageServiceMessageBatchClient) Send(m *MessageBatchRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *messageServiceMessageBatchClient) Recv() (*MessageBatchResponse, error) {
	m := new(MessageBatchResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// MessageServiceServer is the server API for MessageService service.
type MessageServiceServer interface {
	Message(MessageService_MessageServer) error
	MessageBatch(MessageService_MessageBatchServer) error
}

// UnimplementedMessageServiceServer can be embedded to have forward compatible implementations.
type UnimplementedMessageServiceServer struct {
}

func (*UnimplementedMessageServiceServer) Message(srv MessageService_MessageServer) error {
	return status.Errorf(codes.Unimplemented, "method Message not implemented")
}
func (*UnimplementedMessageServiceServer) MessageBatch(srv MessageService_MessageBatchServer) error {
	return status.Errorf(codes.Unimplemented, "method MessageBatch not implemented")
}

func RegisterMessageServiceServer(s *grpc.Server, srv MessageServiceServer) {
	s.RegisterService(&_MessageService_serviceDesc, srv)
}

func _MessageService_Message_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).Message(&messageServiceMessageServer{stream})
}

type MessageService_MessageServer interface {
	Send(*MessageResponse) error
	Recv() (*MessageRequest, error)
	grpc.ServerStream
}

type messageServiceMessageServer struct {
	grpc.ServerStream
}

func (x *messageServiceMessageServer) Send(m *MessageResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceMessageServer) Recv() (*MessageRequest, error) {
	m := new(MessageRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _MessageService_MessageBatch_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(MessageServiceServer).MessageBatch(&messageServiceMessageBatchServer{stream})
}

type MessageService_MessageBatchServer interface {
	Send(*MessageBatchResponse) error
	Recv() (*MessageBatchRequest, error)
	grpc.ServerStream
}

type messageServiceMessageBatchServer struct {
	grpc.ServerStream
}

func (x *messageServiceMessageBatchServer) Send(m *MessageBatchResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *messageServiceMessageBatchServer) Recv() (*MessageBatchRequest, error) {
	m := new(MessageBatchRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _MessageService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.messagepb.MessageService",
	HandlerType: (*MessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Message",
			Handler:       _MessageService_Message_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "MessageBatch",
			Handler:       _MessageService_MessageBatch_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/messagepb/message.proto",
}
