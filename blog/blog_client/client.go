package main

import (
	"context"
	"fmt"
	"io"
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

	//create blog
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
	blogID := createBlogRes.GetBlog().GetId()

	//read blog
	fmt.Println("Read the blog")
	_, err = c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: "test"})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	readBlogRes, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Printf("Error happened while reading: %v\n", err)
	}

	fmt.Printf("Blog was read: %v\n", readBlogRes)

	//update blog
	fmt.Println("Update the blog")
	newBlog := &blogpb.Blog{
		Id:       blogID,
		AuthorId: "Update Test",
		Title:    "Update Test2",
		Content:  "Update Test2",
	}
	updateRes, err := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{Blog: newBlog})
	if err != nil {
		fmt.Printf("Error happened while updating: %v\n", err)
	}
	fmt.Printf("Blog was updated: %v\n", updateRes)

	//delete blog
	fmt.Println("Delete the blog")
	deleteRes, err := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{BlogId: blogID})
	if err != nil {
		fmt.Printf("Error happened while deleting: %v\n", err)
	}
	fmt.Printf("Blog was deleted: %v\n", deleteRes)

	//list blog
	fmt.Println("List the blog")
	resStream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
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
		log.Printf("Result: %v\n", msg.GetBlog())
	}

}
