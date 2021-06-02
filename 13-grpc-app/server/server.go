package main

import (
	"context"
	"fmt"
	"grpc-app/proto"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
}

func (s *server) Add(ctx context.Context, req *proto.CalculatorRequest) (*proto.CalculatorResponse, error) {
	x := req.GetX()
	y := req.GetY()
	result := x + y
	return &proto.CalculatorResponse{Result: result}, nil
}

func (s *server) CalculateAverage(stream proto.AppServices_CalculateAverageServer) error {
	var sum int64 = 0
	var count int64 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			//send the response
			avg := sum / count
			response := &proto.AverageResponse{Result: int64(avg)}
			stream.SendAndClose(response)
			break
		}
		if err != nil {
			log.Fatalln(err)
		}
		sum += req.GetNo()
		count++
	}
	return nil
}

//server streaming
func (s *server) GeneratePrimes(req *proto.PrimeNumbersRequest, stream proto.AppServices_GeneratePrimesServer) error {
	rangeStart := req.GetRangeStart()
	rangeEnd := req.GetRangeEnd()
	for i := rangeStart; i <= rangeEnd; i++ {
		if isPrime(i) {
			fmt.Println("Sending Prime Number : ", i)
			res := &proto.PrimeNumbersResponse{No: i}
			stream.Send(res)
			time.Sleep(500 * time.Millisecond)
		}
	}
	return nil
}

func isPrime(no int32) bool {
	for i := int32(2); i <= (no / int32(2)); i++ {
		if no%int32(i) == 0 {
			return false
		}
	}
	return true
}

func main() {
	listener, err := net.Listen("tcp", "localhost:50000")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterAppServicesServer(grpcServer, &server{})
	e := grpcServer.Serve(listener)
	if e != nil {
		log.Fatalln(e)
	}
}
