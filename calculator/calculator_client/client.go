package main

import (
	"context"
	"fmt"
	"io"
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

	fmt.Println("Sum:")
	doUnary(c)
	fmt.Println("PrimeDecomposition:")
	doManyTimes(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		A: 5,
		B: 3,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC: %v", err)
	}

	log.Printf("Result: %v", res.Result)
}

func doManyTimes(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.PrimeDecompositionRequest{
		A: 120,
	}

	resStream, err := c.PrimeDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeDecomposition RPC: %v", err)
	}
	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}
		log.Printf("Result: %v", msg.GetResult())
	}
}
