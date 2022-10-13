package helper

import (
	"fmt"
	"mangojek-backend/exception"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = "SANGAT_RAHASIA"

func GenerateJWT(claims *jwt.MapClaims) (string, error) {
	var mySigningKey = []byte(secretKey)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webToken, err := token.SignedString(mySigningKey)
	exception.PanicIfNeeded(err)

	return webToken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	exception.PanicIfNeeded(err)
	return token, nil
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)
	exception.PanicIfNeeded(err)
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token")
}
