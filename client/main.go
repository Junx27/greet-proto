package main

import (
	"context"
	"fmt"
	"go-greet/pb"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreeterClient(conn)

	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	var statusChoice int
	fmt.Println("Enter your status:")
	fmt.Println("1. Hadir")
	fmt.Println("2. Tidak Hadir")
	fmt.Print("Please choose 1 or 2: ")

	fmt.Scanln(&statusChoice)
	var status string
	if statusChoice == 1 {
		status = "Hadir"
	} else if statusChoice == 2 {
		status = "Tidak Hadir"
	} else {
		status = "Tidak Tahu"
	}

	req := &pb.HelloRequest{Name: name + " - " + status}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatalf("Error while calling SayHello: %v", err)
	}

	log.Printf("Terima Kasih! %s", res.GetMessage())
}
