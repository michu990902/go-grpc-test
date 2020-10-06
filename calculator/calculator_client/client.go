package main

import (
	"context"
	"fmt"
	"log"

	"github.com/michu990902/go-pb-test/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Client test")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cold not connect: %v", err)
	}

	defer cc.Close()
	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.CalculatorRequest{
		A: 5,
		B: 3,
	}

	res, err := c.CalculatorAdd(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling CalculatorAdd RPC: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}
