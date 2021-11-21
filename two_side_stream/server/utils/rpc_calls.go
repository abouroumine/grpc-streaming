package utils

import (
	server "abouroumine/two_server/protos"
	"fmt"
	"io"
)

type Server2 struct {
	server.UnimplementedTwoStreamServer
}

func (s *Server2) CountMultiDiv(stream server.TwoStream_CountMultiDivServer) error {
	for {
		v, err := stream.Recv()
		if err == io.EOF {
			return stream.Send(&server.Response{
				Multi: v.First * v.Second,
				Div:   v.First / v.Second,
			})
		}
		if err != nil {
			return err
		}
		msg := server.Response{
			Multi: v.First * v.Second,
			Div:   v.First / v.Second,
		}
		fmt.Println(msg.Div, msg.Multi)
		err = stream.SendMsg(&msg)
		if err != nil {
			return err
		}
	}
}
