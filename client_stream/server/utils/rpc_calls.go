package utils

import (
	pb "abouroumine/cs_server/protos"
	"io"
)

type Server2 struct {
	pb.UnimplementedClientStreamServer
}

var allValues int64 = 0
var lastValue int64 = 0

func (s *Server2) MaxAndAverage(stream pb.ClientStream_MaxAndAverageServer) error {
	for {
		v, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.TwoResponse{
				Max:     lastValue,
				Average: allValues / lastValue,
			})
		}

		if err != nil {
			return err
		}

		allValues += v.Value
		lastValue = v.Value
	}
}
