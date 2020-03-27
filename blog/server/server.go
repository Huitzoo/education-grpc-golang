package main

import (
	"blog/blogpb"
	"blog/db"
	"blog/models"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var mongodb db.DB

type server struct {
}

func (*server) CreateBlog(ctx context.Context, request *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := request.GetBlog()

	data := models.ItemBlog{
		AuthorID: blog.GetAuthorID(),
		Content:  blog.GetContent(),
		Title:    blog.GetTitle(),
	}

	_, err := mongodb.InsertOne(0, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error:%v", err))
	}

	return &blogpb.CreateBlogResponse{Result: "ok"}, nil
}

func (*server) ReadBlog(ctx context.Context, request *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogID := request.GetBlogID()

	data, err := mongodb.FindByID(0, blogID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error:%v", err))
	}
	response := &models.ItemBlog{}
	if err := data.Decode(response); err != nil {
		log.Fatal(err)
	}

	return &blogpb.ReadBlogResponse{
		Blog: &blogpb.Blog{
			Id:       1,
			AuthorID: response.AuthorID,
			Title:    response.Title,
			Content:  response.Content,
		},
	}, nil
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	fmt.Println("Server is started")

	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	// Register reflection service on gRPC server.
	s := grpc.NewServer()
	mongodb = db.DB{}
	mongodb.ConnectDB()

	blogpb.RegisterBlogServiceServer(s, &server{})

	go func() {

		if err := s.Serve(list); err != nil {
			log.Fatal(err)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("stopped server")
	s.Stop()
	fmt.Println("stopped listener")
	list.Close()
	mongodb.CloseConnnection()
}
