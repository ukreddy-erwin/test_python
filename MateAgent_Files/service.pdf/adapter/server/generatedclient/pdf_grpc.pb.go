// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pdfpb

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

// PdfServiceClient is the client API for PdfService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PdfServiceClient interface {
	PdfSession(ctx context.Context, opts ...grpc.CallOption) (PdfService_PdfSessionClient, error)
}

type pdfServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPdfServiceClient(cc grpc.ClientConnInterface) PdfServiceClient {
	return &pdfServiceClient{cc}
}

func (c *pdfServiceClient) PdfSession(ctx context.Context, opts ...grpc.CallOption) (PdfService_PdfSessionClient, error) {
	stream, err := c.cc.NewStream(ctx, &PdfService_ServiceDesc.Streams[0], "/pdfgeneratedclient.PdfService/PdfSession", opts...)
	if err != nil {
		return nil, err
	}
	x := &pdfServicePdfSessionClient{stream}
	return x, nil
}

type PdfService_PdfSessionClient interface {
	Send(*PdfSessionReq) error
	Recv() (*PdfSessionResp, error)
	grpc.ClientStream
}

type pdfServicePdfSessionClient struct {
	grpc.ClientStream
}

func (x *pdfServicePdfSessionClient) Send(m *PdfSessionReq) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pdfServicePdfSessionClient) Recv() (*PdfSessionResp, error) {
	m := new(PdfSessionResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PdfServiceServer is the server API for PdfService service.
// All implementations must embed UnimplementedPdfServiceServer
// for forward compatibility
type PdfServiceServer interface {
	PdfSession(PdfService_PdfSessionServer) error
	mustEmbedUnimplementedPdfServiceServer()
}

// UnimplementedPdfServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPdfServiceServer struct {
}

func (UnimplementedPdfServiceServer) PdfSession(PdfService_PdfSessionServer) error {
	return status.Errorf(codes.Unimplemented, "method PdfSession not implemented")
}
func (UnimplementedPdfServiceServer) mustEmbedUnimplementedPdfServiceServer() {}

// UnsafePdfServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PdfServiceServer will
// result in compilation errors.
type UnsafePdfServiceServer interface {
	mustEmbedUnimplementedPdfServiceServer()
}

func RegisterPdfServiceServer(s grpc.ServiceRegistrar, srv PdfServiceServer) {
	s.RegisterService(&PdfService_ServiceDesc, srv)
}

func _PdfService_PdfSession_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PdfServiceServer).PdfSession(&pdfServicePdfSessionServer{stream})
}

type PdfService_PdfSessionServer interface {
	Send(*PdfSessionResp) error
	Recv() (*PdfSessionReq, error)
	grpc.ServerStream
}

type pdfServicePdfSessionServer struct {
	grpc.ServerStream
}

func (x *pdfServicePdfSessionServer) Send(m *PdfSessionResp) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pdfServicePdfSessionServer) Recv() (*PdfSessionReq, error) {
	m := new(PdfSessionReq)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PdfService_ServiceDesc is the grpc.ServiceDesc for PdfService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PdfService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pdfgeneratedclient.PdfService",
	HandlerType: (*PdfServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PdfSession",
			Handler:       _PdfService_PdfSession_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "pdf.proto",
}
