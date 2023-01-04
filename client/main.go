package main

import (
	"context"
	"fmt"
	"log"

	// This import path is based on the name declaration in the go.mod,
	// and the gen/proto/go output location in the buf.gen.yaml.
	pb "finalProject/gen/proto/go"
	"google.golang.org/grpc"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	connectTo := "127.0.0.1:8080"
	conn, err := grpc.Dial(connectTo, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return fmt.Errorf("failed to connect to PetStoreService on %s: %w", connectTo, err)
	}
	defer conn.Close()
	log.Println("Connected to", connectTo)

	client := pb.NewTestApiClient(conn)
	stream, err := client.StreamPayments(context.Background(), &pb.StreamPaymentsRequest{})
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	done := make(chan bool)
	go func() {
		for {
			resp, err := stream.Recv()
			if err != nil {
				log.Fatalf("stream recv error %v", err)
			}
			log.Println("received", resp)
		}
	}()
	<-done

	return nil
}
