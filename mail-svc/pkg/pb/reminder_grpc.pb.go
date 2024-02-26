// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: pkg/pb/reminder.proto

package pb

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

const (
	RemainderService_TodayTask_FullMethodName = "/remainder.remainderService/TodayTask"
)

// RemainderServiceClient is the client API for RemainderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RemainderServiceClient interface {
	TodayTask(ctx context.Context, in *TodayTaskRequest, opts ...grpc.CallOption) (*TodayTaskResponse, error)
}

type remainderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRemainderServiceClient(cc grpc.ClientConnInterface) RemainderServiceClient {
	return &remainderServiceClient{cc}
}

func (c *remainderServiceClient) TodayTask(ctx context.Context, in *TodayTaskRequest, opts ...grpc.CallOption) (*TodayTaskResponse, error) {
	out := new(TodayTaskResponse)
	err := c.cc.Invoke(ctx, RemainderService_TodayTask_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RemainderServiceServer is the server API for RemainderService service.
// All implementations must embed UnimplementedRemainderServiceServer
// for forward compatibility
type RemainderServiceServer interface {
	TodayTask(context.Context, *TodayTaskRequest) (*TodayTaskResponse, error)
	mustEmbedUnimplementedRemainderServiceServer()
}

// UnimplementedRemainderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRemainderServiceServer struct {
}

func (UnimplementedRemainderServiceServer) TodayTask(context.Context, *TodayTaskRequest) (*TodayTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TodayTask not implemented")
}
func (UnimplementedRemainderServiceServer) mustEmbedUnimplementedRemainderServiceServer() {}

// UnsafeRemainderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RemainderServiceServer will
// result in compilation errors.
type UnsafeRemainderServiceServer interface {
	mustEmbedUnimplementedRemainderServiceServer()
}

func RegisterRemainderServiceServer(s grpc.ServiceRegistrar, srv RemainderServiceServer) {
	s.RegisterService(&RemainderService_ServiceDesc, srv)
}

func _RemainderService_TodayTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TodayTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RemainderServiceServer).TodayTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RemainderService_TodayTask_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RemainderServiceServer).TodayTask(ctx, req.(*TodayTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RemainderService_ServiceDesc is the grpc.ServiceDesc for RemainderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RemainderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "remainder.remainderService",
	HandlerType: (*RemainderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TodayTask",
			Handler:    _RemainderService_TodayTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/pb/reminder.proto",
}
