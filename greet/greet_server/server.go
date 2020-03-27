package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

type server struct {
}

func (*server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	fN := req.GetGreeting().GetFirtName()
	result := "Hola " + fN
	res := &greetpb.GreetingResponse{
		Result: result,
	}
	return res, nil
}
func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	fN := req.GetGreeting().GetFirtName()
	for i := 0; i < 10; i++ {
		result := "Hola " + fN + " number " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1 * time.Second)
	}
	return nil
}
func (*server) LongGreet(stream greetpb.GreetService_LongGreetServer) error {
	var result string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&greetpb.LongGreetResponse{
				Result: result,
			})
		}
		if err != nil {
			log.Fatal(err)
		}
		result += "Hello " + req.GetGreeting().GetFirtName() + "! "
	}
}

func (*server) GreetEveryOne(stream greetpb.GreetService_GreetEveryOneServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}

		result := "Hola " + req.GetGreeting().GetFirtName()
		sendErr := stream.Send(&greetpb.GreetEveryOneResponse{
			Result: result,
		})
		if sendErr != nil {
			log.Fatal(sendErr)
		}
	}

}
func (*server) GreetWithDeadLines(ctx context.Context, request *greetpb.WithDeadLineRequest) (*greetpb.WithDeadLineResponse, error) {
	time.Sleep(3 * time.Second)
	if ctx.Err() == context.DeadlineExceeded {
		return nil, status.Errorf(codes.DeadlineExceeded, "The client canceled request")
	}
	req := request.GetGreeting().GetFirtName()

	return &greetpb.WithDeadLineResponse{Result: req}, nil
}
func main() {

	opts := []grpc.ServerOption{}

	tls := true

	s := grpc.NewServer()

	if tls {

		certFile := "../ssl/server.crt"
		keyFile := "../ssl/server.pem"

		creds, sslErr := credentials.NewServerTLSFromFile(certFile, keyFile)

		if sslErr != nil {
			log.Fatal(sslErr)
			return
		}
		opts = append(opts, grpc.Creds(creds))
	}
	s := grpc.NewServer(opts...)

	greetpb.RegisterGreetServiceServer(s, &server{})

	lis, _ := net.Listen("tcp", "localhost:50051")

	// error handling omitted
	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server is started")

	/*
		list, err := net.Listen("tcp", "0.0.0.0:50051")

		if err != nil {
			log.Fatal(err)
		}

		s := grpc.NewServer()
		greetpb.RegisterGreetServiceServer(s, &server{})

		if err := s.Serve(list); err != nil {
			log.Fatal(err)
		}
	*/
}
