package main

import (
	"context"
	"log"
	"net"

	"github.com/byteshiva/simplegrpc/addition"
	"github.com/byteshiva/simplegrpc/subtraction"

	"google.golang.org/grpc"
)

type additionServer struct {
}

func (s *additionServer) Add(c context.Context, addRequest *addition.AddRequest) (*addition.AddResponse, error) {
	result := addRequest.Number + addRequest.AnotherNumber
	response := addition.AddResponse{
		Sum: int64(result),
	}
	return &response, nil
}

func (s *additionServer) Minus(c context.Context, minusRequest *subtraction.MinusRequest) (*subtraction.MinusResponse, error) {
	result := minusRequest.Number - minusRequest.AnotherNumber
	response := subtraction.MinusResponse{
		Subtract: int64(result),
	}
	return &response, nil
}

func newAddServer() *additionServer {
	return &additionServer{}
}

func main() {
	lis, err := net.Listen("tcp", "localhost:10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	addition.RegisterAdditionServer(grpcServer, newAddServer())
	subtraction.RegisterSubtractionServer(grpcServer, newAddServer())
	grpcServer.Serve(lis)
}
