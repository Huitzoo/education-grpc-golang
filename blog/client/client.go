package main

import (
	"blog/blogpb"
	"context"
	"fmt"
	"io"
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
	//doUnaryReadBlog(c)
	//doUnaryUpdateBlog(c)
	doUnaryGetList(c)
}
func doUnaryGetList(c blogpb.BlogServiceClient) {
	req := &blogpb.ListBlogsRequest{}

	stream, err := c.ListBlogs(context.Background(), req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		result, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result.GetBlog())
	}
}

func doUnaryUpdateBlog(c blogpb.BlogServiceClient) {
	ctx := context.Background()
	req := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorID: "Huitzoo",
		},
		Iod: "5e7d8ea9f2d43d6b6b82cb79",
	}

	response, err := c.UpdateBlog(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)

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
			AuthorID: "oscar",
			Content:  "Que no te consuma la calle, ni la plata",
			Title:    "frases",
		},
	}

	response, err := c.CreateBlog(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response)

}
