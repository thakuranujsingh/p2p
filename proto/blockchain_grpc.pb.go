// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto/blockchain.proto

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

// BlockchainClient is the client API for Blockchain service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockchainClient interface {
	AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error)
	GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error)
	BroadCast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
}

type blockchainClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockchainClient(cc grpc.ClientConnInterface) BlockchainClient {
	return &blockchainClient{cc}
}

func (c *blockchainClient) AddBlock(ctx context.Context, in *AddBlockRequest, opts ...grpc.CallOption) (*AddBlockResponse, error) {
	out := new(AddBlockResponse)
	err := c.cc.Invoke(ctx, "/proto.blockchain/AddBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) GetBlockchain(ctx context.Context, in *GetBlockchainRequest, opts ...grpc.CallOption) (*GetBlockchainResponse, error) {
	out := new(GetBlockchainResponse)
	err := c.cc.Invoke(ctx, "/proto.blockchain/GetBlockchain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockchainClient) BroadCast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, "/proto.blockchain/BroadCast", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockchainServer is the server API for Blockchain service.
// All implementations must embed UnimplementedBlockchainServer
// for forward compatibility
type BlockchainServer interface {
	AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error)
	GetBlockchain(context.Context, *GetBlockchainRequest) (*GetBlockchainResponse, error)
	BroadCast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
	mustEmbedUnimplementedBlockchainServer()
}

// UnimplementedBlockchainServer must be embedded to have forward compatible implementations.
type UnimplementedBlockchainServer struct {
}

func (UnimplementedBlockchainServer) AddBlock(context.Context, *AddBlockRequest) (*AddBlockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlock not implemented")
}
func (UnimplementedBlockchainServer) GetBlockchain(context.Context, *GetBlockchainRequest) (*GetBlockchainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockchain not implemented")
}
func (UnimplementedBlockchainServer) BroadCast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadCast not implemented")
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

func _Blockchain_AddBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddBlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).AddBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.blockchain/AddBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).AddBlock(ctx, req.(*AddBlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_GetBlockchain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlockchainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).GetBlockchain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.blockchain/GetBlockchain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).GetBlockchain(ctx, req.(*GetBlockchainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Blockchain_BroadCast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockchainServer).BroadCast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.blockchain/BroadCast",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockchainServer).BroadCast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Blockchain_ServiceDesc is the grpc.ServiceDesc for Blockchain service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Blockchain_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.blockchain",
	HandlerType: (*BlockchainServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddBlock",
			Handler:    _Blockchain_AddBlock_Handler,
		},
		{
			MethodName: "GetBlockchain",
			Handler:    _Blockchain_GetBlockchain_Handler,
		},
		{
			MethodName: "BroadCast",
			Handler:    _Blockchain_BroadCast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/blockchain.proto",
}
