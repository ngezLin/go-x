package snap

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type (
	client struct {
		*ClientOption
	}

	HttpRequestConfig struct {
		Url string
	}

	Client interface {
		PrepareSignAsymmetric(ctx context.Context, req *http.Request) (err error)
		PrepareSignAuth(ctx context.Context, req *http.Request) (err error)
		PrepareHTTPRequest(ctx context.Context, req *http.Request) (err error)
	}
)

type (
	ClientOption struct {
		PrivateKey string
		ClientKey  string
		Secret     string
	}
)

func NewClient(opt *ClientOption) Client {
	return &client{
		opt,
	}
}

func (dep *client) exportPEMStrToPrivKey(priv []byte) *rsa.PrivateKey {
	block, _ := pem.Decode(priv)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

func (dep *client) generateRSAPrivateKey(ctx context.Context, key string) (privKey *rsa.PrivateKey, err error) {
	decodePrivKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		return privKey, errors.New("decode failed")
	}
	privKey = dep.exportPEMStrToPrivKey(decodePrivKey)
	return
}

// generateSign is a helper function to handle common signature generation logic
func (dep *client) generateSign(ctx context.Context, message string) (signatureStr string, err error) {
	msgHAsh := sha256.Sum256([]byte(message))

	privKeyFile, err := dep.generateRSAPrivateKey(ctx, dep.PrivateKey)
	if err != nil {
		return
	}

	signature, err := rsa.SignPSS(rand.Reader, privKeyFile, crypto.SHA256, msgHAsh[:], nil)
	if err != nil {
		return
	}

	signatureStr = base64.StdEncoding.EncodeToString(signature)
	return
}

// PrepareSignAsymmetric implements Client.
func (dep *client) PrepareSignAsymmetric(ctx context.Context, req *http.Request) (err error) {
	var body string
	if req.Body != nil {
		var bodyBytes []byte
		bodyBytes, _ = io.ReadAll(req.Body)
		// write back to request body
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		//unpretty request json
		dst := &bytes.Buffer{}
		_ = json.Compact(dst, []byte(bodyBytes))
		if dst != nil {
			body = dst.String()
		}
	}

	msg := sha256.Sum256([]byte(body))
	sha256_hash := hex.EncodeToString(msg[:])
	stringToSign := strings.ToLower(sha256_hash)

	message := fmt.Sprintf("%s:%s:%s:%s", req.Method, req.URL.Path, stringToSign, req.Header.Get(X_TIMESTAMP))

	signatureStr, err := dep.generateSign(ctx, message)
	if err != nil {
		return
	}
	// Put signature to header
	req.Header.Set(X_SIGNATURE, signatureStr)

	return
}

// PrepareSignAuth implements Client.
func (dep *client) PrepareSignAuth(ctx context.Context, req *http.Request) (err error) {
	if req.Header.Get(X_TIMESTAMP) == "" {
		req.Header.Set(X_TIMESTAMP, fmt.Sprintf("%d", time.Now().Unix()))
	}
	message := fmt.Sprintf("%s|%s", dep.ClientKey, req.Header.Get(X_TIMESTAMP))

	signatureStr, err := dep.generateSign(ctx, message)
	if err != nil {
		return
	}
	// Put signature to header
	req.Header.Set(X_SIGNATURE, signatureStr)
	return
}

// PrepareHTTPRequest implements Client.
func (dep *client) PrepareHTTPRequest(ctx context.Context, req *http.Request) (err error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set(X_CLIENT_KEY, dep.ClientKey)
	req.Header.Set(X_IDEMPOTENTCY, uuid.NewString())

	//prepare signature
	err = dep.PrepareSignAsymmetric(ctx, req)
	if err != nil {
		return
	}

	return
}
