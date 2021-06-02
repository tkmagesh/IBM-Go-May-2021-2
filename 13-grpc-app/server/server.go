package main

import (
	"context"
	"grpc-app/proto"
	"io"
	"log"
	"net"

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
