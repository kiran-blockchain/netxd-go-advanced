package main

import (
	"context"
	"fmt"
	
	"log"

	pb "grpc-mongo/grpc/profile" 
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProfileServiceClient(conn)
	

	response, err := client.CreateProfile(context.Background(), &pb.Profile{Name: "Kiran"})
	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}

	fmt.Printf("Response: %s\n", response.Name)
}
