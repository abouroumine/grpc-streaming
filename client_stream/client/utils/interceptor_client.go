package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func OrderUnaryClientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("Unary Client Interceptor Method: ", method)
	err := invoker(ctx, method, req, reply, cc, opts...)
	fmt.Println("Error: ", err)
	return err
}

type WrappedStream struct {
	grpc.ClientStream
}

func (w *WrappedStream) RecvMsg(m interface{}) error {
	return w.ClientStream.RecvMsg(m)
}

func (w *WrappedStream) SendMsg(m interface{}) error {
	return w.ClientStream.SendMsg(m)
}

func NewWrappedStream(s grpc.ClientStream) grpc.ClientStream {
	return &WrappedStream{s}
}

func OrderStreamClientInterceptor(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	s, err := streamer(ctx, desc, cc, method, opts...)
	if err != nil {
		fmt.Println("Error is: ", err)
		return nil, err
	}
	fmt.Println(method, " || ", s)
	return NewWrappedStream(s), nil
}
