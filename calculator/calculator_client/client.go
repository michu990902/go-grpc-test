package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/michu990902/go-pb-test/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	// fmt.Println("ComputeAverage:")
	// doComputeAverage(c)
	// fmt.Println("FindMaximum:")
	// doFindMaximum(c)
	fmt.Println("Error test:")
	doErrorUanry(c)
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
	numbers := []int32{1, 2, 3, 4}
	stream, err := c.ComputeAverage(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverange: %v\n", err)
	}

	for _, number := range numbers {
		fmt.Printf("Sending request: %v\n", number)
		stream.Send(&calculatorpb.ComputeAverageRequest{
			Number: number,
		})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while reciving response from ComputeAverange: %v\n", err)
	}
	fmt.Printf("ComputeAverange response: %v\n", res)
}

func doFindMaximum(c calculatorpb.CalculatorServiceClient) {
	numbers := []int32{1, 5, 3, 6, 2, 20}
	stream, err := c.FindMaximum(context.Background())
	if err != nil {
		log.Fatalf("Error while calling ComputeAverange: %v\n", err)
	}

	waitc := make(chan struct{})
	go func() {
		for _, number := range numbers {
			fmt.Printf("Sending request: %v\n", number)
			stream.Send(&calculatorpb.FindMaximumRequest{
				Number: number,
			})
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while receiving: %v\n", err)
				break
			}
			fmt.Printf("Received: %v\n", res.GetResult())
		}
		close(waitc)
	}()

	<-waitc
}

func doErrorUanry(c calculatorpb.CalculatorServiceClient) {
	//correct call
	doErrorCall(c, 10)
	//error call
	doErrorCall(c, -2)
}

func doErrorCall(c calculatorpb.CalculatorServiceClient, n int32) {
	req := &calculatorpb.SquareRootRequest{
		Number: n,
	}

	res, err := c.SquareRoot(context.Background(), req)
	if err != nil {
		resError, ok := status.FromError(err)
		if ok {
			fmt.Println(resError.Message())
			fmt.Println(resError.Code())
			if resError.Code() == codes.InvalidArgument {
				fmt.Println("We propably send a negative number!")
				return
			}
		} else {
			log.Fatalf("Big error while calling SquareRoot RPC: %v\n", err)
			return
		}
	}

	fmt.Printf("Result: %v\n", res.GetResult())
}
