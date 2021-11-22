package main

import (
	pb "abouroumine/cs_client/protos"
	"abouroumine/cs_client/utils"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"time"
)

const (
	Server = "localhost:50050"
)

func main() {
	conn, err := grpc.Dial(Server, grpc.WithInsecure(), grpc.WithUnaryInterceptor(utils.OrderUnaryClientInterceptor), grpc.WithStreamInterceptor(utils.OrderStreamClientInterceptor))

	if err != nil {
		return
	}
	defer conn.Close()

	c := pb.NewClientStreamClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.MaxAndAverage(ctx)
	if err != nil {
		return
	}
	for i := 1; i < 10; i++ {
		err = stream.Send(wrapperspb.Int64(int64(i)))
		if err != nil {
			return
		}
		fmt.Println(i)
	}
	r, err := stream.CloseAndRecv()
	if err != nil {
		return
	}

	fmt.Println(r.Max, r.Average)
}
