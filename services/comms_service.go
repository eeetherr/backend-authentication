package services

import (
	"ankit/authentication/constants"
	"ankit/authentication/dto/comms"
	"ankit/authentication/repositories"
	"ankit/authentication/utils"
	"encoding/json"
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type commsService struct {
}

func (s *commsService) SendEmail(to string, data interface{}, templateName string) error {
	commsRepo := repositories.CommsRepository{}
	template, err := commsRepo.GetTemplateUsingName(templateName)
	if err != nil {
		logrus.Errorf("error in getting templates : %v", err)
		return err
	}

	var templateContent comms.TemplateContent
	if err := json.Unmarshal(template.TemplateContent, &templateContent); err != nil {
		logrus.Errorf("error while unmarshalling : %v", err)
		return err
	}

	body, err := utils.RenderTemplate(templateContent.Body, data)
	if err != nil {
		logrus.Errorf("error in making body of email : %v", err)
		return err
	}

	subject := templateContent.Subject
	message := fmt.Sprintf("To: %s\nSubject: %s\n\n%s", to, subject, body)

	auth := smtp.PlainAuth("", viper.GetString(constants.CommsUsername), viper.GetString(constants.CommsAppPassword), viper.GetString(constants.CommsHostName))
	err = smtp.SendMail(
		viper.GetString(constants.CommsHostName)+":"+viper.GetString(constants.CommsHostPort),
		auth,
		viper.GetString(constants.CommsName),
		[]string{to},
		[]byte(message),
	)

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil

}
