package utils

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hastpassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hastpassword), nil
}

func CompareHast(password, hashpassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashpassword), []byte(password))
	return err == nil
}

func CreateJwt(secret string, userID string) string {
	claim := jwt.MapClaims{
		"userid": userID,
		"exp":    time.Now().Unix() + 3600000,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, _ := token.SignedString([]byte(secret))

	return signedToken
}

func VerifyJwt(token, secret string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	if !parsedToken.Valid {
		return "", errors.New("token is not valid")
	}

	claims := parsedToken.Claims.(jwt.MapClaims)
	id, _ := claims["userid"].(string)

	return id, nil
}
