package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/utils"
)

var (
	jwtKey = []byte("secret_key") // output in config file
)

func Authenticate(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := utils.ExtractTokenFromHandler(r)
		token, err := jwt.ParseWithClaims(tokenString, &utils.Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	}
}
