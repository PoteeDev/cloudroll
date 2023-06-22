// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/cloudroll.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CloudrollServiceClient is the client API for CloudrollService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CloudrollServiceClient interface {
	Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EchoMessage, error)
}

type cloudrollServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCloudrollServiceClient(cc grpc.ClientConnInterface) CloudrollServiceClient {
	return &cloudrollServiceClient{cc}
}

func (c *cloudrollServiceClient) Ping(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*EchoMessage, error) {
	out := new(EchoMessage)
	err := c.cc.Invoke(ctx, "/CloudrollService/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CloudrollServiceServer is the server API for CloudrollService service.
// All implementations must embed UnimplementedCloudrollServiceServer
// for forward compatibility
type CloudrollServiceServer interface {
	Ping(context.Context, *Empty) (*EchoMessage, error)
	mustEmbedUnimplementedCloudrollServiceServer()
}

// UnimplementedCloudrollServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCloudrollServiceServer struct {
}

func (UnimplementedCloudrollServiceServer) Ping(context.Context, *Empty) (*EchoMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedCloudrollServiceServer) mustEmbedUnimplementedCloudrollServiceServer() {}

// UnsafeCloudrollServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CloudrollServiceServer will
// result in compilation errors.
type UnsafeCloudrollServiceServer interface {
	mustEmbedUnimplementedCloudrollServiceServer()
}

func RegisterCloudrollServiceServer(s grpc.ServiceRegistrar, srv CloudrollServiceServer) {
	s.RegisterService(&CloudrollService_ServiceDesc, srv)
}

func _CloudrollService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CloudrollServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/CloudrollService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CloudrollServiceServer).Ping(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// CloudrollService_ServiceDesc is the grpc.ServiceDesc for CloudrollService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CloudrollService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CloudrollService",
	HandlerType: (*CloudrollServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CloudrollService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/cloudroll.proto",
}
