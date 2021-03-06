// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package server

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ClientStreamClient is the client API for ClientStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientStreamClient interface {
	MaxAndAverage(ctx context.Context, opts ...grpc.CallOption) (ClientStream_MaxAndAverageClient, error)
}

type clientStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewClientStreamClient(cc grpc.ClientConnInterface) ClientStreamClient {
	return &clientStreamClient{cc}
}

func (c *clientStreamClient) MaxAndAverage(ctx context.Context, opts ...grpc.CallOption) (ClientStream_MaxAndAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientStream_ServiceDesc.Streams[0], "/grpc_stream.ClientStream/maxAndAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientStreamMaxAndAverageClient{stream}
	return x, nil
}

type ClientStream_MaxAndAverageClient interface {
	Send(*wrapperspb.Int64Value) error
	CloseAndRecv() (*TwoResponse, error)
	grpc.ClientStream
}

type clientStreamMaxAndAverageClient struct {
	grpc.ClientStream
}

func (x *clientStreamMaxAndAverageClient) Send(m *wrapperspb.Int64Value) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientStreamMaxAndAverageClient) CloseAndRecv() (*TwoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TwoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStreamServer is the server API for ClientStream service.
// All implementations must embed UnimplementedClientStreamServer
// for forward compatibility
type ClientStreamServer interface {
	MaxAndAverage(ClientStream_MaxAndAverageServer) error
	mustEmbedUnimplementedClientStreamServer()
}

// UnimplementedClientStreamServer must be embedded to have forward compatible implementations.
type UnimplementedClientStreamServer struct {
}

func (UnimplementedClientStreamServer) MaxAndAverage(ClientStream_MaxAndAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method MaxAndAverage not implemented")
}
func (UnimplementedClientStreamServer) mustEmbedUnimplementedClientStreamServer() {}

// UnsafeClientStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientStreamServer will
// result in compilation errors.
type UnsafeClientStreamServer interface {
	mustEmbedUnimplementedClientStreamServer()
}

func RegisterClientStreamServer(s grpc.ServiceRegistrar, srv ClientStreamServer) {
	s.RegisterService(&ClientStream_ServiceDesc, srv)
}

func _ClientStream_MaxAndAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientStreamServer).MaxAndAverage(&clientStreamMaxAndAverageServer{stream})
}

type ClientStream_MaxAndAverageServer interface {
	SendAndClose(*TwoResponse) error
	Recv() (*wrapperspb.Int64Value, error)
	grpc.ServerStream
}

type clientStreamMaxAndAverageServer struct {
	grpc.ServerStream
}

func (x *clientStreamMaxAndAverageServer) SendAndClose(m *TwoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientStreamMaxAndAverageServer) Recv() (*wrapperspb.Int64Value, error) {
	m := new(wrapperspb.Int64Value)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientStream_ServiceDesc is the grpc.ServiceDesc for ClientStream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientStream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc_stream.ClientStream",
	HandlerType: (*ClientStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "maxAndAverage",
			Handler:       _ClientStream_MaxAndAverage_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "client_stream.proto",
}
