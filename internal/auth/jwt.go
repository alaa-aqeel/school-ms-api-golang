package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/alaa-aqeel/school-ms-api-golang/config"
	"github.com/golang-jwt/jwt/v5"
)

func GetTokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}

func CreateToken(sub string) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": sub,
		"exp": config.JWT_EXPIRE_TOKEN,
	})

	return accessToken.SignedString(config.JWT_SECRET_KEY)
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	var claims jwt.MapClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return config.JWT_SECRET_KEY, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
