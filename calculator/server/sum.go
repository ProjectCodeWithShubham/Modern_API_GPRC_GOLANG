package main

import (
	"context"
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum Function was invoked with %v", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
