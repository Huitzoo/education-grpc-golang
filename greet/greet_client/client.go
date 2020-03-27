package main

import (
	"context"
	"fmt"
	"greet/greetpb"
	"io"
	"time"

	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
)

func main() {

	//Client is created

	opts := []grpc.DialOption{}

	tls := true

	if tls {
		certFile := "../ssl/ca.crt"

		creds, _ := credentials.NewClientTLSFromFile(certFile, "")
		opts = append(opts, grpc.WithTransportCredentials(creds))

	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	cc, err := grpc.Dial("localhost:50051", opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

	//doServerStreaming(c)
	//doClientStreming(c)
	//doBiDirecciontal(c)
	//doUnaryWithDeadline(c, 1)
	//doUnaryWithDeadline(c, 5)
}

func doUnaryWithDeadline(c greetpb.GreetServiceClient, seconds int) {
	ctx, cancel := context.WithTimeout(
		context.Background(),
		time.Duration(seconds)*time.Second,
	)
	defer cancel()

	req := &greetpb.WithDeadLineRequest{
		Greeting: &greetpb.Greeting{
			FirtName: "Oscar",
			LastName: "Chavez",
		},
	}

	response, err := c.GreetWithDeadLines(ctx, req)
	if err != nil {
		statusErr, ok := status.FromError(err)
		if ok {
			if statusErr.Code() == codes.DeadlineExceeded {
				fmt.Println("Timeout execeded")
			} else {
				fmt.Println("unexpected error", err)
			}
		} else {

			log.Fatal(err)
		}
	} else {

		fmt.Println(response.Result)
	}
}

func doBiDirecciontal(c greetpb.GreetServiceClient) {
	request := []*greetpb.GreetEveryOneRequest{
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "oscar",
			},
		},
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "jorge",
			},
		},
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "cintia",
			},
		},
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "javier",
			},
		},
		&greetpb.GreetEveryOneRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "fernanda",
			},
		},
	}
	stream, err := c.GreetEveryOne(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range request {
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			result, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				break
			}
			if err != nil {
				log.Fatal(err)
				close(waitc)
				break
			}
			fmt.Println(result.GetResult())
		}
	}()

	<-waitc
}

func doClientStreming(c greetpb.GreetServiceClient) {
	request := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "oscar",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "jorge",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "cintia",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "javier",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirtName: "fernanda",
			},
		},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatal(err)
	}

	for _, req := range request {
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.Result)
}

func doServerStreaming(c greetpb.GreetServiceClient) {
	ctx := context.Background()
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirtName: "Oscar",
			LastName: "Chavez",
		},
	}
	resStream, err := c.GreetManyTimes(ctx, req)
	if err != nil {
		log.Fatal(err)
	}
	for {
		mgs, err := resStream.Recv()
		if err == io.EOF {
			fmt.Println("end of stream")
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(mgs.GetResult())
	}

}

func doUnary(c greetpb.GreetServiceClient) {
	ctx := context.Background()

	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirtName: "Oscar",
			LastName: "Chavez",
		},
	}

	response, _ := c.Greet(ctx, req)

	fmt.Println(response.Result)

}
