// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: internal/proto/credit_card/credit_card.proto

package credit_card

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
	CreditCardService_PostSaveCreditCard_FullMethodName = "/proto.CreditCardService/PostSaveCreditCard"
	CreditCardService_GetLoadCreditCard_FullMethodName  = "/proto.CreditCardService/GetLoadCreditCard"
)

// CreditCardServiceClient is the client API for CreditCardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreditCardServiceClient interface {
	PostSaveCreditCard(ctx context.Context, in *PostCreditCardRequest, opts ...grpc.CallOption) (*PostCreditCardResponse, error)
	GetLoadCreditCard(ctx context.Context, in *GetCreditCardRequest, opts ...grpc.CallOption) (*GetCreditCardResponse, error)
}

type creditCardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreditCardServiceClient(cc grpc.ClientConnInterface) CreditCardServiceClient {
	return &creditCardServiceClient{cc}
}

func (c *creditCardServiceClient) PostSaveCreditCard(ctx context.Context, in *PostCreditCardRequest, opts ...grpc.CallOption) (*PostCreditCardResponse, error) {
	out := new(PostCreditCardResponse)
	err := c.cc.Invoke(ctx, CreditCardService_PostSaveCreditCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardServiceClient) GetLoadCreditCard(ctx context.Context, in *GetCreditCardRequest, opts ...grpc.CallOption) (*GetCreditCardResponse, error) {
	out := new(GetCreditCardResponse)
	err := c.cc.Invoke(ctx, CreditCardService_GetLoadCreditCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreditCardServiceServer is the server API for CreditCardService service.
// All implementations must embed UnimplementedCreditCardServiceServer
// for forward compatibility
type CreditCardServiceServer interface {
	PostSaveCreditCard(context.Context, *PostCreditCardRequest) (*PostCreditCardResponse, error)
	GetLoadCreditCard(context.Context, *GetCreditCardRequest) (*GetCreditCardResponse, error)
	mustEmbedUnimplementedCreditCardServiceServer()
}

// UnimplementedCreditCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCreditCardServiceServer struct {
}

func (UnimplementedCreditCardServiceServer) PostSaveCreditCard(context.Context, *PostCreditCardRequest) (*PostCreditCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostSaveCreditCard not implemented")
}
func (UnimplementedCreditCardServiceServer) GetLoadCreditCard(context.Context, *GetCreditCardRequest) (*GetCreditCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLoadCreditCard not implemented")
}
func (UnimplementedCreditCardServiceServer) mustEmbedUnimplementedCreditCardServiceServer() {}

// UnsafeCreditCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreditCardServiceServer will
// result in compilation errors.
type UnsafeCreditCardServiceServer interface {
	mustEmbedUnimplementedCreditCardServiceServer()
}

func RegisterCreditCardServiceServer(s grpc.ServiceRegistrar, srv CreditCardServiceServer) {
	s.RegisterService(&CreditCardService_ServiceDesc, srv)
}

func _CreditCardService_PostSaveCreditCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCreditCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardServiceServer).PostSaveCreditCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreditCardService_PostSaveCreditCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardServiceServer).PostSaveCreditCard(ctx, req.(*PostCreditCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardService_GetLoadCreditCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCreditCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardServiceServer).GetLoadCreditCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CreditCardService_GetLoadCreditCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardServiceServer).GetLoadCreditCard(ctx, req.(*GetCreditCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CreditCardService_ServiceDesc is the grpc.ServiceDesc for CreditCardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreditCardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.CreditCardService",
	HandlerType: (*CreditCardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSaveCreditCard",
			Handler:    _CreditCardService_PostSaveCreditCard_Handler,
		},
		{
			MethodName: "GetLoadCreditCard",
			Handler:    _CreditCardService_GetLoadCreditCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/proto/credit_card/credit_card.proto",
}
