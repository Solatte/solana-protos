// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
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
	SubscribeEvent(ctx context.Context, opts ...grpc.CallOption) (Anomaly_SubscribeEventClient, error)
	GetPriceAllWindow(ctx context.Context, in *Mint, opts ...grpc.CallOption) (*PriceAllWindow, error)
	GetOneMinuteVolumeByWindow(ctx context.Context, in *GetOneMinuteVolumeByWindowArgs, opts ...grpc.CallOption) (*OneMinuteVolumeByWindow, error)
	GetOHLCPriceAllWindow(ctx context.Context, in *GetOHLCPriceAllWindowArgs, opts ...grpc.CallOption) (*OHLCPriceAllWindow, error)
	IsAmmGood(ctx context.Context, in *IsAmmGoodArgs, opts ...grpc.CallOption) (*Boolean, error)
	GetMostActiveToken(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error)
	GetTokenByTrending(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error)
	GetTokenByBuy(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error)
	GetTokenBySell(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error)
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

func (c *anomalyClient) SubscribeEvent(ctx context.Context, opts ...grpc.CallOption) (Anomaly_SubscribeEventClient, error) {
	stream, err := c.cc.NewStream(ctx, &Anomaly_ServiceDesc.Streams[1], "/solom.Anomaly/SubscribeEvent", opts...)
	if err != nil {
		return nil, err
	}
	x := &anomalySubscribeEventClient{stream}
	return x, nil
}

type Anomaly_SubscribeEventClient interface {
	Send(*SubscribeEventRequest) error
	Recv() (*SubscribeEventUpdate, error)
	grpc.ClientStream
}

type anomalySubscribeEventClient struct {
	grpc.ClientStream
}

