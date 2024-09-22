package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc/calculator/calculatorpb"
	"log"
	"net"
)

func main() {
	fmt.Println("Server starting...")

	listen, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	fmt.Println("Server listening...")

	grpcServer := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(grpcServer, &Server{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type Server struct {
	calculatorpb.CalculatorServiceServer
}

func (s *Server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	fmt.Println("Add function")

	num1 := req.GetNum1()
	num2 := req.GetNum2()

	result := num1 + num2

	res := &calculatorpb.AddResponse{
		Result: result,
	}

	return res, nil
}

func (s *Server) Sub(ctx context.Context, req *calculatorpb.SubRequest) (*calculatorpb.SubResponse, error) {
	fmt.Println("Sub function")

	num1 := req.GetNum1()
	num2 := req.GetNum2()
	result := num1 - num2

	res := &calculatorpb.SubResponse{
		Result: result,
	}

	return res, nil
}
