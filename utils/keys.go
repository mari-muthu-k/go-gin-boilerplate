package utils

import (
	"os"

	"github.com/golang-jwt/jwt"
	"github.com/mari-muthu-k/gin-template/globals"
	"github.com/mari-muthu-k/gin-template/model/appschema"
)

func LoadCertificateKeys()error {
	publicKey,err := os.ReadFile("keys/public.pem")
	if err != nil {
		return err
	}

	privateKey,err := os.ReadFile("keys/private.pem")
	if err != nil {
		return err
	}

	PrivatePem, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return err
	}
	
	PublicPem, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return err
	}
	

	globals.AppKeys = appschema.CertificateKeys{
		PublicKeyPem:   PublicPem,
		PrivateKey:     PrivatePem,
	}

	return nil
}