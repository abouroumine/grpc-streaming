package main

import (
	client "abouroumine/two_client/protos"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"
)

const (
	Server = "localhost:50052"
)

func main() {
	conn, err := grpc.Dial(Server, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return
	}
	defer conn.Close()

	c := client.NewTwoStreamClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.CountMultiDiv(ctx)
	if err != nil {
		return
	}
	for i := int64(0); i < 10; i++ {
		err = stream.Send(&client.Request{
			First:  i + int64(2),
			Second: i + int64(1),
		})
		if err != nil {
			return
		}
		data := new(client.Response)
		err = stream.RecvMsg(data)
		if err != nil {
			return
		}
		fmt.Println("Message is: ", data.Multi, data.Div)
	}
}
