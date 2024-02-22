package utils

import "golang.org/x/crypto/bcrypt"

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
