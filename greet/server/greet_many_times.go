package main

import (
	"fmt"
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest1, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked server main with: %v \n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s , number %d ", in.FirstName, i)
		stream.Send(&pb.GreetResponse1{
			Result: res,
		})
	}
	return nil
}
