package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ketekdude/devcamp-2023-snd/service/database"
	shipperHandler "github.com/ketekdude/devcamp-2023-snd/service/server/handlers/shipper"
	pb "github.com/ketekdude/devcamp-2023-snd/service/server/handlers/shipper/proto"
	"github.com/ketekdude/devcamp-2023-snd/service/shippermodule"
)

func main() {
	dbConfig := database.Config{
		User:     "postgres",
		Password: "12345",
		DBName:   "devcamp",
		Port:     5433,
		Host:     "127.0.0.1",
		SSLMode:  "disable",
	}

	// Init DB connection
	log.Println("Initializing DB Connection")
	db := database.GetDatabaseConnection(dbConfig)

	// Init shipper usecase
	log.Println("Initializing Usecase")
	sm := shippermodule.NewShipperModule(db)

	// Init shipper handler
	log.Println("Initializing Handler")
	sh := shipperHandler.NewShipperHandler(sm)

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterShipperServer(grpcServer, sh)

	// register server using reflection
	reflection.Register(grpcServer)

	log.Println("Devcamp-2023-snd shipper service service is starting...")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
