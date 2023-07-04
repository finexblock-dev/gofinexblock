// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: proto/proxy/proxy.bitcoin.proto

package bitcoin_proxy

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
	BitcoinProxy_GetBlockCount_FullMethodName        = "/bitcoin_proxy.BitcoinProxy/GetBlockCount"
	BitcoinProxy_GetBlock_FullMethodName             = "/bitcoin_proxy.BitcoinProxy/GetBlock"
	BitcoinProxy_GetRawTransaction_FullMethodName    = "/bitcoin_proxy.BitcoinProxy/GetRawTransaction"
	BitcoinProxy_CreateRawTransaction_FullMethodName = "/bitcoin_proxy.BitcoinProxy/CreateRawTransaction"
	BitcoinProxy_ListUnspent_FullMethodName          = "/bitcoin_proxy.BitcoinProxy/ListUnspent"
	BitcoinProxy_SignRawTransaction_FullMethodName   = "/bitcoin_proxy.BitcoinProxy/SignRawTransaction"
	BitcoinProxy_SendRawTransaction_FullMethodName   = "/bitcoin_proxy.BitcoinProxy/SendRawTransaction"
	BitcoinProxy_SendToAddress_FullMethodName        = "/bitcoin_proxy.BitcoinProxy/SendToAddress"
	BitcoinProxy_CreateWalletAddress_FullMethodName  = "/bitcoin_proxy.BitcoinProxy/CreateWalletAddress"
	BitcoinProxy_GetNewAddress_FullMethodName        = "/bitcoin_proxy.BitcoinProxy/GetNewAddress"
	BitcoinProxy_GetPrivateKey_FullMethodName        = "/bitcoin_proxy.BitcoinProxy/GetPrivateKey"
	BitcoinProxy_GetAddressUTXO_FullMethodName       = "/bitcoin_proxy.BitcoinProxy/GetAddressUTXO"
)

// BitcoinProxyClient is the client API for BitcoinProxy service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BitcoinProxyClient interface {
	GetBlockCount(ctx context.Context, in *GetBlockCountRequest, opts ...grpc.CallOption) (*GetBlockCountResponse, error)
	GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error)
	GetRawTransaction(ctx context.Context, in *GetRawTransactionRequest, opts ...grpc.CallOption) (*GetRawTransactionResponse, error)
	CreateRawTransaction(ctx context.Context, in *CreateRawTransactionRequest, opts ...grpc.CallOption) (*CreateRawTransactionResponse, error)
	ListUnspent(ctx context.Context, in *ListUnspentRequest, opts ...grpc.CallOption) (*ListUnspentResponse, error)
	SignRawTransaction(ctx context.Context, in *SignRawTransactionRequest, opts ...grpc.CallOption) (*SignRawTransactionResponse, error)
	SendRawTransaction(ctx context.Context, in *SendRawTransactionRequest, opts ...grpc.CallOption) (*SendRawTransactionResponse, error)
	SendToAddress(ctx context.Context, in *SendToAddressRequest, opts ...grpc.CallOption) (*SendToAddressResponse, error)
	CreateWalletAddress(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error)
	GetNewAddress(ctx context.Context, in *GetNewAddressRequest, opts ...grpc.CallOption) (*GetNewAddressResponse, error)
	GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error)
	GetAddressUTXO(ctx context.Context, in *GetAddressUTXORequest, opts ...grpc.CallOption) (*GetAddressUTXOResponse, error)
}

type bitcoinProxyClient struct {
	cc grpc.ClientConnInterface
}

func NewBitcoinProxyClient(cc grpc.ClientConnInterface) BitcoinProxyClient {
	return &bitcoinProxyClient{cc}
}

