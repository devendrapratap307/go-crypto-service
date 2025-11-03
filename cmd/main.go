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
	// restHandler, grpcServer, err := wire.InitializeApp()
	// if err != nil {
	// 	log.Fatalf("failed to initialize: %v", err)
	// }
	// // Start REST
	// go func() {
	// 	app := fiber.New()
	// 	restHandler.Register(app)
	// 	log.Println("REST listening :8080")
	// 	if err := app.Listen(":8080"); err != nil {
	// 		log.Fatalf("REST server error: %v", err)
	// 	}
	// }()
	// // Start gRPC
	// lis, err := net.Listen("tcp", ":50051")
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterCryptoServiceServer(s, grpcServer)
	// log.Println("gRPC listening :50051")
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("gRPC serve error: %v", err)
	// }

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
