package utils

import (
	server "abouroumine/ss_server/protos"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type Server2 struct {
	server.UnimplementedServerStreamServer
}

func (s *Server2) StartToFinish(req *server.Request, stream server.ServerStream_StartToFinishServer) error {
	start := req.Start
	finish := req.Finish
	for i := start; i < finish; i++ {
		err := stream.Send(wrapperspb.Int64(i))
		if err != nil {
			return err
		}
	}
	return nil
}
