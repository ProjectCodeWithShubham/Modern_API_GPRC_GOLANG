package main

import (
	"context"
	"log"
	"time"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet Was invoked")

	reqs := []*pb.GreetRequest{

		{FirstName: "Shubham_client_streaming"},
		{FirstName: "utkarsh_client_streaming"},
		{FirstName: "ujwal_client_streaming"},
		{FirstName: "samyak_client_streaming"},
	}
	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error While calling longGreet %v \n ", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req:%v \n ", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while receiving response from Long Greet : \n %v", err)
	}
	log.Printf("LongGreet: %s \n ", res.Result)
}
