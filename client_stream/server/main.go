package main

import (
	pb "abouroumine/cs_server/protos"
	"abouroumine/cs_server/utils"
	"google.golang.org/grpc"
	"net"
)

const (
	Server = ":50050"
)

func main() {
	s := grpc.NewServer()
	lis, err := net.Listen("tcp", Server)
	if err != nil {
		return
	}

	pb.RegisterClientStreamServer(s, &utils.Server2{})

	err = s.Serve(lis)
	if err != nil {
		return
	}

}
