package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateApiJWT(secret, issuer string, userId, username string) (string, error) {
	claims := jwt.MapClaims{
		"userId":        userId,
		"username":      username,
		"authenticated": true,
		"exp":           time.Now().AddDate(0, 1, 0).Unix(),
		"iat":           time.Now().Unix(),
		"iss":           issuer,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
