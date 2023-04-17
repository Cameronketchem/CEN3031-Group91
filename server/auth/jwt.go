package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(secret string, addr string, id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"address": addr,
		"user_id": id,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(secret))
	return tokenString, err
}

func VerifyJWT(secret string, tokenStr string) (string, int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

	if err != nil {
		return "", 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims["address"].(string), claims["user_id"].(int), nil
	}

	return "", 0, err
}
