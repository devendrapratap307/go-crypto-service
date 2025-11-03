package crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

func SignRSA(priv *rsa.PrivateKey, message []byte) (string, error) {
	h := sha256.Sum256(message)
	sig, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, h[:], nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sig), nil
}
func VerifyRSA(pub *rsa.PublicKey, message []byte, signatureB64 string) (bool, error) {
	sig, err := base64.StdEncoding.DecodeString(signatureB64)
	if err != nil {
		return false, err
	}
	h := sha256.Sum256(message)
	if err := rsa.VerifyPSS(pub, crypto.SHA256, h[:], sig, nil); err != nil {
		return false, nil
	}
	return true, nil
}
