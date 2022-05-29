package mocks

import (
	"errors"
)

func ValidateAccountPasswordSucceeded(password string) (string, error) {
	return password, nil
}

func ValidateAccountPasswordFailed(password string) (string, error) {
	return "", errors.New("invalid password")
}
