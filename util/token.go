package util

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWT(data any) (string, error) {
	var mySigningKey = []byte(TOKEN_SECRET_KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["data"] = data
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

func ExtractClaims(tokenStr string) (jwt.MapClaims, bool) {
	hmacSecretString := TOKEN_SECRET_KEY
	hmacSecret := []byte(hmacSecretString)
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return hmacSecret, nil
	})

	if err != nil {
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		log.Printf("Invalid JWT Token")
		return nil, false
	}
}
