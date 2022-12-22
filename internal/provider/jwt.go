package provider

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret string = "secret"

func InitJWT(secret string) {
	jwtSecret = secret
	log.Println("Secret: ", jwtSecret)
}

func GenerateToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": userId,
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err	
	}

	return tokenString, nil
}
