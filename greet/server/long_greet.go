package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/greet/proto"
)

func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Printf("Long-Greet function was invoked ")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error While reading client stream: %v \n ", err)
		}
		res += fmt.Sprintf("Hello %s! \n", req.FirstName)

	}
}
