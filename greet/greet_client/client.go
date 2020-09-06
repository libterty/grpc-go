package main

import (
	"../greetpb"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	fmt.Println("Client Serve")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Connect error: %v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Client Created: %f", c)
}
