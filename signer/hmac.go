package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
)

// HMACSigner implements the Signer interface for HMAC
type HMACSigner struct {
	Secret []byte
}

// NewHMACSigner creates a new HMACSigner
func NewHMACSigner(secret string) *HMACSigner {
	return &HMACSigner{
		Secret: []byte(secret),
	}
}

// Sign signs the message using HMAC
func (h *HMACSigner) Sign(message []byte) (string, error) {
	hmac := hmac.New(sha256.New, h.Secret)
	hmac.Write(message)
	signature := hmac.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature), nil
}

// Verify verifies the HMAC signature
func (h *HMACSigner) Verify(message []byte, signature string) error {
	expectedSignature, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return err
	}

	// Create HMAC hash
	hash := hmac.New(sha256.New, h.Secret)
	hash.Write(message)
	computedSignature := hash.Sum(nil)

	// Use hmac.Equal for constant-time comparison
	if !hmac.Equal(computedSignature, expectedSignature) {
		return errors.New("HMAC signature mismatch")
	}
	return nil
}
