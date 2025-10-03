package utils

import (
	"bytes"
	"fmt"
	"html/template"
	"math/rand"

	"golang.org/x/crypto/bcrypt"
)

func GenerateVerificationCode(len int) string {
	maxRange := TrailingDigits(len, 8)
	minRange := TrailingDigits(len, 1)
	verificationCode := rand.Intn(maxRange)
	verificationCode += minRange

	return fmt.Sprintf("%v", verificationCode)
}

func TrailingDigits(len, digit int) int {
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

func RenderTemplate(tplText string, data interface{}) (string, error) {
	t, err := template.New("tpl").Parse(tplText)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
