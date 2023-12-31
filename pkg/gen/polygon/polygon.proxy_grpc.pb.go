// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.1
// source: pkg/proto/polygon/polygon.proxy.proto

package polygon

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

// PolygonProxyClient is the client API for PolygonProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolygonProxyClient interface {
	CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*CreateWalletOutput, error)
	GetBalance(ctx context.Context, in *GetBalanceInput, opts ...grpc.CallOption) (*GetBalanceOutput, error)
	GetReceipt(ctx context.Context, in *GetReceiptInput, opts ...grpc.CallOption) (*GetReceiptOutput, error)
	SendRawTransaction(ctx context.Context, in *SendRawTransactionInput, opts ...grpc.CallOption) (*SendRawTransactionOutput, error)
	CreateRawTransaction(ctx context.Context, in *CreateRawTransactionInput, opts ...grpc.CallOption) (*CreateRawTransactionOutput, error)
	GetBlockNumber(ctx context.Context, in *GetBlockNumberInput, opts ...grpc.CallOption) (*GetBlockNumberOutput, error)
	GetBlocks(ctx context.Context, in *GetBlocksInput, opts ...grpc.CallOption) (*GetBlocksOutput, error)
}

type polygonProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewPolygonProxyClient(cc grpc.ClientConnInterface) PolygonProxyClient {
	return &polygonProxyClient{cc}
}

func (c *polygonProxyClient) CreateWallet(ctx context.Context, in *CreateWalletInput, opts ...grpc.CallOption) (*CreateWalletOutput, error) {
	out := new(CreateWalletOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/CreateWallet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetBalance(ctx context.Context, in *GetBalanceInput, opts ...grpc.CallOption) (*GetBalanceOutput, error) {
	out := new(GetBalanceOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/GetBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetReceipt(ctx context.Context, in *GetReceiptInput, opts ...grpc.CallOption) (*GetReceiptOutput, error) {
	out := new(GetReceiptOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/GetReceipt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) SendRawTransaction(ctx context.Context, in *SendRawTransactionInput, opts ...grpc.CallOption) (*SendRawTransactionOutput, error) {
	out := new(SendRawTransactionOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/SendRawTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) CreateRawTransaction(ctx context.Context, in *CreateRawTransactionInput, opts ...grpc.CallOption) (*CreateRawTransactionOutput, error) {
	out := new(CreateRawTransactionOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/CreateRawTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetBlockNumber(ctx context.Context, in *GetBlockNumberInput, opts ...grpc.CallOption) (*GetBlockNumberOutput, error) {
	out := new(GetBlockNumberOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/GetBlockNumber", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polygonProxyClient) GetBlocks(ctx context.Context, in *GetBlocksInput, opts ...grpc.CallOption) (*GetBlocksOutput, error) {
	out := new(GetBlocksOutput)
	err := c.cc.Invoke(ctx, "/polygon.PolygonProxy/GetBlocks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolygonProxyServer is the server API for PolygonProxy service.
// All implementations must embed UnimplementedPolygonProxyServer
// for forward compatibility
type PolygonProxyServer interface {
	CreateWallet(context.Context, *CreateWalletInput) (*CreateWalletOutput, error)
	GetBalance(context.Context, *GetBalanceInput) (*GetBalanceOutput, error)
	GetReceipt(context.Context, *GetReceiptInput) (*GetReceiptOutput, error)
	SendRawTransaction(context.Context, *SendRawTransactionInput) (*SendRawTransactionOutput, error)
	CreateRawTransaction(context.Context, *CreateRawTransactionInput) (*CreateRawTransactionOutput, error)
	GetBlockNumber(context.Context, *GetBlockNumberInput) (*GetBlockNumberOutput, error)
	GetBlocks(context.Context, *GetBlocksInput) (*GetBlocksOutput, error)
	mustEmbedUnimplementedPolygonProxyServer()
}

// UnimplementedPolygonProxyServer must be embedded to have forward compatible implementations.
type UnimplementedPolygonProxyServer struct {
}

func (UnimplementedPolygonProxyServer) CreateWallet(context.Context, *CreateWalletInput) (*CreateWalletOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWallet not implemented")
}
func (UnimplementedPolygonProxyServer) GetBalance(context.Context, *GetBalanceInput) (*GetBalanceOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBalance not implemented")
}
func (UnimplementedPolygonProxyServer) GetReceipt(context.Context, *GetReceiptInput) (*GetReceiptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedPolygonProxyServer) SendRawTransaction(context.Context, *SendRawTransactionInput) (*SendRawTransactionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRawTransaction not implemented")
}
func (UnimplementedPolygonProxyServer) CreateRawTransaction(context.Context, *CreateRawTransactionInput) (*CreateRawTransactionOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRawTransaction not implemented")
}
func (UnimplementedPolygonProxyServer) GetBlockNumber(context.Context, *GetBlockNumberInput) (*GetBlockNumberOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockNumber not implemented")
}
func (UnimplementedPolygonProxyServer) GetBlocks(context.Context, *GetBlocksInput) (*GetBlocksOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedPolygonProxyServer) mustEmbedUnimplementedPolygonProxyServer() {}

// UnsafePolygonProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolygonProxyServer will
// result in compilation errors.
type UnsafePolygonProxyServer interface {
	mustEmbedUnimplementedPolygonProxyServer()
}

func RegisterPolygonProxyServer(s grpc.ServiceRegistrar, srv PolygonProxyServer) {
	s.RegisterService(&PolygonProxy_ServiceDesc, srv)
}

func _PolygonProxy_CreateWallet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).CreateWallet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/CreateWallet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).CreateWallet(ctx, req.(*CreateWalletInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBalanceInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/GetBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetBalance(ctx, req.(*GetBalanceInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/GetReceipt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetReceipt(ctx, req.(*GetReceiptInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_SendRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRawTransactionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).SendRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/SendRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).SendRawTransaction(ctx, req.(*SendRawTransactionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_CreateRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRawTransactionInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).CreateRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/CreateRawTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).CreateRawTransaction(ctx, req.(*CreateRawTransactionInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockNumberInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/GetBlockNumber",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetBlockNumber(ctx, req.(*GetBlockNumberInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolygonProxy_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolygonProxyServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/polygon.PolygonProxy/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolygonProxyServer).GetBlocks(ctx, req.(*GetBlocksInput))
	}
	return interceptor(ctx, in, info, handler)
}

// PolygonProxy_ServiceDesc is the grpc.ServiceDesc for PolygonProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolygonProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "polygon.PolygonProxy",
	HandlerType: (*PolygonProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateWallet",
			Handler:    _PolygonProxy_CreateWallet_Handler,
		},
		{
			MethodName: "GetBalance",
			Handler:    _PolygonProxy_GetBalance_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _PolygonProxy_GetReceipt_Handler,
		},
		{
			MethodName: "SendRawTransaction",
			Handler:    _PolygonProxy_SendRawTransaction_Handler,
		},
		{
			MethodName: "CreateRawTransaction",
			Handler:    _PolygonProxy_CreateRawTransaction_Handler,
		},
		{
			MethodName: "GetBlockNumber",
			Handler:    _PolygonProxy_GetBlockNumber_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _PolygonProxy_GetBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/polygon/polygon.proxy.proto",
}
