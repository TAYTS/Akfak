// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/clientpb/client.proto

package clientpb

import (
	consumepb "AKFAK/proto/consumepb"
	metadatapb "AKFAK/proto/metadatapb"
	producepb "AKFAK/proto/producepb"
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

func init() {
	proto.RegisterFile("proto/clientpb/client.proto", fileDescriptor_e0b4fa04edbad8e6)
}

var fileDescriptor_e0b4fa04edbad8e6 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x29, 0x48, 0x82, 0x32, 0xf4, 0xc0, 0xa2, 0x42,
	0x7c, 0x60, 0x4a, 0x0f, 0x26, 0x29, 0x25, 0x0b, 0x51, 0x5c, 0x50, 0x94, 0x9f, 0x52, 0x9a, 0x9c,
	0x5a, 0x90, 0x04, 0x63, 0x41, 0x94, 0x4b, 0xc9, 0x43, 0xa4, 0x73, 0x53, 0x4b, 0x12, 0x53, 0x12,
	0x4b, 0x12, 0x0b, 0x92, 0xe0, 0x4c, 0xa8, 0x02, 0xa8, 0xfe, 0xe4, 0xfc, 0xbc, 0xe2, 0xd2, 0x5c,
	0x90, 0x7e, 0x28, 0x0b, 0x22, 0x6d, 0x34, 0x89, 0x89, 0x8b, 0xd7, 0x19, 0x6c, 0x57, 0x70, 0x6a,
	0x51, 0x59, 0x66, 0x72, 0xaa, 0x50, 0x24, 0x17, 0x5f, 0x78, 0x62, 0x66, 0x89, 0x7f, 0x9e, 0x2f,
	0xd4, 0x20, 0x21, 0x45, 0x88, 0x5a, 0x3d, 0x84, 0x25, 0x7a, 0x30, 0xb9, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x29, 0x25, 0x7c, 0x4a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x18, 0x84,
	0x82, 0xb8, 0xd8, 0x03, 0x20, 0xae, 0x17, 0x92, 0x87, 0x6a, 0x80, 0xfb, 0x4b, 0x0f, 0x2a, 0x03,
	0x33, 0x51, 0x01, 0xb7, 0x02, 0x98, 0x79, 0x1a, 0x8c, 0x06, 0x8c, 0x20, 0x33, 0x9d, 0x21, 0x3e,
	0x82, 0x9b, 0x09, 0xf7, 0xab, 0x1e, 0x54, 0x06, 0xdd, 0x4c, 0x2c, 0x0a, 0x90, 0xcd, 0x74, 0x12,
	0x8b, 0x12, 0x71, 0xf4, 0x76, 0x73, 0xf4, 0xd6, 0x47, 0x8d, 0xa8, 0x24, 0x36, 0x30, 0xdf, 0x18,
	0x10, 0x00, 0x00, 0xff, 0xff, 0xd7, 0x03, 0x43, 0xa7, 0xc1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ClientServiceClient is the client API for ClientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ClientServiceClient interface {
	// Shared RPC
	WaitOnMetadata(ctx context.Context, in *metadatapb.MetadataRequest, opts ...grpc.CallOption) (*metadatapb.MetadataResponse, error)
	// Producer RPC
	Produce(ctx context.Context, opts ...grpc.CallOption) (ClientService_ProduceClient, error)
	// Consumer RPC
	Consume(ctx context.Context, opts ...grpc.CallOption) (ClientService_ConsumeClient, error)
}

type clientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientServiceClient(cc grpc.ClientConnInterface) ClientServiceClient {
	return &clientServiceClient{cc}
}

func (c *clientServiceClient) WaitOnMetadata(ctx context.Context, in *metadatapb.MetadataRequest, opts ...grpc.CallOption) (*metadatapb.MetadataResponse, error) {
	out := new(metadatapb.MetadataResponse)
	err := c.cc.Invoke(ctx, "/proto.clientpb.ClientService/WaitOnMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clientServiceClient) Produce(ctx context.Context, opts ...grpc.CallOption) (ClientService_ProduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientService_serviceDesc.Streams[0], "/proto.clientpb.ClientService/Produce", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceProduceClient{stream}
	return x, nil
}

type ClientService_ProduceClient interface {
	Send(*producepb.ProduceRequest) error
	Recv() (*producepb.ProduceResponse, error)
	grpc.ClientStream
}

type clientServiceProduceClient struct {
	grpc.ClientStream
}

func (x *clientServiceProduceClient) Send(m *producepb.ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServiceProduceClient) Recv() (*producepb.ProduceResponse, error) {
	m := new(producepb.ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *clientServiceClient) Consume(ctx context.Context, opts ...grpc.CallOption) (ClientService_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ClientService_serviceDesc.Streams[1], "/proto.clientpb.ClientService/Consume", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientServiceConsumeClient{stream}
	return x, nil
}

type ClientService_ConsumeClient interface {
	Send(*consumepb.ConsumeRequest) error
	Recv() (*consumepb.ConsumeResponse, error)
	grpc.ClientStream
}

type clientServiceConsumeClient struct {
	grpc.ClientStream
}

func (x *clientServiceConsumeClient) Send(m *consumepb.ConsumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientServiceConsumeClient) Recv() (*consumepb.ConsumeResponse, error) {
	m := new(consumepb.ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientServiceServer is the server API for ClientService service.
type ClientServiceServer interface {
	// Shared RPC
	WaitOnMetadata(context.Context, *metadatapb.MetadataRequest) (*metadatapb.MetadataResponse, error)
	// Producer RPC
	Produce(ClientService_ProduceServer) error
	// Consumer RPC
	Consume(ClientService_ConsumeServer) error
}

// UnimplementedClientServiceServer can be embedded to have forward compatible implementations.
type UnimplementedClientServiceServer struct {
}

func (*UnimplementedClientServiceServer) WaitOnMetadata(ctx context.Context, req *metadatapb.MetadataRequest) (*metadatapb.MetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitOnMetadata not implemented")
}
func (*UnimplementedClientServiceServer) Produce(srv ClientService_ProduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (*UnimplementedClientServiceServer) Consume(srv ClientService_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}

func RegisterClientServiceServer(s *grpc.Server, srv ClientServiceServer) {
	s.RegisterService(&_ClientService_serviceDesc, srv)
}

func _ClientService_WaitOnMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(metadatapb.MetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClientServiceServer).WaitOnMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.clientpb.ClientService/WaitOnMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClientServiceServer).WaitOnMetadata(ctx, req.(*metadatapb.MetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClientService_Produce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).Produce(&clientServiceProduceServer{stream})
}

type ClientService_ProduceServer interface {
	Send(*producepb.ProduceResponse) error
	Recv() (*producepb.ProduceRequest, error)
	grpc.ServerStream
}

type clientServiceProduceServer struct {
	grpc.ServerStream
}

func (x *clientServiceProduceServer) Send(m *producepb.ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServiceProduceServer) Recv() (*producepb.ProduceRequest, error) {
	m := new(producepb.ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ClientService_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientServiceServer).Consume(&clientServiceConsumeServer{stream})
}

type ClientService_ConsumeServer interface {
	Send(*consumepb.ConsumeResponse) error
	Recv() (*consumepb.ConsumeRequest, error)
	grpc.ServerStream
}

type clientServiceConsumeServer struct {
	grpc.ServerStream
}

func (x *clientServiceConsumeServer) Send(m *consumepb.ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientServiceConsumeServer) Recv() (*consumepb.ConsumeRequest, error) {
	m := new(consumepb.ConsumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ClientService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.clientpb.ClientService",
	HandlerType: (*ClientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WaitOnMetadata",
			Handler:    _ClientService_WaitOnMetadata_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Produce",
			Handler:       _ClientService_Produce_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Consume",
			Handler:       _ClientService_Consume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/clientpb/client.proto",
}
