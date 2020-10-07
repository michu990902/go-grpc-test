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

	// fmt.Println("Sum:")
	// doSum(c)
	// fmt.Println("PrimeDecomposition:")
	// doPrimeDecomposition(c)
	fmt.Println("ComputeAverage:")
	doComputeAverage(c)
}

func doSum(c calculatorpb.CalculatorServiceClient) {
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

func doPrimeDecomposition(c calculatorpb.CalculatorServiceClient) {
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

func doComputeAverage(c calculatorpb.CalculatorServiceClient) {
	requests := []*calculatorpb.ComputeAverageRequest{
		&calculatorpb.ComputeAverageRequest{
			Number: 1,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 2,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 3,
		},
		&calculatorpb.ComputeAverageRequest{
			Number: 4,
		},
	}

	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverange: %v\n", err)
	}

	for _, req := range requests {
		fmt.Printf("Sending reques: %v\n", req)
		stream.Send(req)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reciving response from ComputeAverange: %v\n", err)
	}
	fmt.Printf("ComputeAverange response: %v\n", res)
}
