package middleware

import (
	"cd-catalog-backend-go/config"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

//Auth is the authorization middleware
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/users/login" {
			next.ServeHTTP(w, r)
		}

		if r.Header.Get("token") != "" {
			token, err := jwt.Parse(r.Header.Get("token"), func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte(config.SecretKey), nil
			})

			if err != nil {
				fmt.Println(err)
				w.WriteHeader(401)
				return
			}

			if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
				next.ServeHTTP(w, r)
			}
		} else {
			w.WriteHeader(401)
			return
		}

	})
}
