package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func EncryptAES(key, plaintext []byte) (ciphertextB64, nonceB64 string, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return
	}
	ct := gcm.Seal(nil, nonce, plaintext, nil)
	ciphertextB64 = base64.StdEncoding.EncodeToString(ct)
	nonceB64 = base64.StdEncoding.EncodeToString(nonce)
	return
}
func DecryptAES(key []byte, ciphertextB64, nonceB64 string) (plaintext []byte, err error) {
	ct, err := base64.StdEncoding.DecodeString(ciphertextB64)
	if err != nil {
		return
	}
	nonce, err := base64.StdEncoding.DecodeString(nonceB64)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return
	}
	plaintext, err = gcm.Open(nil, nonce, ct, nil)
	return
}
