package signer

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

// SignerType defines the type of signer
type SignerType string

const (
	HMAC      SignerType = "hmac"
	RSA_PKCS1 SignerType = "rsa_pkcs1"
	RSA_PSS   SignerType = "rsa_pss"
)

// Signer is the interface for signing and verifying
type Signer interface {
	Sign(message []byte) (string, error)
	Verify(message []byte, signature string) error
}

// Utility functions to generate RSA keys
func GenerateRSAKeys(bits int) (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, bits)
}

func PublicKeyToPEM(pub *rsa.PublicKey) []byte {
	pubASN1 := x509.MarshalPKCS1PublicKey(pub)
	return pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: pubASN1})
}

func PrivateKeyToPEM(priv *rsa.PrivateKey) []byte {
	privASN1 := x509.MarshalPKCS1PrivateKey(priv)
	return pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: privASN1})
}

// LoadPrivateKeyFromPEM loads an RSA private key from PEM encoded data
func LoadPrivateKeyFromPEM(pemData []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("invalid PEM data")
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

// LoadPublicKeyFromPEM loads an RSA public key from PEM encoded data
func LoadPublicKeyFromPEM(pemData []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemData)
	if block == nil || block.Type != "PUBLIC KEY" {
		return nil, errors.New("invalid PEM data")
	}
	pub, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return pub, nil
}
