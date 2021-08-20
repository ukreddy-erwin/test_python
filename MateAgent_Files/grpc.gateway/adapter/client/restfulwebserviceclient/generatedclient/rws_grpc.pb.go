// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package rwspb

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

// RestfulWSServiceClient is the client API for RestfulWSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RestfulWSServiceClient interface {
	RestfulWSSession(ctx context.Context, opts ...grpc.CallOption) (RestfulWSService_RestfulWSSessionClient, error)
}

type restfulWSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRestfulWSServiceClient(cc grpc.ClientConnInterface) RestfulWSServiceClient {
	return &restfulWSServiceClient{cc}
}

func (c *restfulWSServiceClient) RestfulWSSession(ctx context.Context, opts ...grpc.CallOption) (RestfulWSService_RestfulWSSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &RestfulWSService_ServiceDesc.Streams[0], "/rwsgeneratedclient.RestfulWSService/RestfulWSSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &restfulWSServiceRestfulWSSessionClient{stream}
	return x, nil
}

type RestfulWSService_RestfulWSSessionClient interface {
	Send(*RWSSessionReq) error
	Recv() (*RWSSessionResp, error)
	grpc.ClientStream
}

type restfulWSServiceRestfulWSSessionClient struct {
	grpc.ClientStream
}

func (x *restfulWSServiceRestfulWSSessionClient) Send(m *RWSSessionReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *restfulWSServiceRestfulWSSessionClient) Recv() (*RWSSessionResp, error) {
	m := new(RWSSessionResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RestfulWSServiceServer is the server API for RestfulWSService service.
// All implementations must embed UnimplementedRestfulWSServiceServer
// for forward compatibility
type RestfulWSServiceServer interface {
	RestfulWSSession(RestfulWSService_RestfulWSSessionServer) error
	mustEmbedUnimplementedRestfulWSServiceServer()
}

// UnimplementedRestfulWSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRestfulWSServiceServer struct {
}

func (UnimplementedRestfulWSServiceServer) RestfulWSSession(RestfulWSService_RestfulWSSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method RestfulWSSession not implemented")
}
func (UnimplementedRestfulWSServiceServer) mustEmbedUnimplementedRestfulWSServiceServer() {}

// UnsafeRestfulWSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RestfulWSServiceServer will
// result in compilation errors.
type UnsafeRestfulWSServiceServer interface {
	mustEmbedUnimplementedRestfulWSServiceServer()
}

func RegisterRestfulWSServiceServer(s grpc.ServiceRegistrar, srv RestfulWSServiceServer) {
	s.RegisterService(&RestfulWSService_ServiceDesc, srv)
}

func _RestfulWSService_RestfulWSSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RestfulWSServiceServer).RestfulWSSession(&restfulWSServiceRestfulWSSessionServer{stream})
}

type RestfulWSService_RestfulWSSessionServer interface {
	Send(*RWSSessionResp) error
	Recv() (*RWSSessionReq, error)
	grpc.ServerStream
}

type restfulWSServiceRestfulWSSessionServer struct {
	grpc.ServerStream
}

func (x *restfulWSServiceRestfulWSSessionServer) Send(m *RWSSessionResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *restfulWSServiceRestfulWSSessionServer) Recv() (*RWSSessionReq, error) {
	m := new(RWSSessionReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// RestfulWSService_ServiceDesc is the grpc.ServiceDesc for RestfulWSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RestfulWSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rwsgeneratedclient.RestfulWSService",
	HandlerType: (*RestfulWSServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RestfulWSSession",
			Handler:       _RestfulWSService_RestfulWSSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "rws.proto",
}
