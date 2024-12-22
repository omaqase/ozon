package main

import (
	"github.com/joho/godotenv"
	"github.com/oqamase/ozon/notification/internal/service"
	"github.com/oqamase/ozon/notification/pkg/notification"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}

	services, err := service.NewService()
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	notification.RegisterNotificationServiceServer(server, services)

	if err := server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
