package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"log"
	"time"

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

	//client streaming
	fmt.Println("Client streaming")
	data := []int64{30, 10, 20, 40, 50}
	avgClientStream, err := client.CalculateAverage(context.Background())
	if err != nil {
		log.Fatalln(err)
	}

	for _, no := range data {
		fmt.Println("Sending : ", no)
		req := &proto.AverageRequest{No: no}
		avgClientStream.Send(req)
		time.Sleep(500 * time.Millisecond)
	}

	avgResponse, err := avgClientStream.CloseAndRecv()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Average : ", avgResponse.GetResult())
}
