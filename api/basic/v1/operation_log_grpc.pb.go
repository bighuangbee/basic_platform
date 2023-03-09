// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// OperationLogClient is the client API for OperationLog service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OperationLogClient interface {
	// 增加操作日志.
	Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error)
	// 查询操作日志.
	ListOperationLog(ctx context.Context, in *ListOperationLogRequest, opts ...grpc.CallOption) (*ListOperationLogReply, error)
	// 操作日志操作人搜索.
	ListOperationLogUser(ctx context.Context, in *ListOperationLogUserRequest, opts ...grpc.CallOption) (*ListOperationLogUserReply, error)
}

type operationLogClient struct {
	cc grpc.ClientConnInterface
}

func NewOperationLogClient(cc grpc.ClientConnInterface) OperationLogClient {
	return &operationLogClient{cc}
}

func (c *operationLogClient) Add(ctx context.Context, in *AddRequest, opts ...grpc.CallOption) (*AddReply, error) {
	out := new(AddReply)
	err := c.cc.Invoke(ctx, "/api.basic.v1.OperationLog/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationLogClient) ListOperationLog(ctx context.Context, in *ListOperationLogRequest, opts ...grpc.CallOption) (*ListOperationLogReply, error) {
	out := new(ListOperationLogReply)
	err := c.cc.Invoke(ctx, "/api.basic.v1.OperationLog/ListOperationLog", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *operationLogClient) ListOperationLogUser(ctx context.Context, in *ListOperationLogUserRequest, opts ...grpc.CallOption) (*ListOperationLogUserReply, error) {
	out := new(ListOperationLogUserReply)
	err := c.cc.Invoke(ctx, "/api.basic.v1.OperationLog/ListOperationLogUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OperationLogServer is the server API for OperationLog service.
// All implementations must embed UnimplementedOperationLogServer
// for forward compatibility
type OperationLogServer interface {
	// 增加操作日志.
	Add(context.Context, *AddRequest) (*AddReply, error)
	// 查询操作日志.
	ListOperationLog(context.Context, *ListOperationLogRequest) (*ListOperationLogReply, error)
	// 操作日志操作人搜索.
	ListOperationLogUser(context.Context, *ListOperationLogUserRequest) (*ListOperationLogUserReply, error)
	mustEmbedUnimplementedOperationLogServer()
}

// UnimplementedOperationLogServer must be embedded to have forward compatible implementations.
type UnimplementedOperationLogServer struct {
}

func (UnimplementedOperationLogServer) Add(context.Context, *AddRequest) (*AddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedOperationLogServer) ListOperationLog(context.Context, *ListOperationLogRequest) (*ListOperationLogReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperationLog not implemented")
}
func (UnimplementedOperationLogServer) ListOperationLogUser(context.Context, *ListOperationLogUserRequest) (*ListOperationLogUserReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListOperationLogUser not implemented")
}
func (UnimplementedOperationLogServer) mustEmbedUnimplementedOperationLogServer() {}

// UnsafeOperationLogServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OperationLogServer will
// result in compilation errors.
type UnsafeOperationLogServer interface {
	mustEmbedUnimplementedOperationLogServer()
}

func RegisterOperationLogServer(s grpc.ServiceRegistrar, srv OperationLogServer) {
	s.RegisterService(&OperationLog_ServiceDesc, srv)
}

func _OperationLog_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.basic.v1.OperationLog/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServer).Add(ctx, req.(*AddRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationLog_ListOperationLog_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServer).ListOperationLog(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.basic.v1.OperationLog/ListOperationLog",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServer).ListOperationLog(ctx, req.(*ListOperationLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OperationLog_ListOperationLogUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListOperationLogUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OperationLogServer).ListOperationLogUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.basic.v1.OperationLog/ListOperationLogUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OperationLogServer).ListOperationLogUser(ctx, req.(*ListOperationLogUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OperationLog_ServiceDesc is the grpc.ServiceDesc for OperationLog service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OperationLog_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.basic.v1.OperationLog",
	HandlerType: (*OperationLogServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _OperationLog_Add_Handler,
		},
		{
			MethodName: "ListOperationLog",
			Handler:    _OperationLog_ListOperationLog_Handler,
		},
		{
			MethodName: "ListOperationLogUser",
			Handler:    _OperationLog_ListOperationLogUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/basic/v1/operation_log.proto",
}
