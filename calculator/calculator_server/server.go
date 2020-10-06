package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/michu990902/go-pb-test/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) CalculatorAdd(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Add tow numbers: %v", req)
	a := req.GetA()
	b := req.GetB()
	result := a + b

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
