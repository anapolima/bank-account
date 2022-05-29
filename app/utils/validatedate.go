package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ValidateDate(date string) (string, error) {
	log.Printf("Starting date validation with %s", date)
	date = strings.TrimSpace(date)
	regexExpression := `\d{4}-d{2}-d{2}$`
	r := regexp.MustCompile(regexExpression)
	d := r.Find([]byte(date))

	if d != nil {
		return "", errors.New("invalid date format: expected YYYY-MM-DD")
	}

	if checkDateValue(date) {
		return date, nil
	}

	return "", errors.New("invalid birthdate")
}

func checkDateValue(date string) bool {

	d := strings.Split(date, "-")

	day, err := strconv.Atoi(d[2])
	if err != nil {
		return false
	}
	month, err := strconv.Atoi(d[1])
	if err != nil {
		return false
	}
	year, err := strconv.Atoi(d[0])
	if err != nil {
		return false
	}

	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		if day > 31 || day < 1 {
			return false
		} else {
			return true
		}

	case 4, 6, 9, 11:
		if day > 30 || day < 1 {
			return false
		} else {
			return true
		}

	case 2:
		if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
			if day > 29 || day < 1 {
				return false
			} else {
				return true
			}
		} else {
			if day > 28 || day < 1 {
				return false
			} else {
				return true
			}
		}

	default:
		return false
	}
}
