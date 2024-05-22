// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: internal/proto/user/user.proto

package user

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
	UserService_PostRegisterUser_FullMethodName = "/proto.UserService/PostRegisterUser"
	UserService_PostLoginUser_FullMethodName    = "/proto.UserService/PostLoginUser"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	PostRegisterUser(ctx context.Context, in *PostUserRegisterRequest, opts ...grpc.CallOption) (*PostUserRegisterResponse, error)
	PostLoginUser(ctx context.Context, in *PostUserLoginRequest, opts ...grpc.CallOption) (*PostUserLoginResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) PostRegisterUser(ctx context.Context, in *PostUserRegisterRequest, opts ...grpc.CallOption) (*PostUserRegisterResponse, error) {
	out := new(PostUserRegisterResponse)
	err := c.cc.Invoke(ctx, UserService_PostRegisterUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) PostLoginUser(ctx context.Context, in *PostUserLoginRequest, opts ...grpc.CallOption) (*PostUserLoginResponse, error) {
	out := new(PostUserLoginResponse)
	err := c.cc.Invoke(ctx, UserService_PostLoginUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	PostRegisterUser(context.Context, *PostUserRegisterRequest) (*PostUserRegisterResponse, error)
	PostLoginUser(context.Context, *PostUserLoginRequest) (*PostUserLoginResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) PostRegisterUser(context.Context, *PostUserRegisterRequest) (*PostUserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostRegisterUser not implemented")
}
func (UnimplementedUserServiceServer) PostLoginUser(context.Context, *PostUserLoginRequest) (*PostUserLoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostLoginUser not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_PostRegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserRegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PostRegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_PostRegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PostRegisterUser(ctx, req.(*PostUserRegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_PostLoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostUserLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).PostLoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_PostLoginUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).PostLoginUser(ctx, req.(*PostUserLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostRegisterUser",
			Handler:    _UserService_PostRegisterUser_Handler,
		},
		{
			MethodName: "PostLoginUser",
			Handler:    _UserService_PostLoginUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/user/user.proto",
}
