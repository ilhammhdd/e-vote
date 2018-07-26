package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"
)

func GenerateECDSAKeyPairToFile() {
	_, errPrivate := os.Stat("private_key.pem")
	_, errPublic := os.Stat("public_key.pem")

	if os.IsNotExist(errPrivate) && os.IsNotExist(errPublic) {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			PanicRecover(err)
		}

		publicKey := &privateKey.PublicKey

		pemPrivateKeyFile, err := os.Create("private_key.pem")
		if err != nil {
			PanicRecover(err)
		}

		pemPublicKeyFile, err := os.Create("public_key.pem")
		if err != nil {
			PanicRecover(err)
		}

		defer pemPrivateKeyFile.Close()
		defer pemPublicKeyFile.Close()

		marshalPrivate, err := x509.MarshalECPrivateKey(privateKey)
		if err != nil {
			PanicRecover(err)
		}

		pemPrivateKeyBlock := &pem.Block{
			Type:  "E256 PRIVATE KEY",
			Bytes: marshalPrivate,
		}

		asn1Bytes, err := x509.MarshalPKIXPublicKey(publicKey)
		if err != nil {
			PanicRecover(err)
		}

		pemPublicKeyBlock := &pem.Block{
			Type:  "E256 PUBLIC KEY",
			Bytes: asn1Bytes,
		}

		err = pem.Encode(pemPrivateKeyFile, pemPrivateKeyBlock)
		if err != nil {
			PanicRecover(err)
		}

		err = pem.Encode(pemPublicKeyFile, pemPublicKeyBlock)
		if err != nil {
			PanicRecover(err)
		}
	}
}

func getPublicKeyFromPEM(filePath string) *ecdsa.PublicKey {
	publicKeyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		PanicRecover(err)
		return nil
	}

	publicPemBlock, _ := pem.Decode(publicKeyData)
	if publicPemBlock == nil || publicPemBlock.Type != "E256 PUBLIC KEY" {
		PanicRecover(errors.New("Failed to decode PEM block containing public key"))
		return nil
	}

	parsedPublicKey, err := x509.ParsePKIXPublicKey(publicPemBlock.Bytes)
	if err != nil {
		if cert, err := x509.ParseCertificate(publicPemBlock.Bytes); err == nil {
			parsedPublicKey = cert.PublicKey
		} else {
			PanicRecover(err)
			return nil
		}
	}

	publicKey, ok := parsedPublicKey.(*ecdsa.PublicKey)
	if !ok {
		PanicRecover(errors.New("not ECDSA public key"))
		return nil
	}

	return publicKey
}

func getPrivateFromPem(filePath string) *ecdsa.PrivateKey {
	privateKeyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		PanicRecover(err)
		return nil
	}

	privatePemBlock, _ := pem.Decode(privateKeyData)
	if privatePemBlock == nil || privatePemBlock.Type != "E256 PRIVATE KEY" {
		PanicRecover(errors.New("Failed to decode PEM block containing private key"))
		return nil
	}

	privateKey, err := x509.ParseECPrivateKey(privatePemBlock.Bytes)
	if err != nil {
		PanicRecover(err)
		return nil
	}

	return privateKey
}
