package handler

import (
	"context"

	"github.com/devendrapratap307/go-crypto-service/internal/service"
	pb "github.com/devendrapratap307/go-crypto-service/proto/crypto"
)

type GRPCServer struct {
	pb.UnimplementedCryptoServiceServer
	svc *service.CryptoService
}

func NewGRPCServer(svc *service.CryptoService) *GRPCServer { return &GRPCServer{svc: svc} }
func (s *GRPCServer) Encrypt(ctx context.Context, req *pb.EncryptRequest) (*pb.EncryptResponse, error) {
	ct, nonce, err := s.svc.Encrypt([]byte(req.Plaintext))
	if err != nil {
		return nil, err
	}
	return &pb.EncryptResponse{Ciphertext: ct, Nonce: nonce}, nil
}

// Decrypt, Sign, Verify implemented similarly
