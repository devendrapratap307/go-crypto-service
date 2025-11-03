package service

import (
	"github.com/devendrapratap307/go-crypto-service/internal/crypto"
	"github.com/devendrapratap307/go-crypto-service/internal/keys"
)

type CryptoService struct {
	km *keys.KeyManager
}

func NewCryptoService(km *keys.KeyManager) *CryptoService { return &CryptoService{km: km} }
func (s *CryptoService) Encrypt(plaintext []byte) (ciphertext, nonce string,
	err error) {
	return crypto.EncryptAES(s.km.AESKey, plaintext)
}
func (s *CryptoService) Decrypt(ciphertext, nonce string) ([]byte, error) {
	return crypto.DecryptAES(s.km.AESKey, ciphertext, nonce)
}
func (s *CryptoService) Sign(message []byte) (string, error) {
	return crypto.SignRSA(s.km.RSAPrivate, message)
}
func (s *CryptoService) Verify(message []byte, sig string) (bool, error) {
	return crypto.VerifyRSA(s.km.RSAPublic, message, sig)
}
