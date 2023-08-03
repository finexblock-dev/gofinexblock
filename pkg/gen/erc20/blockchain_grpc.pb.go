// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.3
// source: pkg/proto/erc20/blockchain.proto

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
	Blockchain_GetCurrentBlockNumber_FullMethodName = "/erc20.Blockchain/GetCurrentBlockNumber"
	Blockchain_GetBlocks_FullMethodName             = "/erc20.Blockchain/GetBlocks"
	Blockchain_GetReceipt_FullMethodName            = "/erc20.Blockchain/GetReceipt"
)

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainClient interface {
	GetCurrentBlockNumber(ctx context.Context, in *GetCurrentBlockNumberInput, opts ...grpc.CallOption) (*GetCurrentBlockNumberOutput, error)
	GetBlocks(ctx context.Context, in *GetBlocksInput, opts ...grpc.CallOption) (*GetBlocksOutput, error)
	GetReceipt(ctx context.Context, in *GetReceiptInput, opts ...grpc.CallOption) (*GetReceiptOutput, error)
}

type blockchainClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainClient(cc grpc.ClientConnInterface) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) GetCurrentBlockNumber(ctx context.Context, in *GetCurrentBlockNumberInput, opts ...grpc.CallOption) (*GetCurrentBlockNumberOutput, error) {
	out := new(GetCurrentBlockNumberOutput)
	err := c.cc.Invoke(ctx, Blockchain_GetCurrentBlockNumber_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlocks(ctx context.Context, in *GetBlocksInput, opts ...grpc.CallOption) (*GetBlocksOutput, error) {
	out := new(GetBlocksOutput)
	err := c.cc.Invoke(ctx, Blockchain_GetBlocks_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetReceipt(ctx context.Context, in *GetReceiptInput, opts ...grpc.CallOption) (*GetReceiptOutput, error) {
	out := new(GetReceiptOutput)
	err := c.cc.Invoke(ctx, Blockchain_GetReceipt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
// All implementations must embed UnimplementedBlockchainServer
// for forward compatibility
type BlockchainServer interface {
	GetCurrentBlockNumber(context.Context, *GetCurrentBlockNumberInput) (*GetCurrentBlockNumberOutput, error)
	GetBlocks(context.Context, *GetBlocksInput) (*GetBlocksOutput, error)
	GetReceipt(context.Context, *GetReceiptInput) (*GetReceiptOutput, error)
	mustEmbedUnimplementedBlockchainServer()
}

// UnimplementedBlockchainServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainServer struct {
}

func (UnimplementedBlockchainServer) GetCurrentBlockNumber(context.Context, *GetCurrentBlockNumberInput) (*GetCurrentBlockNumberOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentBlockNumber not implemented")
}
func (UnimplementedBlockchainServer) GetBlocks(context.Context, *GetBlocksInput) (*GetBlocksOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlocks not implemented")
}
func (UnimplementedBlockchainServer) GetReceipt(context.Context, *GetReceiptInput) (*GetReceiptOutput, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetReceipt not implemented")
}
func (UnimplementedBlockchainServer) mustEmbedUnimplementedBlockchainServer() {}

// UnsafeBlockchainServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockchainServer will
// result in compilation errors.
type UnsafeBlockchainServer interface {
	mustEmbedUnimplementedBlockchainServer()
}

func RegisterBlockchainServer(s grpc.ServiceRegistrar, srv BlockchainServer) {
	s.RegisterService(&Blockchain_ServiceDesc, srv)
}

func _Blockchain_GetCurrentBlockNumber_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCurrentBlockNumberInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetCurrentBlockNumber(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_GetCurrentBlockNumber_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetCurrentBlockNumber(ctx, req.(*GetCurrentBlockNumberInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlocksInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_GetBlocks_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlocks(ctx, req.(*GetBlocksInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetReceipt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReceiptInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetReceipt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Blockchain_GetReceipt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetReceipt(ctx, req.(*GetReceiptInput))
	}
	return interceptor(ctx, in, info, handler)
}

// Blockchain_ServiceDesc is the grpc.ServiceDesc for Blockchain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blockchain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "erc20.Blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentBlockNumber",
			Handler:    _Blockchain_GetCurrentBlockNumber_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _Blockchain_GetBlocks_Handler,
		},
		{
			MethodName: "GetReceipt",
			Handler:    _Blockchain_GetReceipt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/proto/erc20/blockchain.proto",
}
