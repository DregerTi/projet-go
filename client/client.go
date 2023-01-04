package main

import (
	"context"
	pb "finalProject/gen/proto/go"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
	}

	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Message: "Hello"})
	if err != nil {
		log.Println(err)
	}

	fmt.Print(resp.Message)
}
