// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package flowpb

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

// FlowServiceClient is the client API for FlowService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FlowServiceClient interface {
	DebugFlow(ctx context.Context, opts ...grpc.CallOption) (FlowService_DebugFlowClient, error)
	RunFlow(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunFlowClient, error)
	RunAction(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunActionClient, error)
	RunProcess(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunProcessClient, error)
	StopCronJob(ctx context.Context, in *StopCronJobRequest, opts ...grpc.CallOption) (*StopCronJobResponse, error)
	RunDirect(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunDirectClient, error)
}

type flowServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFlowServiceClient(cc grpc.ClientConnInterface) FlowServiceClient {
	return &flowServiceClient{cc}
}

func (c *flowServiceClient) DebugFlow(ctx context.Context, opts ...grpc.CallOption) (FlowService_DebugFlowClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlowService_ServiceDesc.Streams[0], "/generatedclient.FlowService/DebugFlow", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceDebugFlowClient{stream}
	return x, nil
}

type FlowService_DebugFlowClient interface {
	Send(*RunFlowRequest) error
	Recv() (*RunFlowResponse, error)
	grpc.ClientStream
}

type flowServiceDebugFlowClient struct {
	grpc.ClientStream
}

func (x *flowServiceDebugFlowClient) Send(m *RunFlowRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceDebugFlowClient) Recv() (*RunFlowResponse, error) {
	m := new(RunFlowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowServiceClient) RunFlow(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunFlowClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlowService_ServiceDesc.Streams[1], "/generatedclient.FlowService/RunFlow", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceRunFlowClient{stream}
	return x, nil
}

type FlowService_RunFlowClient interface {
	Send(*RunFlowRequest) error
	Recv() (*RunFlowResponse, error)
	grpc.ClientStream
}

type flowServiceRunFlowClient struct {
	grpc.ClientStream
}

func (x *flowServiceRunFlowClient) Send(m *RunFlowRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceRunFlowClient) Recv() (*RunFlowResponse, error) {
	m := new(RunFlowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowServiceClient) RunAction(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunActionClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlowService_ServiceDesc.Streams[2], "/generatedclient.FlowService/RunAction", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceRunActionClient{stream}
	return x, nil
}

type FlowService_RunActionClient interface {
	Send(*RunActionRequest) error
	Recv() (*RunActionResponse, error)
	grpc.ClientStream
}

type flowServiceRunActionClient struct {
	grpc.ClientStream
}

func (x *flowServiceRunActionClient) Send(m *RunActionRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceRunActionClient) Recv() (*RunActionResponse, error) {
	m := new(RunActionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowServiceClient) RunProcess(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunProcessClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlowService_ServiceDesc.Streams[3], "/generatedclient.FlowService/RunProcess", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceRunProcessClient{stream}
	return x, nil
}

type FlowService_RunProcessClient interface {
	Send(*RunProcessRequest) error
	Recv() (*RunFlowResponse, error)
	grpc.ClientStream
}

type flowServiceRunProcessClient struct {
	grpc.ClientStream
}

func (x *flowServiceRunProcessClient) Send(m *RunProcessRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceRunProcessClient) Recv() (*RunFlowResponse, error) {
	m := new(RunFlowResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *flowServiceClient) StopCronJob(ctx context.Context, in *StopCronJobRequest, opts ...grpc.CallOption) (*StopCronJobResponse, error) {
	out := new(StopCronJobResponse)
	err := c.cc.Invoke(ctx, "/generatedclient.FlowService/StopCronJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *flowServiceClient) RunDirect(ctx context.Context, opts ...grpc.CallOption) (FlowService_RunDirectClient, error) {
	stream, err := c.cc.NewStream(ctx, &FlowService_ServiceDesc.Streams[4], "/generatedclient.FlowService/RunDirect", opts...)
	if err != nil {
		return nil, err
	}
	x := &flowServiceRunDirectClient{stream}
	return x, nil
}

type FlowService_RunDirectClient interface {
	Send(*RunDirectRequest) error
	Recv() (*RunDirectResponse, error)
	grpc.ClientStream
}

type flowServiceRunDirectClient struct {
	grpc.ClientStream
}

func (x *flowServiceRunDirectClient) Send(m *RunDirectRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *flowServiceRunDirectClient) Recv() (*RunDirectResponse, error) {
	m := new(RunDirectResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowServiceServer is the server API for FlowService service.
// All implementations must embed UnimplementedFlowServiceServer
// for forward compatibility
type FlowServiceServer interface {
	DebugFlow(FlowService_DebugFlowServer) error
	RunFlow(FlowService_RunFlowServer) error
	RunAction(FlowService_RunActionServer) error
	RunProcess(FlowService_RunProcessServer) error
	StopCronJob(context.Context, *StopCronJobRequest) (*StopCronJobResponse, error)
	RunDirect(FlowService_RunDirectServer) error
	mustEmbedUnimplementedFlowServiceServer()
}

// UnimplementedFlowServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFlowServiceServer struct {
}

func (UnimplementedFlowServiceServer) DebugFlow(FlowService_DebugFlowServer) error {
	return status.Errorf(codes.Unimplemented, "method DebugFlow not implemented")
}
func (UnimplementedFlowServiceServer) RunFlow(FlowService_RunFlowServer) error {
	return status.Errorf(codes.Unimplemented, "method RunFlow not implemented")
}
func (UnimplementedFlowServiceServer) RunAction(FlowService_RunActionServer) error {
	return status.Errorf(codes.Unimplemented, "method RunAction not implemented")
}
func (UnimplementedFlowServiceServer) RunProcess(FlowService_RunProcessServer) error {
	return status.Errorf(codes.Unimplemented, "method RunProcess not implemented")
}
func (UnimplementedFlowServiceServer) StopCronJob(context.Context, *StopCronJobRequest) (*StopCronJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopCronJob not implemented")
}
func (UnimplementedFlowServiceServer) RunDirect(FlowService_RunDirectServer) error {
	return status.Errorf(codes.Unimplemented, "method RunDirect not implemented")
}
func (UnimplementedFlowServiceServer) mustEmbedUnimplementedFlowServiceServer() {}

// UnsafeFlowServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FlowServiceServer will
// result in compilation errors.
type UnsafeFlowServiceServer interface {
	mustEmbedUnimplementedFlowServiceServer()
}

func RegisterFlowServiceServer(s grpc.ServiceRegistrar, srv FlowServiceServer) {
	s.RegisterService(&FlowService_ServiceDesc, srv)
}

func _FlowService_DebugFlow_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).DebugFlow(&flowServiceDebugFlowServer{stream})
}

type FlowService_DebugFlowServer interface {
	Send(*RunFlowResponse) error
	Recv() (*RunFlowRequest, error)
	grpc.ServerStream
}

type flowServiceDebugFlowServer struct {
	grpc.ServerStream
}

func (x *flowServiceDebugFlowServer) Send(m *RunFlowResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceDebugFlowServer) Recv() (*RunFlowRequest, error) {
	m := new(RunFlowRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowService_RunFlow_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).RunFlow(&flowServiceRunFlowServer{stream})
}

type FlowService_RunFlowServer interface {
	Send(*RunFlowResponse) error
	Recv() (*RunFlowRequest, error)
	grpc.ServerStream
}

type flowServiceRunFlowServer struct {
	grpc.ServerStream
}

func (x *flowServiceRunFlowServer) Send(m *RunFlowResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceRunFlowServer) Recv() (*RunFlowRequest, error) {
	m := new(RunFlowRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowService_RunAction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).RunAction(&flowServiceRunActionServer{stream})
}

type FlowService_RunActionServer interface {
	Send(*RunActionResponse) error
	Recv() (*RunActionRequest, error)
	grpc.ServerStream
}

type flowServiceRunActionServer struct {
	grpc.ServerStream
}

func (x *flowServiceRunActionServer) Send(m *RunActionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceRunActionServer) Recv() (*RunActionRequest, error) {
	m := new(RunActionRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowService_RunProcess_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).RunProcess(&flowServiceRunProcessServer{stream})
}

type FlowService_RunProcessServer interface {
	Send(*RunFlowResponse) error
	Recv() (*RunProcessRequest, error)
	grpc.ServerStream
}

type flowServiceRunProcessServer struct {
	grpc.ServerStream
}

func (x *flowServiceRunProcessServer) Send(m *RunFlowResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceRunProcessServer) Recv() (*RunProcessRequest, error) {
	m := new(RunProcessRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _FlowService_StopCronJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopCronJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FlowServiceServer).StopCronJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/generatedclient.FlowService/StopCronJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FlowServiceServer).StopCronJob(ctx, req.(*StopCronJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FlowService_RunDirect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(FlowServiceServer).RunDirect(&flowServiceRunDirectServer{stream})
}

type FlowService_RunDirectServer interface {
	Send(*RunDirectResponse) error
	Recv() (*RunDirectRequest, error)
	grpc.ServerStream
}

type flowServiceRunDirectServer struct {
	grpc.ServerStream
}

func (x *flowServiceRunDirectServer) Send(m *RunDirectResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *flowServiceRunDirectServer) Recv() (*RunDirectRequest, error) {
	m := new(RunDirectRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// FlowService_ServiceDesc is the grpc.ServiceDesc for FlowService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FlowService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "generatedclient.FlowService",
	HandlerType: (*FlowServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StopCronJob",
			Handler:    _FlowService_StopCronJob_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DebugFlow",
			Handler:       _FlowService_DebugFlow_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunFlow",
			Handler:       _FlowService_RunFlow_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunAction",
			Handler:       _FlowService_RunAction_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunProcess",
			Handler:       _FlowService_RunProcess_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "RunDirect",
			Handler:       _FlowService_RunDirect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "flow.proto",
}
