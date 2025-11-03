package main

import (
	"log"
	"net"

	"github.com/devendrapratap307/go-crypto-service/internal/wire"

	pb "github.com/devendrapratap307/go-crypto-service/proto/crypto"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"
)

func main() {
	// Initialize via wire (generated wire_gen.go)
	appHandlers, err := wire.InitializeApp()
	if err != nil {
		log.Fatalf("failed to init app: %v", err)
	}

	// Start REST server
	go func() {
		fiberApp := fiber.New()
		appHandlers.RestHandler.Register(fiberApp)
		log.Println("REST server listening on :8080")
		if err := fiberApp.Listen(":8080"); err != nil {
			log.Fatalf("REST server error: %v", err)
		}
	}()

	// Start gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCryptoServiceServer(grpcServer, appHandlers.GRPCServer)
	log.Println("gRPC server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC server error: %v", err)
	}

}
