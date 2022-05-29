package mocks

import "errors"

func ValidateDateSucceeded(date string) (string, error) {
	return date, nil
}

func ValidateDateFailed(date string) (string, error) {
	return "", errors.New("invalid date")
}
