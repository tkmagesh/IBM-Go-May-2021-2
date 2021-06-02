package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
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

	//server streaming
	fmt.Println()
	fmt.Println("Generating Prime numbers from 25 to 75 (server streaming)")
	primeReq := &proto.PrimeNumbersRequest{
		RangeStart: 25,
		RangeEnd:   75,
	}
	primeNoStream, err := client.GeneratePrimes(context.Background(), primeReq)
	if err != nil {
		log.Fatalln(err)
	}
	for {
		res, err := primeNoStream.Recv()
		if err == io.EOF {
			fmt.Println("All prime generators are generated")
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Prime No : ", res.GetNo())
	}

	//bidirectional streaming
	doBiDiStreaming(client)
}

func doBiDiStreaming(c proto.AppServicesClient) {
	fmt.Println("Starting to do a BiDi Streaming RPC...")

	// we create a stream by invoking the client
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
		return
	}

	requests := []*proto.GreetEveryoneRequest{
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Stephane",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "John",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Lucy",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Mark",
			},
		},
		&proto.GreetEveryoneRequest{
			Greeting: &proto.Greeting{
				FirstName: "Piper",
			},
		},
	}

	waitc := make(chan struct{})
	// we send a bunch of messages to the client (go routine)
	go func() {
		// function to send a bunch of messages
		for _, req := range requests {
			fmt.Printf("Sending message: %v\n", req)
			stream.Send(req)
			time.Sleep(1000 * time.Millisecond)
		}
		stream.CloseSend()
	}()
	// we receive a bunch of messages from the client (go routine)
	go func() {
		// function to receive a bunch of messages
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	// block until everything is done
	<-waitc
}
