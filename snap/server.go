package snap

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/super-saga/go-x/signer"
)

type (
	server struct {
		PublicKey  string
		Secret     string
		ClientKey  string
		SignerType signer.SignerType
	}
	Server interface {
		VerifySignature(ctx context.Context, req *http.Request) (err error)
		VerifySignatureAuth(ctx context.Context, req *http.Request) (err error)
	}
)

func NewServer(publicKey, secret, clientKey string, signerType signer.SignerType) Server {
	return &server{
		PublicKey:  publicKey,
		Secret:     secret,
		ClientKey:  clientKey,
		SignerType: signerType,
	}
}

// VerifySignature implements Server.
func (s *server) VerifySignature(ctx context.Context, req *http.Request) (err error) {
	// 1. Get the signature from the request header
	signature := req.Header.Get(X_SIGNATURE)
	// 2. Get the timestamp from the request header
	timestamp := req.Header.Get(X_TIMESTAMP)
	// 3. Get the request body
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return
	}
	// 4. Get the request method
	method := req.Method
	// 5. Get the request path
	path := req.URL.Path
	// 6. Get the request query
	query := req.URL.RawQuery
	// 7. Get the request client key
	clientKey := req.Header.Get(X_CLIENT_KEY)
	// 8. Verify the signature
	if s.SignerType == signer.RSA_PKCS1 {
		err = s.verifySignatureRSAPKCS1(ctx, signature, timestamp, body, method, path, query, clientKey)
	} else if s.SignerType == signer.RSA_PSS {
		err = s.verifySignatureRSAPSS(ctx, signature, timestamp, body, method, path, query, clientKey)
	} else if s.SignerType == signer.HMAC {
		err = s.verifySignatureHMAC(ctx, signature, timestamp, body, method, path, query, clientKey)
	} else {
		err = errors.New("signer type not supported")
	}
	return
}

func (s *server) verifySignatureRSAPKCS1(ctx context.Context, signature string, timestamp string, body []byte, method string, path string, query string, clientKey string) (err error) {
	message := []byte(timestamp + string(body) + method + path + query + clientKey)

	decodedPublicKey, err := base64.StdEncoding.DecodeString(s.PublicKey)
	if err != nil {
		return
	}

	rsaPublicKey, err := signer.LoadPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return
	}

	rsaPKCS1Signer := signer.NewRSA_PKCS1Signer(nil, rsaPublicKey)
	err = rsaPKCS1Signer.Verify(message, signature)
	return
}

func (s *server) verifySignatureHMAC(ctx context.Context, signature string, timestamp string, body []byte, method string, path string, query string, clientKey string) (err error) {
	message := []byte(timestamp + string(body) + method + path + query + clientKey)
	hmacSigner := signer.NewHMACSigner(s.Secret)
	err = hmacSigner.Verify(message, signature)
	return
}

func (s *server) verifySignatureRSAPSS(ctx context.Context, signature string, timestamp string, body []byte, method string, path string, query string, clientKey string) (err error) {
	message := []byte(timestamp + string(body) + method + path + query + clientKey)

	decodedPublicKey, err := base64.StdEncoding.DecodeString(s.PublicKey)
	if err != nil {
		return
	}

	rsaPublicKey, err := signer.LoadPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return
	}

	rsaPSSSigner := signer.NewRSA_PSSSigner(nil, rsaPublicKey)
	err = rsaPSSSigner.Verify(message, signature)
	return
}

func (s *server) readBodyFromRequest(req *http.Request) (body []byte, err error) {
	if req.Body != nil {
		var bodyBytes []byte
		bodyBytes, _ = io.ReadAll(req.Body)
		// write back to request body
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
		//unpretty request json
		dst := &bytes.Buffer{}
		json.Compact(dst, []byte(bodyBytes))
		body = dst.Bytes()
		if body == nil {
			err = errors.New("body is nil")
		}
	}
	return
}

// VerifySignatureAuth implements Server.
func (s *server) VerifySignatureAuth(ctx context.Context, req *http.Request) (err error) {
	signature := req.Header.Get(X_SIGNATURE)
	timestamp := req.Header.Get(X_TIMESTAMP)

	if signature == "" {
		err = errors.New("signature not found")
		return
	}

	if timestamp == "" {
		err = errors.New("timestamp not found")
		return
	}

	decodedPublicKey, err := base64.StdEncoding.DecodeString(s.PublicKey)
	if err != nil {
		return
	}

	rsaPublicKey, err := signer.LoadPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return
	}

	var sign signer.Signer

	switch s.SignerType {
	case signer.RSA_PKCS1:
		sign = signer.NewRSA_PKCS1Signer(nil, rsaPublicKey)
	case signer.RSA_PSS:
		sign = signer.NewRSA_PSSSigner(nil, rsaPublicKey)
	case signer.HMAC:
		sign = signer.NewHMACSigner(s.Secret)
	default:
		err = errors.New("signer type not supported")
		return
	}

	body, err := s.readBodyFromRequest(req)
	if err != nil {
		return
	}

	err = sign.Verify(body, signature)
	if err != nil {
		return
	}

	return
}
