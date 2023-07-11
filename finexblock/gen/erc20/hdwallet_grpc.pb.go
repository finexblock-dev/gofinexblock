// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/erc20/hdwallet.proto

package erc20

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
	HDWallet_CreateWallet_FullMethodName = "/erc20.HDWallet/CreateWallet"
	HDWallet_GetBalance_FullMethodName   = "/erc20.HDWallet/GetBalance"
)

// HDWalletClient is the client API for HDWallet service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HDWalletClient interface {
	CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*CreateWalletOutput, error)
	GetBalance(ctx context.Context, in *GetBalanceInput, opts ...grpc.CallOption) (*GetBalanceOutput, error)
}

type hDWalletClient struct {
	cc grpc.ClientConnInterface
}

func NewHDWalletClient(cc grpc.ClientConnInterface) HDWalletClient {
	return &hDWalletClient{cc}
}

func (c *hDWalletClient) CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*CreateWalletOutput, error) {
	out := new(CreateWalletOutput)
	err := c.cc.Invoke(ctx, HDWallet_CreateWallet_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hDWalletClient) GetBalance(ctx context.Context, in *GetBalanceInput, opts ...grpc.CallOption) (*GetBalanceOutput, error) {
	out := new(GetBalanceOutput)
	err := c.cc.Invoke(ctx, HDWallet_GetBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HDWalletServer is the server API for HDWallet service.
// All implementations must embed UnimplementedHDWalletServer
// for forward compatibility
type HDWalletServer interface {
	CreateWallet(context.Context, *CreateWalletInput) (*CreateWalletOutput, error)
	GetBalance(context.Context, *GetBalanceInput) (*GetBalanceOutput, error)
	mustEmbedUnimplementedHDWalletServer()
}

// UnimplementedHDWalletServer must be embedded to have forward compatible implementations.
type UnimplementedHDWalletServer struct {
}

func (UnimplementedHDWalletServer) CreateWallet(context.Context, *CreateWalletInput) (*CreateWalletOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedHDWalletServer) GetBalance(context.Context, *GetBalanceInput) (*GetBalanceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedHDWalletServer) mustEmbedUnimplementedHDWalletServer() {}

// UnsafeHDWalletServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HDWalletServer will
// result in compilation errors.
type UnsafeHDWalletServer interface {
	mustEmbedUnimplementedHDWalletServer()
}

func RegisterHDWalletServer(s grpc.ServiceRegistrar, srv HDWalletServer) {
	s.RegisterService(&HDWallet_ServiceDesc, srv)
}

func _HDWallet_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HDWalletServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HDWallet_CreateWallet_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HDWalletServer).CreateWallet(ctx, req.(*CreateWalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _HDWallet_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HDWalletServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HDWallet_GetBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HDWalletServer).GetBalance(ctx, req.(*GetBalanceInput))
	}
	return interceptor(ctx, in, info, handler)
}

// HDWallet_ServiceDesc is the grpc.ServiceDesc for HDWallet service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HDWallet_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "erc20.HDWallet",
	HandlerType: (*HDWalletServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _HDWallet_CreateWallet_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _HDWallet_GetBalance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/erc20/hdwallet.proto",
}
