package utils

import (
	"fmt"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateVerificationCode(len int) string {
	maxRange := TraillingDigits(len, 8)
	minRange := TraillingDigits(len, 1)
	verificationCode := rand.Intn(maxRange)
	verificationCode += minRange

	return fmt.Sprintf("%v", verificationCode)
}

func TraillingDigits(len, digit int) int {
	temp := 0
	for i := 0; i < len; i++ {
		temp = temp*10 + digit
	}
	return temp
}

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}
