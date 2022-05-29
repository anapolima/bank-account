package mocks

import "errors"

func ValidateDocumentSucceeded(document string) (string, error) {
	return document, nil
}

func ValidateDocumentFailed(document string) (string, error) {
	return "", errors.New("invalid document")
}
