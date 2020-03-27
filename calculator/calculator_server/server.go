package main

import (
	"calculator/calculatorpb"
	"io"
	"math"

	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
}

func (*server) Calculator(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	first := req.GetCalculate().GetFirtNumber()
	second := req.GetCalculate().GetSecondNumber()
	result := first + second

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeFactorDescomptitionManyTimeRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	number := req.GetPrime().GetNumber()
	var factor int32 = 2

	for {
		if number == 1 {
			break
		}
		if number%factor == 0 {
			number = number / factor
			res := &calculatorpb.PrimeFactorDescomptitionManyTimeResponse{
				Result: factor,
			}
			stream.Send(res)
		} else {
			factor++
		}

	}
	return nil
}

func (*server) LongCalculator(stream calculatorpb.CalculatorService_LongCalculatorServer) error {
	var number int32

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&calculatorpb.LongCalculatorResponse{
				Number: float64(number) / 4.0,
			})
		}
		if err != nil {
			log.Fatal(err)
		}

		number += req.GetNumber()
	}

}

func (*server) MaximumEveryOne(stream calculatorpb.CalculatorService_MaximumEveryOneServer) error {
	var number int32 = 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatal(err)
		}
		if number < req.GetCalculate().GetFirtNumber() {
			number = req.GetCalculate().GetFirtNumber()
			sendErr := stream.Send(&calculatorpb.MaximumEveryOneResponse{
				Result: number,
			})
			if sendErr != nil {
				log.Fatal(sendErr)
			}
		}
	}

}

func (*server) SquareRoots(ctx context.Context, request *calculatorpb.SquareRootsRequest) (*calculatorpb.SquareRootsResponse, error) {
	req := request.GetNumber()
	if req < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received negative number: %v", req),
		)
	}
	result := math.Sqrt(float64(req))
	return &calculatorpb.SquareRootsResponse{Result: result}, nil
}

func main() {
	fmt.Println("Server is started")
	list, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	// Register reflection service on gRPC server.
	s := grpc.NewServer()

	reflection.Register(s)

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(list); err != nil {
		log.Fatal(err)
	}

}
