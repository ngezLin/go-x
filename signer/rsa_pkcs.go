package signer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
)

// RSA_PKCS1Signer implements the Signer interface for RSA PKCS#1 v1.5
type RSA_PKCS1Signer struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewRSA_PKCS1Signer creates a new RSA_PKCS1Signer
func NewRSA_PKCS1Signer(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *RSA_PKCS1Signer {
	return &RSA_PKCS1Signer{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

// Sign signs the message using RSA PKCS#1 v1.5
func (r *RSA_PKCS1Signer) Sign(message []byte) (string, error) {
	hashed := sha256.Sum256(message)
	signature, err := rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA256, hashed[:])
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify verifies the RSA PKCS#1 v1.5 signature
func (r *RSA_PKCS1Signer) Verify(message []byte, signature string) error {
	hashed := sha256.Sum256(message)
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA256, hashed[:], sigBytes)
}
