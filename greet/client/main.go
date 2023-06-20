package main

import (
	"log"

	pb "github.com/ProjectCodeWithShubham/Modern_API_GPRC_GOLANG/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v \n ", err)
	}
	log.Printf("Client is running.......")
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
	doGreetManyTimes(c)

}
