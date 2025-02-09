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
		*ServerOption
	}
	Server interface {
		VerifySignature(ctx context.Context, req *http.Request) (err error)
		VerifySignatureAuth(ctx context.Context, req *http.Request) (err error)
	}
)

type (
	ServerOption struct {
		PublicKey  string
		Secret     string
		ClientKey  string
		SignerType signer.SignerType
	}
)

func NewServer(opt *ServerOption) Server {
	return &server{
		opt,
	}
}

// VerifySignature implements Server.
func (dep *server) VerifySignature(ctx context.Context, req *http.Request) (err error) {
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
	if dep.SignerType == signer.RSA_PKCS1 {
		err = dep.verifySignatureRSAPKCS1(ctx, signature, timestamp, body, method, path, query, clientKey)
	} else if dep.SignerType == signer.RSA_PSS {
		err = dep.verifySignatureRSAPSS(ctx, signature, timestamp, body, method, path, query, clientKey)
	} else if dep.SignerType == signer.HMAC {
		err = dep.verifySignatureHMAC(ctx, signature, timestamp, body, method, path, query, clientKey)
	} else {
		err = errors.New("signer type not supported")
	}
	return
}

func (dep *server) verifySignatureRSAPKCS1(ctx context.Context, signature string, timestamp string, body []byte, method string, path string, query string, clientKey string) (err error) {
	message := []byte(timestamp + string(body) + method + path + query + clientKey)

	decodedPublicKey, err := base64.StdEncoding.DecodeString(dep.PublicKey)
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

func (dep *server) verifySignatureHMAC(ctx context.Context, signature string, timestamp string, body []byte, method string, path string, query string, clientKey string) (err error) {
	message := []byte(timestamp + string(body) + method + path + query + clientKey)
	hmacSigner := signer.NewHMACSigner(dep.Secret)
	err = hmacSigner.Verify(message, signature)
	return
}

func (dep *server) verifySignatureRSAPSS(ctx context.Context, signature string, timestamp string, body []byte, method string, path string, query string, clientKey string) (err error) {
	message := []byte(timestamp + string(body) + method + path + query + clientKey)

	decodedPublicKey, err := base64.StdEncoding.DecodeString(dep.PublicKey)
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

func (dep *server) readBodyFromRequest(req *http.Request) (body []byte, err error) {
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
func (dep *server) VerifySignatureAuth(ctx context.Context, req *http.Request) (err error) {
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

	decodedPublicKey, err := base64.StdEncoding.DecodeString(dep.PublicKey)
	if err != nil {
		return
	}

	rsaPublicKey, err := signer.LoadPublicKeyFromPEM(decodedPublicKey)
	if err != nil {
		return
	}

	var sign signer.Signer

	switch dep.SignerType {
	case signer.RSA_PKCS1:
		sign = signer.NewRSA_PKCS1Signer(nil, rsaPublicKey)
	case signer.RSA_PSS:
		sign = signer.NewRSA_PSSSigner(nil, rsaPublicKey)
	case signer.HMAC:
		sign = signer.NewHMACSigner(dep.Secret)
	default:
		err = errors.New("signer type not supported")
		return
	}

	body, err := dep.readBodyFromRequest(req)
	if err != nil {
		return
	}

	err = sign.Verify(body, signature)
	if err != nil {
		return
	}

	return
}
