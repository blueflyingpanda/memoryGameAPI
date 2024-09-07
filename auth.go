package main

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"regexp"
	"time"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET")) // Replace with a strong secret
var tokenExpiry = time.Hour * 24

type JWTClaims struct {
	Login string `json:"login"`
	jwt.RegisteredClaims
}

func VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}

// GenerateJWT generates a JWT token for a given user
func GenerateJWT(login string) (string, error) {
	claims := jwt.MapClaims{
		"login": login,
		"exp":   time.Now().Add(tokenExpiry).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func IsValidSHA256Hash(s string) bool {
	if len(s) != 64 {
		return false
	}

	hexRegex := regexp.MustCompile("^[0-9a-fA-F]+$")
	return hexRegex.MatchString(s)
}
