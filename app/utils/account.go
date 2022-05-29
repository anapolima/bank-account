package utils

import (
	"math/rand"
	"time"
)

func GenerateDigit() int {
	rand.Seed(time.Now().UnixNano())
	min := 1
	max := 9
	return rand.Intn(max-min) + min
}

func GenerateAccountNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 100000
	max := 999999
	return rand.Intn(max-min) + min
}

func GenerateAgencyNumber() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	return rand.Intn(max-min) + min
}
