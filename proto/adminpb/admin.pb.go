// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/adminpb/admin.proto

package adminpb

import (
	adminclientpb "AKFAK/proto/adminclientpb"
	heartbeatspb "AKFAK/proto/heartbeatspb"
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
	proto.RegisterFile("proto/adminpb/admin.proto", fileDescriptor_b89f257325c17cfa)
}

var fileDescriptor_b89f257325c17cfa = []byte{
	// 383 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x94, 0xcb, 0x4e, 0xf3, 0x30,
	0x10, 0x85, 0xff, 0x6e, 0xba, 0xb0, 0xda, 0x7f, 0xe1, 0x0a, 0x21, 0xba, 0x04, 0x81, 0xda, 0x22,
	0xd2, 0x8a, 0xf6, 0x05, 0x42, 0xc5, 0x4d, 0x05, 0x84, 0xb8, 0x6c, 0xd8, 0x58, 0x4e, 0x32, 0xa2,
	0x11, 0x69, 0x12, 0x9c, 0x29, 0x95, 0x78, 0x00, 0xde, 0x86, 0x77, 0x44, 0x75, 0x6c, 0x37, 0x11,
	0x16, 0x98, 0x95, 0x35, 0xf1, 0x77, 0xce, 0x99, 0x38, 0x13, 0x93, 0x9d, 0x5c, 0x64, 0x98, 0x0d,
	0x79, 0xb4, 0x88, 0xd3, 0x3c, 0x28, 0x57, 0x4f, 0x3e, 0xa3, 0x6d, 0xb9, 0x78, 0x6a, 0xab, 0x3b,
	0xa9, 0x90, 0x61, 0x12, 0x43, 0x8a, 0x9a, 0x67, 0x65, 0xc9, 0x52, 0x58, 0xb1, 0x9c, 0x0b, 0x8c,
	0x31, 0xce, 0x94, 0x49, 0x77, 0xe4, 0xa4, 0xc2, 0x2c, 0x8f, 0x43, 0xa5, 0x38, 0xb2, 0x29, 0xc2,
	0x2c, 0x45, 0x91, 0x25, 0x09, 0x08, 0x06, 0x09, 0x84, 0x95, 0x80, 0x9e, 0x0d, 0x4f, 0x80, 0x47,
	0x20, 0x18, 0x4f, 0x23, 0x16, 0x17, 0x42, 0x91, 0x7d, 0x1b, 0xb9, 0xcc, 0x23, 0x8e, 0xc0, 0x16,
	0x80, 0x3c, 0xe2, 0xc8, 0x7f, 0x32, 0x7d, 0x06, 0x64, 0x9b, 0x3e, 0x14, 0xb9, 0x57, 0x92, 0x73,
	0xe0, 0x02, 0x03, 0xe0, 0x58, 0xe4, 0x41, 0xa5, 0x28, 0xa1, 0xe3, 0xcf, 0x26, 0x69, 0xf9, 0x6b,
	0xaf, 0x7b, 0x10, 0x6f, 0x71, 0x08, 0x74, 0x45, 0xe8, 0xd4, 0x38, 0x9d, 0xaa, 0x17, 0xa2, 0x9e,
	0x57, 0x39, 0x71, 0x1d, 0xeb, 0x7d, 0x07, 0xef, 0xe0, 0x75, 0x09, 0x05, 0x76, 0x87, 0xce, 0x7c,
	0x91, 0x67, 0x69, 0x01, 0xbb, 0xff, 0xe8, 0x3b, 0xe9, 0xc8, 0x46, 0xa6, 0x92, 0xbe, 0x81, 0xd5,
	0xc3, 0xfa, 0xe4, 0xa9, 0xdd, 0xc9, 0x42, 0xea, 0xe8, 0x91, 0xbb, 0xc0, 0x64, 0x7f, 0x34, 0xc8,
	0x76, 0x9d, 0xb8, 0xd5, 0xc3, 0x42, 0xc7, 0x0e, 0x7e, 0x86, 0xd6, 0x4d, 0x4c, 0xfe, 0x26, 0x32,
	0x8d, 0x00, 0x69, 0x5d, 0xc9, 0x01, 0xf1, 0xd3, 0xe8, 0xb2, 0x10, 0xb4, 0x67, 0xf5, 0xa9, 0x22,
	0x3a, 0xb1, 0xef, 0x40, 0x9a, 0x98, 0x17, 0xf2, 0xff, 0x51, 0x4e, 0xd7, 0xb5, 0x1a, 0x2e, 0x3a,
	0xb0, 0xca, 0xeb, 0x90, 0x8e, 0x3a, 0x74, 0x62, 0x4d, 0xd8, 0x9c, 0xb4, 0xcf, 0x01, 0x37, 0xdf,
	0x9e, 0xda, 0x5b, 0xad, 0x31, 0x3a, 0x6a, 0xe0, 0x82, 0x9a, 0x24, 0x4e, 0xc8, 0x85, 0x19, 0x70,
	0xba, 0xaf, 0xb4, 0xd5, 0x1f, 0xc0, 0xdb, 0xec, 0xeb, 0x88, 0x83, 0xdf, 0x30, 0x6d, 0xdf, 0x6b,
	0x8c, 0x1a, 0x27, 0x5b, 0x4f, 0x1d, 0x7f, 0x76, 0xe6, 0xcf, 0x86, 0xb5, 0xcb, 0x29, 0x68, 0xca,
	0x72, 0xfc, 0x15, 0x00, 0x00, 0xff, 0xff, 0x7a, 0xe1, 0xc2, 0xc5, 0xb4, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AdminServiceClient is the client API for AdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AdminServiceClient interface {
	ControllerElection(ctx context.Context, in *adminclientpb.ControllerElectionRequest, opts ...grpc.CallOption) (*adminclientpb.ControllerElectionResponse, error)
	AdminClientNewTopic(ctx context.Context, in *adminclientpb.AdminClientNewTopicRequest, opts ...grpc.CallOption) (*adminclientpb.AdminClientNewTopicResponse, error)
	AdminClientNewPartition(ctx context.Context, in *adminclientpb.AdminClientNewPartitionRequest, opts ...grpc.CallOption) (*adminclientpb.AdminClientNewPartitionResponse, error)
	LeaderAndIsr(ctx context.Context, in *adminclientpb.LeaderAndIsrRequest, opts ...grpc.CallOption) (*adminclientpb.LeaderAndIsrResponse, error)
	UpdateMetadata(ctx context.Context, in *adminclientpb.UpdateMetadataRequest, opts ...grpc.CallOption) (*adminclientpb.UpdateMetadataResponse, error)
	GetController(ctx context.Context, in *adminclientpb.GetControllerRequest, opts ...grpc.CallOption) (*adminclientpb.GetControllerResponse, error)
	Heartbeats(ctx context.Context, opts ...grpc.CallOption) (AdminService_HeartbeatsClient, error)
}

type adminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAdminServiceClient(cc grpc.ClientConnInterface) AdminServiceClient {
	return &adminServiceClient{cc}
}

func (c *adminServiceClient) ControllerElection(ctx context.Context, in *adminclientpb.ControllerElectionRequest, opts ...grpc.CallOption) (*adminclientpb.ControllerElectionResponse, error) {
	out := new(adminclientpb.ControllerElectionResponse)
	err := c.cc.Invoke(ctx, "/proto.adminpb.AdminService/ControllerElection", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminClientNewTopic(ctx context.Context, in *adminclientpb.AdminClientNewTopicRequest, opts ...grpc.CallOption) (*adminclientpb.AdminClientNewTopicResponse, error) {
	out := new(adminclientpb.AdminClientNewTopicResponse)
	err := c.cc.Invoke(ctx, "/proto.adminpb.AdminService/AdminClientNewTopic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) AdminClientNewPartition(ctx context.Context, in *adminclientpb.AdminClientNewPartitionRequest, opts ...grpc.CallOption) (*adminclientpb.AdminClientNewPartitionResponse, error) {
	out := new(adminclientpb.AdminClientNewPartitionResponse)
	err := c.cc.Invoke(ctx, "/proto.adminpb.AdminService/AdminClientNewPartition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) LeaderAndIsr(ctx context.Context, in *adminclientpb.LeaderAndIsrRequest, opts ...grpc.CallOption) (*adminclientpb.LeaderAndIsrResponse, error) {
	out := new(adminclientpb.LeaderAndIsrResponse)
	err := c.cc.Invoke(ctx, "/proto.adminpb.AdminService/LeaderAndIsr", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) UpdateMetadata(ctx context.Context, in *adminclientpb.UpdateMetadataRequest, opts ...grpc.CallOption) (*adminclientpb.UpdateMetadataResponse, error) {
	out := new(adminclientpb.UpdateMetadataResponse)
	err := c.cc.Invoke(ctx, "/proto.adminpb.AdminService/UpdateMetadata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) GetController(ctx context.Context, in *adminclientpb.GetControllerRequest, opts ...grpc.CallOption) (*adminclientpb.GetControllerResponse, error) {
	out := new(adminclientpb.GetControllerResponse)
	err := c.cc.Invoke(ctx, "/proto.adminpb.AdminService/GetController", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *adminServiceClient) Heartbeats(ctx context.Context, opts ...grpc.CallOption) (AdminService_HeartbeatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_AdminService_serviceDesc.Streams[0], "/proto.adminpb.AdminService/Heartbeats", opts...)
	if err != nil {
		return nil, err
	}
	x := &adminServiceHeartbeatsClient{stream}
	return x, nil
}

type AdminService_HeartbeatsClient interface {
	Send(*heartbeatspb.HeartbeatsRequest) error
	Recv() (*heartbeatspb.HeartbeatsResponse, error)
	grpc.ClientStream
}

type adminServiceHeartbeatsClient struct {
	grpc.ClientStream
}

func (x *adminServiceHeartbeatsClient) Send(m *heartbeatspb.HeartbeatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *adminServiceHeartbeatsClient) Recv() (*heartbeatspb.HeartbeatsResponse, error) {
	m := new(heartbeatspb.HeartbeatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// AdminServiceServer is the server API for AdminService service.
type AdminServiceServer interface {
	ControllerElection(context.Context, *adminclientpb.ControllerElectionRequest) (*adminclientpb.ControllerElectionResponse, error)
	AdminClientNewTopic(context.Context, *adminclientpb.AdminClientNewTopicRequest) (*adminclientpb.AdminClientNewTopicResponse, error)
	AdminClientNewPartition(context.Context, *adminclientpb.AdminClientNewPartitionRequest) (*adminclientpb.AdminClientNewPartitionResponse, error)
	LeaderAndIsr(context.Context, *adminclientpb.LeaderAndIsrRequest) (*adminclientpb.LeaderAndIsrResponse, error)
	UpdateMetadata(context.Context, *adminclientpb.UpdateMetadataRequest) (*adminclientpb.UpdateMetadataResponse, error)
	GetController(context.Context, *adminclientpb.GetControllerRequest) (*adminclientpb.GetControllerResponse, error)
	Heartbeats(AdminService_HeartbeatsServer) error
}

// UnimplementedAdminServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAdminServiceServer struct {
}

func (*UnimplementedAdminServiceServer) ControllerElection(ctx context.Context, req *adminclientpb.ControllerElectionRequest) (*adminclientpb.ControllerElectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ControllerElection not implemented")
}
func (*UnimplementedAdminServiceServer) AdminClientNewTopic(ctx context.Context, req *adminclientpb.AdminClientNewTopicRequest) (*adminclientpb.AdminClientNewTopicResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminClientNewTopic not implemented")
}
func (*UnimplementedAdminServiceServer) AdminClientNewPartition(ctx context.Context, req *adminclientpb.AdminClientNewPartitionRequest) (*adminclientpb.AdminClientNewPartitionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AdminClientNewPartition not implemented")
}
func (*UnimplementedAdminServiceServer) LeaderAndIsr(ctx context.Context, req *adminclientpb.LeaderAndIsrRequest) (*adminclientpb.LeaderAndIsrResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaderAndIsr not implemented")
}
func (*UnimplementedAdminServiceServer) UpdateMetadata(ctx context.Context, req *adminclientpb.UpdateMetadataRequest) (*adminclientpb.UpdateMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMetadata not implemented")
}
func (*UnimplementedAdminServiceServer) GetController(ctx context.Context, req *adminclientpb.GetControllerRequest) (*adminclientpb.GetControllerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetController not implemented")
}
func (*UnimplementedAdminServiceServer) Heartbeats(srv AdminService_HeartbeatsServer) error {
	return status.Errorf(codes.Unimplemented, "method Heartbeats not implemented")
}

func RegisterAdminServiceServer(s *grpc.Server, srv AdminServiceServer) {
	s.RegisterService(&_AdminService_serviceDesc, srv)
}

func _AdminService_ControllerElection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adminclientpb.ControllerElectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).ControllerElection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.adminpb.AdminService/ControllerElection",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).ControllerElection(ctx, req.(*adminclientpb.ControllerElectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminClientNewTopic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adminclientpb.AdminClientNewTopicRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminClientNewTopic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.adminpb.AdminService/AdminClientNewTopic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminClientNewTopic(ctx, req.(*adminclientpb.AdminClientNewTopicRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_AdminClientNewPartition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adminclientpb.AdminClientNewPartitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).AdminClientNewPartition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.adminpb.AdminService/AdminClientNewPartition",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).AdminClientNewPartition(ctx, req.(*adminclientpb.AdminClientNewPartitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_LeaderAndIsr_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adminclientpb.LeaderAndIsrRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).LeaderAndIsr(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.adminpb.AdminService/LeaderAndIsr",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).LeaderAndIsr(ctx, req.(*adminclientpb.LeaderAndIsrRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_UpdateMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adminclientpb.UpdateMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).UpdateMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.adminpb.AdminService/UpdateMetadata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).UpdateMetadata(ctx, req.(*adminclientpb.UpdateMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_GetController_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(adminclientpb.GetControllerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AdminServiceServer).GetController(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.adminpb.AdminService/GetController",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AdminServiceServer).GetController(ctx, req.(*adminclientpb.GetControllerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AdminService_Heartbeats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AdminServiceServer).Heartbeats(&adminServiceHeartbeatsServer{stream})
}

type AdminService_HeartbeatsServer interface {
	Send(*heartbeatspb.HeartbeatsResponse) error
	Recv() (*heartbeatspb.HeartbeatsRequest, error)
	grpc.ServerStream
}

type adminServiceHeartbeatsServer struct {
	grpc.ServerStream
}

func (x *adminServiceHeartbeatsServer) Send(m *heartbeatspb.HeartbeatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *adminServiceHeartbeatsServer) Recv() (*heartbeatspb.HeartbeatsRequest, error) {
	m := new(heartbeatspb.HeartbeatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _AdminService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.adminpb.AdminService",
	HandlerType: (*AdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ControllerElection",
			Handler:    _AdminService_ControllerElection_Handler,
		},
		{
			MethodName: "AdminClientNewTopic",
			Handler:    _AdminService_AdminClientNewTopic_Handler,
		},
		{
			MethodName: "AdminClientNewPartition",
			Handler:    _AdminService_AdminClientNewPartition_Handler,
		},
		{
			MethodName: "LeaderAndIsr",
			Handler:    _AdminService_LeaderAndIsr_Handler,
		},
		{
			MethodName: "UpdateMetadata",
			Handler:    _AdminService_UpdateMetadata_Handler,
		},
		{
			MethodName: "GetController",
			Handler:    _AdminService_GetController_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Heartbeats",
			Handler:       _AdminService_Heartbeats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/adminpb/admin.proto",
}
