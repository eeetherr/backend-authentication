package services

import (
	"ankit/authentication/dto/comms"
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

type commsService struct {
	SMTPHost string
	SMTPPort string
	Username string
	Password string
	From     string
}

//var CommsService = &commsService{}

func NewCommsService() *commsService {
	return &commsService{
		SMTPHost: "smtp.gmail.com",
		SMTPPort: "587",
		Username: "yytiyiy@gmail.com",
		Password: "your-app-password",
		From:     "your-email@gmail.com",
	}
} // SignUp handles user registration
func (s *commsService) SendEmail(to string, verificationCode string) error {

	//send email
	tmpl, err := template.ParseFiles("templates/verification_email.txt")
	if err != nil {
		return fmt.Errorf("failed to load email template: %w", err)
	}

	var data = comms.EmailData{
		Name: "User",
		Code: "hdh",
	} //after sending to user save into database

	var body bytes.Buffer
	err = tmpl.Execute(&body, data)
	if err != nil {
		return fmt.Errorf("failed to execute email template: %w", err)
	}

	// Split into subject and body (optional)
	subject := "Verify your email"
	message := fmt.Sprintf("To: %s\nSubject: %s\n\n%s", to, subject, body.String())

	// Send the email
	auth := smtp.PlainAuth("", s.Username, s.Password, s.SMTPHost)
	err = smtp.SendMail(
		s.SMTPHost+":"+s.SMTPPort,
		auth,
		s.From,
		[]string{to},
		[]byte(message),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil

}
