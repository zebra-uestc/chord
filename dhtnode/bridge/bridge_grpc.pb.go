// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package bridge

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

// BridgeToOrderClient is the client API for BridgeToOrder service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeToOrderClient interface {
	// 由发送方调用函数，接收方实现函数
	TransBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Status, error)
}

type bridgeToOrderClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeToOrderClient(cc grpc.ClientConnInterface) BridgeToOrderClient {
	return &bridgeToOrderClient{cc}
}

func (c *bridgeToOrderClient) TransBlock(ctx context.Context, in *Block, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BridgeToOrder/TransBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeToOrderServer is the server API for BridgeToOrder service.
// All implementations must embed UnimplementedBridgeToOrderServer
// for forward compatibility
type BridgeToOrderServer interface {
	// 由发送方调用函数，接收方实现函数
	TransBlock(context.Context, *Block) (*Status, error)
	mustEmbedUnimplementedBridgeToOrderServer()
}

// UnimplementedBridgeToOrderServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeToOrderServer struct {
}

func (UnimplementedBridgeToOrderServer) TransBlock(context.Context, *Block) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransBlock not implemented")
}
func (UnimplementedBridgeToOrderServer) mustEmbedUnimplementedBridgeToOrderServer() {}

// UnsafeBridgeToOrderServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeToOrderServer will
// result in compilation errors.
type UnsafeBridgeToOrderServer interface {
	mustEmbedUnimplementedBridgeToOrderServer()
}

func RegisterBridgeToOrderServer(s grpc.ServiceRegistrar, srv BridgeToOrderServer) {
	s.RegisterService(&BridgeToOrder_ServiceDesc, srv)
}

func _BridgeToOrder_TransBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Block)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeToOrderServer).TransBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BridgeToOrder/TransBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeToOrderServer).TransBlock(ctx, req.(*Block))
	}
	return interceptor(ctx, in, info, handler)
}

// BridgeToOrder_ServiceDesc is the grpc.ServiceDesc for BridgeToOrder service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BridgeToOrder_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BridgeToOrder",
	HandlerType: (*BridgeToOrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransBlock",
			Handler:    _BridgeToOrder_TransBlock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dhtnode/bridge/bridge.proto",
}

// BridgeToDhtClient is the client API for BridgeToDht service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BridgeToDhtClient interface {
	TransMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Status, error)
	// dht调用，orderer实现
	LoadConfig(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Config, error)
}

type bridgeToDhtClient struct {
	cc grpc.ClientConnInterface
}

func NewBridgeToDhtClient(cc grpc.ClientConnInterface) BridgeToDhtClient {
	return &bridgeToDhtClient{cc}
}

func (c *bridgeToDhtClient) TransMsg(ctx context.Context, in *Msg, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/BridgeToDht/TransMsg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bridgeToDhtClient) LoadConfig(ctx context.Context, in *Status, opts ...grpc.CallOption) (*Config, error) {
	out := new(Config)
	err := c.cc.Invoke(ctx, "/BridgeToDht/LoadConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BridgeToDhtServer is the server API for BridgeToDht service.
// All implementations must embed UnimplementedBridgeToDhtServer
// for forward compatibility
type BridgeToDhtServer interface {
	TransMsg(context.Context, *Msg) (*Status, error)
	// dht调用，orderer实现
	LoadConfig(context.Context, *Status) (*Config, error)
	mustEmbedUnimplementedBridgeToDhtServer()
}

// UnimplementedBridgeToDhtServer must be embedded to have forward compatible implementations.
type UnimplementedBridgeToDhtServer struct {
}

func (UnimplementedBridgeToDhtServer) TransMsg(context.Context, *Msg) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransMsg not implemented")
}
func (UnimplementedBridgeToDhtServer) LoadConfig(context.Context, *Status) (*Config, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadConfig not implemented")
}
func (UnimplementedBridgeToDhtServer) mustEmbedUnimplementedBridgeToDhtServer() {}

// UnsafeBridgeToDhtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BridgeToDhtServer will
// result in compilation errors.
type UnsafeBridgeToDhtServer interface {
	mustEmbedUnimplementedBridgeToDhtServer()
}

func RegisterBridgeToDhtServer(s grpc.ServiceRegistrar, srv BridgeToDhtServer) {
	s.RegisterService(&BridgeToDht_ServiceDesc, srv)
}

func _BridgeToDht_TransMsg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Msg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeToDhtServer).TransMsg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BridgeToDht/TransMsg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeToDhtServer).TransMsg(ctx, req.(*Msg))
	}
	return interceptor(ctx, in, info, handler)
}

func _BridgeToDht_LoadConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Status)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BridgeToDhtServer).LoadConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BridgeToDht/LoadConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BridgeToDhtServer).LoadConfig(ctx, req.(*Status))
	}
	return interceptor(ctx, in, info, handler)
}

// BridgeToDht_ServiceDesc is the grpc.ServiceDesc for BridgeToDht service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BridgeToDht_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BridgeToDht",
	HandlerType: (*BridgeToDhtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransMsg",
			Handler:    _BridgeToDht_TransMsg_Handler,
		},
		{
			MethodName: "LoadConfig",
			Handler:    _BridgeToDht_LoadConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dhtnode/bridge/bridge.proto",
}
