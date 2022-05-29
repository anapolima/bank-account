package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ValidateDocument(document string) (string, error) {
	log.Printf("Starting document validation with")
	document = strings.TrimSpace(document)

	if len(document) == 11 {
		if validateCpf(document) {
			return document, nil
		}

		return "", errors.New("invalid document: CPF")
	}

	if len(document) == 14 {
		if validateCnpj(document) {
			return document, nil
		}

		return "", errors.New("invalid document: CNPJ")
	}

	return "", errors.New("invalid document")
}

func validateCpf(cpf string) bool {
	regexExpression := `\d{11}`

	log.Printf("Checking CPF regex")
	r := regexp.MustCompile(regexExpression)
	d := r.Find([]byte(cpf))

	if d == nil {
		log.Printf("Invalid CPF format")
		return false
	}

	return checkCpf(cpf)
}

func checkCpf(cpf string) bool {
	sum := 0
	var rest int
	var aux int

	if cpf == "00000000000" {
		return false
	}

	runes := []rune(cpf)

	for i := 1; i <= 9; i++ {
		substring := string(runes[i-1 : i])
		value, err := strconv.Atoi(substring)

		if err != nil {
			return false
		}

		sum = sum + (value * (11 - i))
	}

	rest = (sum * 10) % 11

	if (rest == 10) || (rest == 11) {
		rest = 0
	}

	aux, err := strconv.Atoi(string(runes[9:10]))
	if err != nil {
		return false
	}
	if rest != aux {
		return false
	}

	sum = 0
	for i := 1; i <= 10; i++ {
		aux, err := strconv.Atoi(string(runes[i-1 : i]))

		if err != nil {
			return false
		}
		sum = sum + (aux * (12 - i))
	}

	rest = (sum * 10) % 11
	if (rest == 10) || (rest == 11) {
		rest = 0
	}

	aux, err = strconv.Atoi(string(runes[10:11]))
	if err != nil {
		return false
	}
	if rest != aux {
		return false
	}

	return true
}

func validateCnpj(cnpj string) bool {
	cnpj = strings.TrimSpace(cnpj)
	regexExpression := `\d{14}`
	r := regexp.MustCompile(regexExpression)
	d := r.Find([]byte(cnpj))

	if d == nil {
		log.Printf("Invalid CNPJ format")
		return false
	}

	return checkCnpj(cnpj)
}

func checkCnpj(cnpj string) bool {
	if cnpj == "" {
		return false
	}

	if len(cnpj) != 14 {
		return false
	}

	// Elimina CNPJs invalidos conhecidos
	if cnpj == "00000000000000" ||
		cnpj == "11111111111111" ||
		cnpj == "22222222222222" ||
		cnpj == "33333333333333" ||
		cnpj == "44444444444444" ||
		cnpj == "55555555555555" ||
		cnpj == "66666666666666" ||
		cnpj == "77777777777777" ||
		cnpj == "88888888888888" ||
		cnpj == "99999999999999" {
		return false
	}

	// Valida DVs
	var result int
	runes := []rune(cnpj)
	size := len(cnpj) - 2
	numbers := runes[0:size]
	digits := runes[size:]
	sum := 0
	pos := size - 7

	for i := size; i >= 1; i-- {

		aux, err := strconv.Atoi(string(numbers[size-i]))
		if err != nil {
			return false
		}
		sum += aux * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}

	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}

	aux, err := strconv.Atoi(string(digits[0]))
	if err != nil {
		return false
	}
	if result != aux {
		return false
	}

	size = size + 1
	numbers = runes[0:size]
	sum = 0
	pos = size - 7

	for i := size; i >= 1; i-- {
		aux, err := strconv.Atoi(string(numbers[size-i]))
		if err != nil {
			return false
		}
		sum += aux * pos
		pos--
		if pos < 2 {
			pos = 9
		}
	}

	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}

	aux, err = strconv.Atoi(string(digits[1]))
	if err != nil {
		return false
	}
	if result != aux {
		return false
	}

	return true
}
