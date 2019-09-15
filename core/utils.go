package core

import (
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

//GetEnv get env var
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

//HashAndSalt Generates a hashed password
func HashAndSalt(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

//ComparePasswords compare the user password with hashed password in database
func ComparePasswords(hashedPwd string, plainPwd string) bool {

	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
