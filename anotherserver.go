package main

import (
	"context"
	"log"
	"net"

	"github.com/byteshiva/simplegrpc/subtraction"

	"google.golang.org/grpc"
)

type additionServer struct {
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
	subtraction.RegisterSubtractionServer(grpcServer, newAddServer())
	grpcServer.Serve(lis)
}
