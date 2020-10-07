package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/michu990902/go-pb-test/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	fmt.Printf("Sum: %v\n", req)
	a := req.GetA()
	b := req.GetB()
	result := a + b

	res := &calculatorpb.SumResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeDecomposition(req *calculatorpb.PrimeDecompositionRequest, stream calculatorpb.CalculatorService_PrimeDecompositionServer) error {
	fmt.Printf("Prime Decomposition: %v\n", req)
	n := req.GetA()
	var k int32
	k = 2

	for n > 1 {
		if n%k == 0 {
			res := &calculatorpb.PrimeDecompositionResponse{
				Result: k,
			}
			stream.Send(res)
			n = n / k
		} else {
			k++
		}
	}

	return nil
}

func (*server) ComputeAverage(stream calculatorpb.CalculatorService_ComputeAverageServer) error {
	fmt.Printf("Compute Average\n")
	result := float64(0)
	resultQuantity := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			result /= float64(resultQuantity)
			return stream.SendAndClose(&calculatorpb.ComputeAverageResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		number := req.GetNumber()
		result += float64(number)
		resultQuantity++
	}

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
