package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect : %v", err)
	}
	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)

	doUnaryRPC(c)

}

func doUnaryRPC(c greetpb.GreetServiceClient) {
	fmt.Println("Starting Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Rajat",
			LastName:  "Rawat",
		},
	}
	//fmt.Printf("Created Client %f", c)
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while Calling rpc: %v", err)
	}
	fmt.Printf("Response from Greet : %v", res.Result)
}
