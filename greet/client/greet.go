package main

import (
	"context"
	"log"

	pb "grpc-protobuf-go/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "FirstName"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", r.Result)
}
