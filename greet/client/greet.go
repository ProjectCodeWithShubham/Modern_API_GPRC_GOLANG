package main

import (
	"context"
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked and it is client file ")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Shubham",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v \n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}
