// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: anomaly.proto

package pb

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

// AnomalyClient is the client API for Anomaly service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AnomalyClient interface {
	Subscribe(ctx context.Context, opts ...grpc.CallOption) (Anomaly_SubscribeClient, error)
	GetPriceAllWindow(ctx context.Context, in *AmmId, opts ...grpc.CallOption) (*PriceAllWindow, error)
	GetVolumeAllWindow(ctx context.Context, in *AmmId, opts ...grpc.CallOption) (*VolumeAllWindow, error)
	GetOHLCPriceAllWindow(ctx context.Context, in *GetOHLCPriceAllWindowArgs, opts ...grpc.CallOption) (*OHLCPriceAllWindow, error)
}

type anomalyClient struct {
	cc grpc.ClientConnInterface
}

func NewAnomalyClient(cc grpc.ClientConnInterface) AnomalyClient {
	return &anomalyClient{cc}
}

func (c *anomalyClient) Subscribe(ctx context.Context, opts ...grpc.CallOption) (Anomaly_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &Anomaly_ServiceDesc.Streams[0], "/solom.Anomaly/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &anomalySubscribeClient{stream}
	return x, nil
}

type Anomaly_SubscribeClient interface {
	Send(*SubscribeRequest) error
	Recv() (*SubscribeUpdate, error)
	grpc.ClientStream
}

type anomalySubscribeClient struct {
	grpc.ClientStream
}

func (x *anomalySubscribeClient) Send(m *SubscribeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *anomalySubscribeClient) Recv() (*SubscribeUpdate, error) {
	m := new(SubscribeUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *anomalyClient) GetPriceAllWindow(ctx context.Context, in *AmmId, opts ...grpc.CallOption) (*PriceAllWindow, error) {
	out := new(PriceAllWindow)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetPriceAllWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetVolumeAllWindow(ctx context.Context, in *AmmId, opts ...grpc.CallOption) (*VolumeAllWindow, error) {
	out := new(VolumeAllWindow)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetVolumeAllWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetOHLCPriceAllWindow(ctx context.Context, in *GetOHLCPriceAllWindowArgs, opts ...grpc.CallOption) (*OHLCPriceAllWindow, error) {
	out := new(OHLCPriceAllWindow)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetOHLCPriceAllWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnomalyServer is the server API for Anomaly service.
// All implementations must embed UnimplementedAnomalyServer
// for forward compatibility
type AnomalyServer interface {
	Subscribe(Anomaly_SubscribeServer) error
	GetPriceAllWindow(context.Context, *AmmId) (*PriceAllWindow, error)
	GetVolumeAllWindow(context.Context, *AmmId) (*VolumeAllWindow, error)
	GetOHLCPriceAllWindow(context.Context, *GetOHLCPriceAllWindowArgs) (*OHLCPriceAllWindow, error)
	mustEmbedUnimplementedAnomalyServer()
}

// UnimplementedAnomalyServer must be embedded to have forward compatible implementations.
type UnimplementedAnomalyServer struct {
}

func (UnimplementedAnomalyServer) Subscribe(Anomaly_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedAnomalyServer) GetPriceAllWindow(context.Context, *AmmId) (*PriceAllWindow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPriceAllWindow not implemented")
}
func (UnimplementedAnomalyServer) GetVolumeAllWindow(context.Context, *AmmId) (*VolumeAllWindow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVolumeAllWindow not implemented")
}
func (UnimplementedAnomalyServer) GetOHLCPriceAllWindow(context.Context, *GetOHLCPriceAllWindowArgs) (*OHLCPriceAllWindow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOHLCPriceAllWindow not implemented")
}
func (UnimplementedAnomalyServer) mustEmbedUnimplementedAnomalyServer() {}

// UnsafeAnomalyServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AnomalyServer will
// result in compilation errors.
type UnsafeAnomalyServer interface {
	mustEmbedUnimplementedAnomalyServer()
}

func RegisterAnomalyServer(s grpc.ServiceRegistrar, srv AnomalyServer) {
	s.RegisterService(&Anomaly_ServiceDesc, srv)
}

func _Anomaly_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnomalyServer).Subscribe(&anomalySubscribeServer{stream})
}

type Anomaly_SubscribeServer interface {
	Send(*SubscribeUpdate) error
	Recv() (*SubscribeRequest, error)
	grpc.ServerStream
}

type anomalySubscribeServer struct {
	grpc.ServerStream
}

func (x *anomalySubscribeServer) Send(m *SubscribeUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *anomalySubscribeServer) Recv() (*SubscribeRequest, error) {
	m := new(SubscribeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Anomaly_GetPriceAllWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmmId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetPriceAllWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetPriceAllWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetPriceAllWindow(ctx, req.(*AmmId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetVolumeAllWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AmmId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetVolumeAllWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetVolumeAllWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetVolumeAllWindow(ctx, req.(*AmmId))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetOHLCPriceAllWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOHLCPriceAllWindowArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetOHLCPriceAllWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetOHLCPriceAllWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetOHLCPriceAllWindow(ctx, req.(*GetOHLCPriceAllWindowArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// Anomaly_ServiceDesc is the grpc.ServiceDesc for Anomaly service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Anomaly_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "solom.Anomaly",
	HandlerType: (*AnomalyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPriceAllWindow",
			Handler:    _Anomaly_GetPriceAllWindow_Handler,
		},
		{
			MethodName: "GetVolumeAllWindow",
			Handler:    _Anomaly_GetVolumeAllWindow_Handler,
		},
		{
			MethodName: "GetOHLCPriceAllWindow",
			Handler:    _Anomaly_GetOHLCPriceAllWindow_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Anomaly_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "anomaly.proto",
}
