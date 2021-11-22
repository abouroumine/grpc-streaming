package utils

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func OrderUnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	fmt.Println("Unary Server Interceptor: ", info.FullMethod)
	m, err := handler(ctx, req)
	fmt.Println("Error: ", err, " || Message: ", m)
	return m, err
}

type WrappedStream struct {
	grpc.ServerStream
}

func (w *WrappedStream) RecvMsg(m interface{}) error {
	return w.ServerStream.RecvMsg(m)
}

func (w *WrappedStream) SendMsg(m interface{}) error {
	return w.ServerStream.SendMsg(m)
}

func NewWrappedStream(s grpc.ServerStream) grpc.ServerStream {
	return &WrappedStream{s}
}

func OrderStreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	fmt.Println("Unary Server Interceptor: ", info.FullMethod)
	err := handler(srv, NewWrappedStream(ss))
	fmt.Println("Error: ", err)
	return err
}
