package main

import (
	"finalProject/adapter"
	pb "finalProject/gen/proto/go"
	"finalProject/helpers"
	"finalProject/payment"
	"finalProject/product"
	"finalProject/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
)

func main() {
	roomManager := service.GetRoomManager()

	listenOn := "127.0.0.1:8080"
	listener, err := net.Listen("tcp", listenOn)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterTestApiServer(server, adapter.NewTestApiServer(roomManager))
	reflection.Register(server)

	go func() {
		log.Println("gRPC Server Started on : 8080")
		db, err := helpers.GetDatabase()
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&product.Product{})
		if err != nil {
			log.Fatalln(err)
		}
		err = db.AutoMigrate(&payment.Payment{})
		if err != nil {
			log.Fatalln(err)
		}
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	// Stop the server
	log.Println("stopping the server")
	server.Stop()
	log.Println("server stopped")
}
