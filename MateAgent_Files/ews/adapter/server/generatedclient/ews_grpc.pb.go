// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package ewspb

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

// EwsServiceClient is the client API for EwsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EwsServiceClient interface {
	EwsSession(ctx context.Context, opts ...grpc.CallOption) (EwsService_EwsSessionClient, error)
}

type ewsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEwsServiceClient(cc grpc.ClientConnInterface) EwsServiceClient {
	return &ewsServiceClient{cc}
}

func (c *ewsServiceClient) EwsSession(ctx context.Context, opts ...grpc.CallOption) (EwsService_EwsSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &EwsService_ServiceDesc.Streams[0], "/ewsgeneratedclient.EwsService/EwsSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &ewsServiceEwsSessionClient{stream}
	return x, nil
}

type EwsService_EwsSessionClient interface {
	Send(*EwsSessionReq) error
	Recv() (*EwsSessionResp, error)
	grpc.ClientStream
}

type ewsServiceEwsSessionClient struct {
	grpc.ClientStream
}

func (x *ewsServiceEwsSessionClient) Send(m *EwsSessionReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *ewsServiceEwsSessionClient) Recv() (*EwsSessionResp, error) {
	m := new(EwsSessionResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EwsServiceServer is the server API for EwsService service.
// All implementations must embed UnimplementedEwsServiceServer
// for forward compatibility
type EwsServiceServer interface {
	EwsSession(EwsService_EwsSessionServer) error
	mustEmbedUnimplementedEwsServiceServer()
}

// UnimplementedEwsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEwsServiceServer struct {
}

func (UnimplementedEwsServiceServer) EwsSession(EwsService_EwsSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method EwsSession not implemented")
}
func (UnimplementedEwsServiceServer) mustEmbedUnimplementedEwsServiceServer() {}

// UnsafeEwsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EwsServiceServer will
// result in compilation errors.
type UnsafeEwsServiceServer interface {
	mustEmbedUnimplementedEwsServiceServer()
}

func RegisterEwsServiceServer(s grpc.ServiceRegistrar, srv EwsServiceServer) {
	s.RegisterService(&EwsService_ServiceDesc, srv)
}

func _EwsService_EwsSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EwsServiceServer).EwsSession(&ewsServiceEwsSessionServer{stream})
}

type EwsService_EwsSessionServer interface {
	Send(*EwsSessionResp) error
	Recv() (*EwsSessionReq, error)
	grpc.ServerStream
}

type ewsServiceEwsSessionServer struct {
	grpc.ServerStream
}

func (x *ewsServiceEwsSessionServer) Send(m *EwsSessionResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *ewsServiceEwsSessionServer) Recv() (*EwsSessionReq, error) {
	m := new(EwsSessionReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// EwsService_ServiceDesc is the grpc.ServiceDesc for EwsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EwsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ewsgeneratedclient.EwsService",
	HandlerType: (*EwsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "EwsSession",
			Handler:       _EwsService_EwsSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ews.proto",
}
