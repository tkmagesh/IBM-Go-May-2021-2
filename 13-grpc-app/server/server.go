package main

import (
	"context"
	"grpc-app/proto"
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
