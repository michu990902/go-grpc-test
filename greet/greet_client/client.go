package main

import (
	"fmt"
	"log"

	"github.com/michu990902/go-pb-test/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client test")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cold not connect: %v", err)
	}

	defer cc.Close()
	c := greetpb.NewGreetServiceClient(cc)
	fmt.Printf("Created client: %f", c)
}
