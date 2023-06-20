package main

import (
	"context"
	"io"
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/greet/proto"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest1{
		FirstName: "shubham12",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v \n", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading the Stream : %v \n", err)
		}

		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}

}
