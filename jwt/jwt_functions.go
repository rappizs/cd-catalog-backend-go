package jwt

import (
	"cd-catalog-backend-go/config"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	jwt "github.com/dgrijalva/jwt-go"
)

//Generate creates a token string
func Generate(userID uuid.UUID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	var myKey = []byte(config.SecretKey)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	return tokenString, nil
}
