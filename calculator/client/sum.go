package main

import (
	"context"
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/calculator/proto"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked and it is client file ")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 3,
	})

	if err != nil {
		log.Fatalf("Could not greet: %v \n", err)
	}

	log.Printf("Sum: %d\n", res.Result)
}
