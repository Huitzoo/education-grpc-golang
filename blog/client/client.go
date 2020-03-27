package main

import (
	"blog/blogpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {

	opts := []grpc.DialOption{}

	opts = append(opts, grpc.WithInsecure())

	cc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	//doUnary(c)
	doUnaryReadBlog(c)
	//doServerStreaming(c)
	//doClientStreming(c)
	//doBiDirecciontal(c)
	//doUnaryWithDeadline(c, 1)
	//doUnaryWithDeadline(c, 5)
}

func doUnaryReadBlog(c blogpb.BlogServiceClient) {
	ctx := context.Background()
	req := &blogpb.ReadBlogRequest{
		BlogID: "5e7d8e60f2d43d6b6b82cb78",
	}
	response, err := c.ReadBlog(ctx, req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(response.GetBlog())

}

func doUnary(c blogpb.BlogServiceClient) {
	ctx := context.Background()
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorID: "1245",
			Content:  "Hola, este es un contenido de este blog alv",
			Title:    "Alv",
		},
	}

	response, err := c.CreateBlog(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Result)

}
