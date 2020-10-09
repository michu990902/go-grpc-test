package main

import (
	"context"
	"fmt"
	"log"

	"github.com/michu990902/go-pb-test/blog/blogpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Blog client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cold not connect: %v", err)
	}

	defer cc.Close()
	c := blogpb.NewBlogServiceClient(cc)

	fmt.Println("Create the blog")
	blog := &blogpb.Blog{
		AuthorId: "Test",
		Title:    "Test2",
		Content:  "Test2",
	}
	createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{
		Blog: blog,
	})
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	fmt.Printf("Blog has been created: %v\n", createBlogRes)
}
