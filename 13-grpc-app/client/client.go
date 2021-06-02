package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	//creating a client instance
	client := proto.NewAppServicesClient(conn)

	//creating the request object
	addRequest := &proto.CalculatorRequest{X: 100, Y: 200}

	//making the call to the server
	addResult, err := client.Add(context.Background(), addRequest)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(addResult.GetResult())
}
