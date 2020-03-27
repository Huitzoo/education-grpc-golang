package main

import (
	"calculator/calculatorpb"
	"io"
	"time"

	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"

	"fmt"
)

func main() {

	//Client is created
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)
	//unitaryCommunication(c)
	//doServerStream(c)
	//doClientStreaming(c)
	//biDirectionalStreaming(c)
	doErrorUnary(c)
}
func doErrorUnary(c calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SquareRootsRequest{
		Number: -1,
	}
	response, err := c.SquareRoots(context.Background(), req)

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			//error del usuario
			fmt.Println(e.Message())
			fmt.Println(e.Code())
		} else {
			//error del server
			log.Fatal(err)
		}
	} else {
		fmt.Println(response.GetResult())

	}
}

func biDirectionalStreaming(c calculatorpb.CalculatorServiceClient) {
	request := [6]int32{1, 5, 3, 6, 2, 20}
	stream, err := c.MaximumEveryOne(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range request {
			stream.Send(&calculatorpb.MaximumEveryOneRequest{
				Calculate: &calculatorpb.Calculator{
					FirtNumber:   req,
					SecondNumber: 0,
				},
			})
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
				close(waitc)
				log.Fatal(err)
				break
			}

			fmt.Print(result.GetResult())
		}
	}()

	<-waitc
}

func doClientStreaming(c calculatorpb.CalculatorServiceClient) {

	request := []*calculatorpb.LongCalculatorRequest{
		&calculatorpb.LongCalculatorRequest{
			Number: 1,
		},
		&calculatorpb.LongCalculatorRequest{
			Number: 2,
		},
		&calculatorpb.LongCalculatorRequest{
			Number: 3,
		},
		&calculatorpb.LongCalculatorRequest{
			Number: 4,
		},
	}

	stream, err := c.LongCalculator(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	for _, req := range request {
		stream.Send(req)
		time.Sleep(100 * time.Millisecond)
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result.GetNumber())
}

func doServerStream(c calculatorpb.CalculatorServiceClient) {
	ctx := context.Background()

	req := &calculatorpb.PrimeFactorDescomptitionManyTimeRequest{
		Prime: &calculatorpb.PrimeFactorDescomptition{
			Number: 120,
		},
	}

	response, _ := c.PrimeNumberDecomposition(ctx, req)
	for {
		factor, err := response.Recv()
		if err == io.EOF {
			fmt.Println()
			fmt.Println("end factor")
			break
		}
		fmt.Print(factor.Result, ",")
	}
}

func unitaryCommunication(c calculatorpb.CalculatorServiceClient) {

	ctx := context.Background()

	req := &calculatorpb.CalculatorRequest{
		Calculate: &calculatorpb.Calculator{
			FirtNumber:   1,
			SecondNumber: 10,
		},
	}

	response, _ := c.Calculator(ctx, req)

	fmt.Println(response.Result)

}
