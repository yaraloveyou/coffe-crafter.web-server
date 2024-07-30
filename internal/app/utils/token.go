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

func GetJwtKey() []byte {
	return []byte(config.SecretKey)
}

func GenerateJwt(u *model.User) (string, string, error) {
	accessExp := time.Now().Add(time.Second * time.Duration(config.AccessExp))
	refreshExp := time.Now().Add(time.Hour * time.Duration(config.RefreshExp))

	accessClaims := &Claims{
		Email:    u.Email,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExp),
		},
	}

	refreshClaims := &Claims{
		Email:    u.Email,
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExp),
		},
	}

	aToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	rToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)

	accessTokenString, err := aToken.SignedString(GetJwtKey())
	if err != nil {
		return "", "", err
	}

	refreshTokenString, err := rToken.SignedString(GetJwtKey())
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil // output refresh token in Redis storage
}

func RefreshJwt(token string) (string, string, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return GetJwtKey(), nil
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
