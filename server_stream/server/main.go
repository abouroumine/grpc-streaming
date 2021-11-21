package main

import (
	server "abouroumine/ss_server/protos"
	"abouroumine/ss_server/utils"
	"google.golang.org/grpc"
	"net"
)

const (
	Server = ":50051"
)

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", Server)
	if err != nil {
		return
	}

	server.RegisterServerStreamServer(s, &utils.Server2{})

	err = s.Serve(lis)
	if err != nil {
		return
	}
}
