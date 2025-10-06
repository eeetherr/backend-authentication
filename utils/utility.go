package utils

import (
	"ankit/authentication/configs"
	"bytes"
	"fmt"
	"html/template"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt/v5"

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

func HashPassword(password string) (string, error) {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed), nil
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

func GenerateJWT(email string) (string, error) {
	// Define expiration time
	jwtSecret := []byte(configs.Config.JWT.Secret)
	expirationTime := time.Now().Add(1 * time.Hour)

	// Create the JWT claims
	claims := jwt.MapClaims{
		"email": email,
		"exp":   expirationTime.Unix(),
	}

	// Create the token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
