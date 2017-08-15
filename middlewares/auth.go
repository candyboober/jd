package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jd/core"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// TODO: Don't forget to validate the exp is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return core.Secret, nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
}
