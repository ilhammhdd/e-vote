package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"log"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateRSAPairToFile() {
	_, errPrivate := os.Stat("rsa_private_key")
	_, errPublic := os.Stat("rsa_public_key.pem")

	if os.IsNotExist(errPrivate) && os.IsNotExist(errPublic) {

		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			log.Println(err)
		}

		publicKey := privateKey.PublicKey

		pemPrivateKeyFile, err := os.Create("rsa_private_key")
		if err != nil {
			log.Println(err)
		}

		pemPublicKeyFile, err := os.Create("rsa_public_key.pem")
		if err != nil {
			log.Println(err)
		}

		pemPrivateKeyBlock := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		}

		asn1Bytes, err := x509.MarshalPKIXPublicKey(&publicKey)
		if err != nil {
			log.Println(err)
		}

		pemPublicKeyBlock := &pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: asn1Bytes,
		}

		err = pem.Encode(pemPrivateKeyFile, pemPrivateKeyBlock)
		if err != nil {
			log.Println(err)
		}

		err = pem.Encode(pemPublicKeyFile, pemPublicKeyBlock)
		if err != nil {
			log.Println(err)
		}

		pemPrivateKeyFile.Close()
		pemPublicKeyFile.Close()
	}
}

func VerifyRSAFromPEM(tokenString string) error {
	key, err := ioutil.ReadFile("rsa_public_key.pem")
	if err != nil {
		log.Println(err)
	}
	parsedKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		log.Println(err)
	}

	parts := strings.Split(tokenString, ".")
	err = jwt.SigningMethodRS256.Verify(strings.Join(parts[0:2], "."), parts[2], parsedKey)
	parts = parts[:0]
	return err
}

func GenerateRSATokenWithClaims(claims jwt.Claims) (string, error) {
	key, err := ioutil.ReadFile("rsa_private_key")
	if err != nil {
		log.Println(err)
	}
	parsedKey, err := jwt.ParseRSAPrivateKeyFromPEM(key)
	if err != nil {
		log.Println(err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(parsedKey)
}
