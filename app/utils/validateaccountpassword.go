package utils

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

func ValidateAccountPassword(password string) (string, error) {
	log.Printf("Starting validate account password")
	password = strings.TrimSpace(password)

	if len(password) != 4 {
		return "", errors.New("invalid account password: expected 4 digits")
	}

	_, err := strconv.ParseInt(password, 10, 0)

	if err != nil {
		return "", errors.New("invalid password: expected numeric password")
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		log.Print(err)
		return "", errors.New("unable to encrypt password")
	}

	return string(encryptedPassword), nil
}
