package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc/calculator/calculatorpb"
	"log"
)

func main() {
	fmt.Println("Client...")
	conn, err := grpc.NewClient(":12345", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()

	c := calculatorpb.NewCalculatorServiceClient(conn)

	ctx := context.Background()

	fmt.Println("--------------Add Function")
	addRequest := calculatorpb.AddRequest{
		Num1: 1,
		Num2: 3,
	}
	addRes, err := c.Add(ctx, &addRequest)
	if err != nil {
		log.Fatalf("could not add: %v", err)
	}
	fmt.Printf("Result: %v\n", addRes.Result)

	fmt.Println("--------------Sub Function")
	subRequest := calculatorpb.SubRequest{
		Num1: 5,
		Num2: 2,
	}
	subRes, err := c.Sub(ctx, &subRequest)
	if err != nil {
		log.Fatalf("could not sub: %v", err)
	}
	fmt.Printf("Result: %v\n", subRes.Result)
}
