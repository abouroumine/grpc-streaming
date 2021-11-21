package main

import (
	server2 "abouroumine/two_server/protos"
	"abouroumine/two_server/utils"
	"google.golang.org/grpc"
	"net"
)

const (
	Server = ":50052"
)

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", Server)
	if err != nil {
		return
	}

	server2.RegisterTwoStreamServer(s, &utils.Server2{})

	err = s.Serve(lis)
	if err != nil {
		return
	}
}
