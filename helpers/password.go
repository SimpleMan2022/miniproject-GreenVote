package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(request string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(request), bcrypt.DefaultCost)
	return string(hashPass), err
}

func VerifyPassword(hashPass string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(password))
	return err
}