func (x *anomalySubscribeEventClient) Send(m *SubscribeEventRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *anomalySubscribeEventClient) Recv() (*SubscribeEventUpdate, error) {
	m := new(SubscribeEventUpdate)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *anomalyClient) GetPriceAllWindow(ctx context.Context, in *Mint, opts ...grpc.CallOption) (*PriceAllWindow, error) {
	out := new(PriceAllWindow)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetPriceAllWindow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetOneMinuteVolumeByWindow(ctx context.Context, in *GetOneMinuteVolumeByWindowArgs, opts ...grpc.CallOption) (*OneMinuteVolumeByWindow, error) {
	out := new(OneMinuteVolumeByWindow)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetOneMinuteVolumeByWindow", in, out, opts...)
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

func (c *anomalyClient) IsAmmGood(ctx context.Context, in *IsAmmGoodArgs, opts ...grpc.CallOption) (*Boolean, error) {
	out := new(Boolean)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/IsAmmGood", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetMostActiveToken(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error) {
	out := new(TokenBy)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetMostActiveToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetTokenByTrending(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error) {
	out := new(TokenBy)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetTokenByTrending", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetTokenByBuy(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error) {
	out := new(TokenBy)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetTokenByBuy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *anomalyClient) GetTokenBySell(ctx context.Context, in *GetTokenByArgs, opts ...grpc.CallOption) (*TokenBy, error) {
	out := new(TokenBy)
	err := c.cc.Invoke(ctx, "/solom.Anomaly/GetTokenBySell", in, out, opts...)
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
	SubscribeEvent(Anomaly_SubscribeEventServer) error
	GetPriceAllWindow(context.Context, *Mint) (*PriceAllWindow, error)
	GetOneMinuteVolumeByWindow(context.Context, *GetOneMinuteVolumeByWindowArgs) (*OneMinuteVolumeByWindow, error)
	GetOHLCPriceAllWindow(context.Context, *GetOHLCPriceAllWindowArgs) (*OHLCPriceAllWindow, error)
	IsAmmGood(context.Context, *IsAmmGoodArgs) (*Boolean, error)
	GetMostActiveToken(context.Context, *GetTokenByArgs) (*TokenBy, error)
	GetTokenByTrending(context.Context, *GetTokenByArgs) (*TokenBy, error)
	GetTokenByBuy(context.Context, *GetTokenByArgs) (*TokenBy, error)
	GetTokenBySell(context.Context, *GetTokenByArgs) (*TokenBy, error)
	mustEmbedUnimplementedAnomalyServer()
}

// UnimplementedAnomalyServer must be embedded to have forward compatible implementations.
type UnimplementedAnomalyServer struct {
}

func (UnimplementedAnomalyServer) Subscribe(Anomaly_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedAnomalyServer) SubscribeEvent(Anomaly_SubscribeEventServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeEvent not implemented")
}
func (UnimplementedAnomalyServer) GetPriceAllWindow(context.Context, *Mint) (*PriceAllWindow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPriceAllWindow not implemented")
}
func (UnimplementedAnomalyServer) GetOneMinuteVolumeByWindow(context.Context, *GetOneMinuteVolumeByWindowArgs) (*OneMinuteVolumeByWindow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneMinuteVolumeByWindow not implemented")
}
func (UnimplementedAnomalyServer) GetOHLCPriceAllWindow(context.Context, *GetOHLCPriceAllWindowArgs) (*OHLCPriceAllWindow, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOHLCPriceAllWindow not implemented")
}
func (UnimplementedAnomalyServer) IsAmmGood(context.Context, *IsAmmGoodArgs) (*Boolean, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsAmmGood not implemented")
}
func (UnimplementedAnomalyServer) GetMostActiveToken(context.Context, *GetTokenByArgs) (*TokenBy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMostActiveToken not implemented")
}
func (UnimplementedAnomalyServer) GetTokenByTrending(context.Context, *GetTokenByArgs) (*TokenBy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByTrending not implemented")
}
func (UnimplementedAnomalyServer) GetTokenByBuy(context.Context, *GetTokenByArgs) (*TokenBy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenByBuy not implemented")
}
func (UnimplementedAnomalyServer) GetTokenBySell(context.Context, *GetTokenByArgs) (*TokenBy, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenBySell not implemented")
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

func _Anomaly_SubscribeEvent_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AnomalyServer).SubscribeEvent(&anomalySubscribeEventServer{stream})
}

type Anomaly_SubscribeEventServer interface {
	Send(*SubscribeEventUpdate) error
	Recv() (*SubscribeEventRequest, error)
	grpc.ServerStream
}

type anomalySubscribeEventServer struct {
	grpc.ServerStream
}

func (x *anomalySubscribeEventServer) Send(m *SubscribeEventUpdate) error {
	return x.ServerStream.SendMsg(m)
}

func (x *anomalySubscribeEventServer) Recv() (*SubscribeEventRequest, error) {
	m := new(SubscribeEventRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Anomaly_GetPriceAllWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Mint)
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
		return srv.(AnomalyServer).GetPriceAllWindow(ctx, req.(*Mint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetOneMinuteVolumeByWindow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneMinuteVolumeByWindowArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetOneMinuteVolumeByWindow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetOneMinuteVolumeByWindow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetOneMinuteVolumeByWindow(ctx, req.(*GetOneMinuteVolumeByWindowArgs))
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

func _Anomaly_IsAmmGood_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAmmGoodArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).IsAmmGood(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/IsAmmGood",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).IsAmmGood(ctx, req.(*IsAmmGoodArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetMostActiveToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetMostActiveToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetMostActiveToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetMostActiveToken(ctx, req.(*GetTokenByArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetTokenByTrending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetTokenByTrending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetTokenByTrending",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetTokenByTrending(ctx, req.(*GetTokenByArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetTokenByBuy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetTokenByBuy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetTokenByBuy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetTokenByBuy(ctx, req.(*GetTokenByArgs))
	}
	return interceptor(ctx, in, info, handler)
}

func _Anomaly_GetTokenBySell_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenByArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnomalyServer).GetTokenBySell(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/solom.Anomaly/GetTokenBySell",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnomalyServer).GetTokenBySell(ctx, req.(*GetTokenByArgs))
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
			MethodName: "GetOneMinuteVolumeByWindow",
			Handler:    _Anomaly_GetOneMinuteVolumeByWindow_Handler,
		},
		{
			MethodName: "GetOHLCPriceAllWindow",
			Handler:    _Anomaly_GetOHLCPriceAllWindow_Handler,
		},
		{
			MethodName: "IsAmmGood",
			Handler:    _Anomaly_IsAmmGood_Handler,
		},
		{
			MethodName: "GetMostActiveToken",
			Handler:    _Anomaly_GetMostActiveToken_Handler,
		},
		{
			MethodName: "GetTokenByTrending",
			Handler:    _Anomaly_GetTokenByTrending_Handler,
		},
		{
			MethodName: "GetTokenByBuy",
			Handler:    _Anomaly_GetTokenByBuy_Handler,
		},
		{
			MethodName: "GetTokenBySell",
			Handler:    _Anomaly_GetTokenBySell_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _Anomaly_Subscribe_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "SubscribeEvent",
			Handler:       _Anomaly_SubscribeEvent_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "anomaly.proto",
}
