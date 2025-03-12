package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword はパスワードをハッシュ化する
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CompareHAshAndPasswordで比較する
func CheckPasswordHash(password,  hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
