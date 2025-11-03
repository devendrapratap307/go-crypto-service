package keys

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/devendrapratap307/go-crypto-service/internal/config"
)

type KeyManager struct {
	AESKey     []byte
	RSAPrivate *rsa.PrivateKey
	RSAPublic  *rsa.PublicKey
}

func NewKeyManager(cfg *config.Config) (*KeyManager, error) {
	km := &KeyManager{}

	// Load AES key
	b, err := ioutil.ReadFile(cfg.AESKeyPath)
	if err != nil {
		return nil, err
	}
	km.AESKey = b

	// Load RSA private key
	privPEM, err := ioutil.ReadFile(filepath.Join(cfg.RSADir, "private.pem"))
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(privPEM)
	if block == nil {
		return nil, errors.New("invalid private key PEM")
	}
	// priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	// Try PKCS#1 first (traditional RSA)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		// If that fails, try PKCS#8 (modern format)
		key, err2 := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err2 != nil {
			return nil, err2
		}
		// Assert type: must be *rsa.PrivateKey
		rsaPriv, ok := key.(*rsa.PrivateKey)
		if !ok {
			return nil, errors.New("not an RSA private key")
		}
		priv = rsaPriv
	}

	if err != nil {
		return nil, err
	}
	km.RSAPrivate = priv
	km.RSAPublic = &priv.PublicKey
	return km, nil
}
