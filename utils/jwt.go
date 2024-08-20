package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mari-muthu-k/gin-template/globals"
	"github.com/mari-muthu-k/gin-template/model/appschema"
)

func generateJWTClaims(data *appschema.JwtData, ttl time.Duration) jwt.MapClaims {
	now := time.Now().UTC()
	claims := make(jwt.MapClaims)
	claims["email"] = data.Email
	claims["first_name"] = data.FirstName
	claims["last_name"] = data.LastName
	claims["id"] = data.ID
	claims["exp"] = now.Add(ttl).Unix() // The expiration time after which the token must be disregarded.
	claims["jwt_created"] = now.Unix()
	return claims
}


func GenerateJWTToken(ttl time.Duration, data *appschema.JwtData) (signedToken string, err error) {

	claims := generateJWTClaims(data, ttl)
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(globals.AppKeys.PrivateKey)
	if err != nil {
		fmt.Println("Error generating token: ", err)
		return
	}

	return token, err
}

func ParseJWTToken(token string) (jwt.MapClaims, error) {

	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return globals.AppKeys.PublicKeyPem, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, err
	}

	return claims, nil
}