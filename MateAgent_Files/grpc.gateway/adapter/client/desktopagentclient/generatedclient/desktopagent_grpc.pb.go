// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dapb

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

// DesktopAgentServiceClient is the client API for DesktopAgentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DesktopAgentServiceClient interface {
	DesktopAgentSession(ctx context.Context, opts ...grpc.CallOption) (DesktopAgentService_DesktopAgentSessionClient, error)
}

type desktopAgentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDesktopAgentServiceClient(cc grpc.ClientConnInterface) DesktopAgentServiceClient {
	return &desktopAgentServiceClient{cc}
}

func (c *desktopAgentServiceClient) DesktopAgentSession(ctx context.Context, opts ...grpc.CallOption) (DesktopAgentService_DesktopAgentSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &DesktopAgentService_ServiceDesc.Streams[0], "/desktopagentgeneratedclient.DesktopAgentService/DesktopAgentSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &desktopAgentServiceDesktopAgentSessionClient{stream}
	return x, nil
}

type DesktopAgentService_DesktopAgentSessionClient interface {
	Send(*DesktopAgentSessionReq) error
	Recv() (*DesktopAgentSessionResp, error)
	grpc.ClientStream
}

type desktopAgentServiceDesktopAgentSessionClient struct {
	grpc.ClientStream
}

func (x *desktopAgentServiceDesktopAgentSessionClient) Send(m *DesktopAgentSessionReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *desktopAgentServiceDesktopAgentSessionClient) Recv() (*DesktopAgentSessionResp, error) {
	m := new(DesktopAgentSessionResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DesktopAgentServiceServer is the server API for DesktopAgentService service.
// All implementations must embed UnimplementedDesktopAgentServiceServer
// for forward compatibility
type DesktopAgentServiceServer interface {
	DesktopAgentSession(DesktopAgentService_DesktopAgentSessionServer) error
	mustEmbedUnimplementedDesktopAgentServiceServer()
}

// UnimplementedDesktopAgentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDesktopAgentServiceServer struct {
}

func (UnimplementedDesktopAgentServiceServer) DesktopAgentSession(DesktopAgentService_DesktopAgentSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method DesktopAgentSession not implemented")
}
func (UnimplementedDesktopAgentServiceServer) mustEmbedUnimplementedDesktopAgentServiceServer() {}

// UnsafeDesktopAgentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DesktopAgentServiceServer will
// result in compilation errors.
type UnsafeDesktopAgentServiceServer interface {
	mustEmbedUnimplementedDesktopAgentServiceServer()
}

func RegisterDesktopAgentServiceServer(s grpc.ServiceRegistrar, srv DesktopAgentServiceServer) {
	s.RegisterService(&DesktopAgentService_ServiceDesc, srv)
}

func _DesktopAgentService_DesktopAgentSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DesktopAgentServiceServer).DesktopAgentSession(&desktopAgentServiceDesktopAgentSessionServer{stream})
}

type DesktopAgentService_DesktopAgentSessionServer interface {
	Send(*DesktopAgentSessionResp) error
	Recv() (*DesktopAgentSessionReq, error)
	grpc.ServerStream
}

type desktopAgentServiceDesktopAgentSessionServer struct {
	grpc.ServerStream
}

func (x *desktopAgentServiceDesktopAgentSessionServer) Send(m *DesktopAgentSessionResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *desktopAgentServiceDesktopAgentSessionServer) Recv() (*DesktopAgentSessionReq, error) {
	m := new(DesktopAgentSessionReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DesktopAgentService_ServiceDesc is the grpc.ServiceDesc for DesktopAgentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DesktopAgentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "desktopagentgeneratedclient.DesktopAgentService",
	HandlerType: (*DesktopAgentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DesktopAgentSession",
			Handler:       _DesktopAgentService_DesktopAgentSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "desktopagent.proto",
}
