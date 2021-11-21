package main

import (
	client "abouroumine/ss_client/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"time"
)

const (
	Server = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(Server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()

	c := client.NewServerStreamClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	data, err := c.StartToFinish(ctx, &client.Request{
		Start:  1,
		Finish: 10,
	})
	if err != nil {
		return
	}
	for {
		v, er := data.Recv()
		if er == io.EOF {
			break
		}
		fmt.Println("The Value is: ", v)
	}
	fmt.Println("Finished")
}