func (c *bitcoinProxyClient) GetBlockCount(ctx context.Context, in *GetBlockCountRequest, opts ...grpc.CallOption) (*GetBlockCountResponse, error) {
	out := new(GetBlockCountResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_GetBlockCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) GetBlock(ctx context.Context, in *GetBlockRequest, opts ...grpc.CallOption) (*GetBlockResponse, error) {
	out := new(GetBlockResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_GetBlock_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) GetRawTransaction(ctx context.Context, in *GetRawTransactionRequest, opts ...grpc.CallOption) (*GetRawTransactionResponse, error) {
	out := new(GetRawTransactionResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_GetRawTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) CreateRawTransaction(ctx context.Context, in *CreateRawTransactionRequest, opts ...grpc.CallOption) (*CreateRawTransactionResponse, error) {
	out := new(CreateRawTransactionResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_CreateRawTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) ListUnspent(ctx context.Context, in *ListUnspentRequest, opts ...grpc.CallOption) (*ListUnspentResponse, error) {
	out := new(ListUnspentResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_ListUnspent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) SignRawTransaction(ctx context.Context, in *SignRawTransactionRequest, opts ...grpc.CallOption) (*SignRawTransactionResponse, error) {
	out := new(SignRawTransactionResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_SignRawTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) SendRawTransaction(ctx context.Context, in *SendRawTransactionRequest, opts ...grpc.CallOption) (*SendRawTransactionResponse, error) {
	out := new(SendRawTransactionResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_SendRawTransaction_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) SendToAddress(ctx context.Context, in *SendToAddressRequest, opts ...grpc.CallOption) (*SendToAddressResponse, error) {
	out := new(SendToAddressResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_SendToAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) CreateWalletAddress(ctx context.Context, in *CreateWalletRequest, opts ...grpc.CallOption) (*CreateWalletResponse, error) {
	out := new(CreateWalletResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_CreateWalletAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) GetNewAddress(ctx context.Context, in *GetNewAddressRequest, opts ...grpc.CallOption) (*GetNewAddressResponse, error) {
	out := new(GetNewAddressResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_GetNewAddress_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) GetPrivateKey(ctx context.Context, in *GetPrivateKeyRequest, opts ...grpc.CallOption) (*GetPrivateKeyResponse, error) {
	out := new(GetPrivateKeyResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_GetPrivateKey_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bitcoinProxyClient) GetAddressUTXO(ctx context.Context, in *GetAddressUTXORequest, opts ...grpc.CallOption) (*GetAddressUTXOResponse, error) {
	out := new(GetAddressUTXOResponse)
	err := c.cc.Invoke(ctx, BitcoinProxy_GetAddressUTXO_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BitcoinProxyServer is the server API for BitcoinProxy service.
// All implementations must embed UnimplementedBitcoinProxyServer
// for forward compatibility
type BitcoinProxyServer interface {
	GetBlockCount(context.Context, *GetBlockCountRequest) (*GetBlockCountResponse, error)
	GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error)
	GetRawTransaction(context.Context, *GetRawTransactionRequest) (*GetRawTransactionResponse, error)
	CreateRawTransaction(context.Context, *CreateRawTransactionRequest) (*CreateRawTransactionResponse, error)
	ListUnspent(context.Context, *ListUnspentRequest) (*ListUnspentResponse, error)
	SignRawTransaction(context.Context, *SignRawTransactionRequest) (*SignRawTransactionResponse, error)
	SendRawTransaction(context.Context, *SendRawTransactionRequest) (*SendRawTransactionResponse, error)
	SendToAddress(context.Context, *SendToAddressRequest) (*SendToAddressResponse, error)
	CreateWalletAddress(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error)
	GetNewAddress(context.Context, *GetNewAddressRequest) (*GetNewAddressResponse, error)
	GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error)
	GetAddressUTXO(context.Context, *GetAddressUTXORequest) (*GetAddressUTXOResponse, error)
	mustEmbedUnimplementedBitcoinProxyServer()
}

// UnimplementedBitcoinProxyServer must be embedded to have forward compatible implementations.
type UnimplementedBitcoinProxyServer struct {
}

func (UnimplementedBitcoinProxyServer) GetBlockCount(context.Context, *GetBlockCountRequest) (*GetBlockCountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockCount not implemented")
}
func (UnimplementedBitcoinProxyServer) GetBlock(context.Context, *GetBlockRequest) (*GetBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlock not implemented")
}
func (UnimplementedBitcoinProxyServer) GetRawTransaction(context.Context, *GetRawTransactionRequest) (*GetRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRawTransaction not implemented")
}
func (UnimplementedBitcoinProxyServer) CreateRawTransaction(context.Context, *CreateRawTransactionRequest) (*CreateRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateRawTransaction not implemented")
}
func (UnimplementedBitcoinProxyServer) ListUnspent(context.Context, *ListUnspentRequest) (*ListUnspentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUnspent not implemented")
}
func (UnimplementedBitcoinProxyServer) SignRawTransaction(context.Context, *SignRawTransactionRequest) (*SignRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignRawTransaction not implemented")
}
func (UnimplementedBitcoinProxyServer) SendRawTransaction(context.Context, *SendRawTransactionRequest) (*SendRawTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendRawTransaction not implemented")
}
func (UnimplementedBitcoinProxyServer) SendToAddress(context.Context, *SendToAddressRequest) (*SendToAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendToAddress not implemented")
}
func (UnimplementedBitcoinProxyServer) CreateWalletAddress(context.Context, *CreateWalletRequest) (*CreateWalletResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateWalletAddress not implemented")
}
func (UnimplementedBitcoinProxyServer) GetNewAddress(context.Context, *GetNewAddressRequest) (*GetNewAddressResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNewAddress not implemented")
}
func (UnimplementedBitcoinProxyServer) GetPrivateKey(context.Context, *GetPrivateKeyRequest) (*GetPrivateKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPrivateKey not implemented")
}
func (UnimplementedBitcoinProxyServer) GetAddressUTXO(context.Context, *GetAddressUTXORequest) (*GetAddressUTXOResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAddressUTXO not implemented")
}
func (UnimplementedBitcoinProxyServer) mustEmbedUnimplementedBitcoinProxyServer() {}

// UnsafeBitcoinProxyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BitcoinProxyServer will
// result in compilation errors.
type UnsafeBitcoinProxyServer interface {
	mustEmbedUnimplementedBitcoinProxyServer()
}

func RegisterBitcoinProxyServer(s grpc.ServiceRegistrar, srv BitcoinProxyServer) {
	s.RegisterService(&BitcoinProxy_ServiceDesc, srv)
}

func _BitcoinProxy_GetBlockCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).GetBlockCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_GetBlockCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).GetBlockCount(ctx, req.(*GetBlockCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_GetBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).GetBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_GetBlock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).GetBlock(ctx, req.(*GetBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_GetRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).GetRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_GetRawTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).GetRawTransaction(ctx, req.(*GetRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_CreateRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).CreateRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_CreateRawTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).CreateRawTransaction(ctx, req.(*CreateRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_ListUnspent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUnspentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).ListUnspent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_ListUnspent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).ListUnspent(ctx, req.(*ListUnspentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_SignRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).SignRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_SignRawTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).SignRawTransaction(ctx, req.(*SignRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_SendRawTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendRawTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).SendRawTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_SendRawTransaction_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).SendRawTransaction(ctx, req.(*SendRawTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_SendToAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendToAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).SendToAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_SendToAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).SendToAddress(ctx, req.(*SendToAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_CreateWalletAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateWalletRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).CreateWalletAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_CreateWalletAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).CreateWalletAddress(ctx, req.(*CreateWalletRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_GetNewAddress_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNewAddressRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).GetNewAddress(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_GetNewAddress_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).GetNewAddress(ctx, req.(*GetNewAddressRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_GetPrivateKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPrivateKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).GetPrivateKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_GetPrivateKey_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).GetPrivateKey(ctx, req.(*GetPrivateKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BitcoinProxy_GetAddressUTXO_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAddressUTXORequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BitcoinProxyServer).GetAddressUTXO(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BitcoinProxy_GetAddressUTXO_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BitcoinProxyServer).GetAddressUTXO(ctx, req.(*GetAddressUTXORequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BitcoinProxy_ServiceDesc is the grpc.ServiceDesc for BitcoinProxy service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BitcoinProxy_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bitcoin_proxy.BitcoinProxy",
	HandlerType: (*BitcoinProxyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBlockCount",
			Handler:    _BitcoinProxy_GetBlockCount_Handler,
		},
		{
			MethodName: "GetBlock",
			Handler:    _BitcoinProxy_GetBlock_Handler,
		},
		{
			MethodName: "GetRawTransaction",
			Handler:    _BitcoinProxy_GetRawTransaction_Handler,
		},
		{
			MethodName: "CreateRawTransaction",
			Handler:    _BitcoinProxy_CreateRawTransaction_Handler,
		},
		{
			MethodName: "ListUnspent",
			Handler:    _BitcoinProxy_ListUnspent_Handler,
		},
		{
			MethodName: "SignRawTransaction",
			Handler:    _BitcoinProxy_SignRawTransaction_Handler,
		},
		{
			MethodName: "SendRawTransaction",
			Handler:    _BitcoinProxy_SendRawTransaction_Handler,
		},
		{
			MethodName: "SendToAddress",
			Handler:    _BitcoinProxy_SendToAddress_Handler,
		},
		{
			MethodName: "CreateWalletAddress",
			Handler:    _BitcoinProxy_CreateWalletAddress_Handler,
		},
		{
			MethodName: "GetNewAddress",
			Handler:    _BitcoinProxy_GetNewAddress_Handler,
		},
		{
			MethodName: "GetPrivateKey",
			Handler:    _BitcoinProxy_GetPrivateKey_Handler,
		},
		{
			MethodName: "GetAddressUTXO",
			Handler:    _BitcoinProxy_GetAddressUTXO_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proxy/proxy.bitcoin.proto",
}