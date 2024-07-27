package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/model"
)

type Claims struct {
	Email    string
	Username string
	jwt.RegisteredClaims
}

var (
	jwtKey = []byte("secret_key") // output in config file
)

func GenerateJwt(u *model.User) (string, string, error) {
	aExp := time.Now().Add(time.Second * 20)   // output in config file
	rExp := time.Now().Add(time.Hour * 24 * 7) // output in config file

	aClaims := &Claims{
		Email:    u.Email,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(aExp),
		},
	}

	rClaims := &Claims{
		Email:    u.Email,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(rExp),
		},
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aClaims)
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rClaims)

	aTokenString, err := aToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	rTokenString, err := rToken.SignedString(jwtKey)
	if err != nil {
		return "", "", err
	}

	return aTokenString, rTokenString, nil // output refresh token in Redis storage
}

func RefreshJwt(token string) (string, string, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return "", "", err
	}

	u := &model.User{
		Email:    claims.Email,
		Username: claims.Username,
	}

	return GenerateJwt(u)
}

func ExtractTokenFromHandler(r *http.Request) string {
	bearerToken := r.Header.Get("Authorization")
	token := strings.Split(bearerToken, " ")[1]

	return token
}
