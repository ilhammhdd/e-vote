package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"math/big"
	"strings"
	"time"
)

type Header struct {
	Algorithm string `json:"alg"`
	Type      string `json:"typ"`
}

type Payload struct {
	Id        string      `json:"jti,omitempty"`
	Audience  string      `json:"aud,omitempty"`
	Issuer    string      `json:"iss,omitempty"`
	IssuedAt  int64       `json:"iat,omitempty"`
	Expired   int64       `json:"exp,omitempty"`
	NotBefore int64       `json:"nbf,omitempty"`
	Claims    interface{} `json:"claims,omitempty"`
}

type Signature struct {
	Hashed []byte   `json:"hashed"`
	R      *big.Int `json:"r"`
	S      *big.Int `json:"s"`
}

type Token struct {
	Header    *Header   `json:"header"`
	Payload   *Payload  `json:"payload"`
	Signature Signature `json:"signature"`
}

func GenerateSignedToken(claims ...interface{}) (string, error) {
	privateKey := getPrivateFromPem("private_key.pem")
	publicKey := getPublicKeyFromPEM("public_key.pem")

	token := &Token{
		Header: &Header{
			Algorithm: "ECDSA",
			Type:      "JWT",
		},
		Payload: &Payload{
			Audience: "client e-vote",
			Issuer:   "server e-vote",
			IssuedAt: time.Now().UnixNano() / 1000000,
			Expired:  (time.Now().UnixNano() / 1000000) + 21600000,
		},
	}

	if len(claims) != 0 {
		token.Payload.Claims = claims[0]
	}

	jsonHeader, err := json.Marshal(token.Header)
	if err != nil {
		PanicRecover(err)
		return "", err
	}
	encodedHeader := base64.StdEncoding.EncodeToString(jsonHeader)

	jsonPayload, err := json.Marshal(token.Payload)
	if err != nil {
		PanicRecover(err)
		return "", err
	}
	encodedPayload := base64.StdEncoding.EncodeToString(jsonPayload)

	hasher := sha256.New()
	hasher.Write(jsonPayload)
	hashed := hasher.Sum(nil)
	r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed)
	if err != nil {
		PanicRecover(err)
		return "", err
	}

	if ecdsa.Verify(publicKey, hashed, r, s) {
		// Info("User", User.(claims), "generated hash, r, and s verified")
		Info("Generate hash, r, and s verified")
	} else {
		err := errors.New("Generated hash, r, and s fails to verify")
		PanicRecover(err)
		return "", err
	}

	token.Signature.Hashed = hashed
	token.Signature.R = r
	token.Signature.S = s
	jsonSignature, err := json.Marshal(token.Signature)
	if err != nil {
		PanicRecover(err)
	}
	signature := base64.StdEncoding.EncodeToString(jsonSignature)

	signedToken := strings.Join([]string{encodedHeader, encodedPayload, signature}, ".")

	return signedToken, nil
}

func VerifySignedToken(tokenString string) bool {
	publicKey := getPublicKeyFromPEM("public_key.pem")

	tokenParts := strings.Split(tokenString, ".")

	decodedSignature, err := base64.StdEncoding.DecodeString(tokenParts[2])
	if err != nil {
		PanicRecover(err)
		return false
	}

	var signature Signature
	err = json.Unmarshal([]byte(decodedSignature), &signature)
	if err != nil {
		PanicRecover(err)
		return false
	}

	if !ecdsa.Verify(publicKey, signature.Hashed, signature.R, signature.S) {
		return false
	}

	signature.Hashed[0] ^= 0xff
	if ecdsa.Verify(publicKey, signature.Hashed, signature.R, signature.S) {
		PanicRecover(errors.New("Verify always true"))
		return false
	}
	return true
}
