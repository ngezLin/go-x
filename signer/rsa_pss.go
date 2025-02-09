package signer

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// RSA_PPSSigner implements the Signer interface for RSA PSS
type RSA_PSSSigner struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

// NewRSA_PSSSigner creates a new RSA_PSSSigner
func NewRSA_PSSSigner(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) *RSA_PSSSigner {
	return &RSA_PSSSigner{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
	}
}

// Sign signs the message using RSA PSS
func (r *RSA_PSSSigner) Sign(message []byte) (res string, err error) {
	if r.PrivateKey == nil {
		err = errors.New("private key is nil")
		return
	}
	hashed := sha256.Sum256(message)
	signature, err := rsa.SignPSS(rand.Reader, r.PrivateKey, crypto.SHA256, hashed[:], nil)
	if err != nil {
		return
	}
	res = base64.StdEncoding.EncodeToString(signature)
	return
}

// Verify verifies the RSA PSS signature
func (r *RSA_PSSSigner) Verify(message []byte, signature string) (err error) {
	if r.PublicKey == nil {
		err = errors.New("public key is nil")
		return
	}

	hashed := sha256.Sum256(message)
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return
	}
	err = rsa.VerifyPSS(r.PublicKey, crypto.SHA256, hashed[:], sigBytes, nil)
	return
}
