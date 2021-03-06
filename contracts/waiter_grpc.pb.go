// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: contracts/waiter.proto

package WaiterProto

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

// WaiterServiceClient is the client API for WaiterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WaiterServiceClient interface {
	Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderAcknowledged, error)
}

type waiterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWaiterServiceClient(cc grpc.ClientConnInterface) WaiterServiceClient {
	return &waiterServiceClient{cc}
}

func (c *waiterServiceClient) Order(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderAcknowledged, error) {
	out := new(OrderAcknowledged)
	err := c.cc.Invoke(ctx, "/WaiterProto.WaiterService/Order", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WaiterServiceServer is the server API for WaiterService service.
// All implementations must embed UnimplementedWaiterServiceServer
// for forward compatibility
type WaiterServiceServer interface {
	Order(context.Context, *OrderRequest) (*OrderAcknowledged, error)
	mustEmbedUnimplementedWaiterServiceServer()
}

// UnimplementedWaiterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWaiterServiceServer struct {
}

func (UnimplementedWaiterServiceServer) Order(context.Context, *OrderRequest) (*OrderAcknowledged, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Order not implemented")
}
func (UnimplementedWaiterServiceServer) mustEmbedUnimplementedWaiterServiceServer() {}

// UnsafeWaiterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WaiterServiceServer will
// result in compilation errors.
type UnsafeWaiterServiceServer interface {
	mustEmbedUnimplementedWaiterServiceServer()
}

func RegisterWaiterServiceServer(s grpc.ServiceRegistrar, srv WaiterServiceServer) {
	s.RegisterService(&WaiterService_ServiceDesc, srv)
}

func _WaiterService_Order_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WaiterServiceServer).Order(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WaiterProto.WaiterService/Order",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WaiterServiceServer).Order(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// WaiterService_ServiceDesc is the grpc.ServiceDesc for WaiterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WaiterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "WaiterProto.WaiterService",
	HandlerType: (*WaiterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Order",
			Handler:    _WaiterService_Order_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contracts/waiter.proto",
}
