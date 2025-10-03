package repositories

import (
	"ankit/authentication/database"
	"github.com/sirupsen/logrus"

	//"ankit/authentication/dto"
	"ankit/authentication/dto/comms"
)

type CommsRepository struct {
}

func NewCommsRepository() *CommsRepository {
	return &CommsRepository{}
}

func (r *CommsRepository) GetTemplateUsingName(templateName string) (comms.Template, error) {
	db := database.GetDB()
	var template comms.Template
	err := db.Model(&comms.Template{}).Where("template_name = ?", templateName).First(&template).Error
	if err != nil {
		logrus.Errorf("Error in getting the template %v", err)
		return template, err
	}
	return template, nil
}

func (r *CommsRepository) SaveCommunicationLogs(commsLogs *comms.CommunicationLog) error {
	db := database.GetDB()
	err := db.Create(&commsLogs).Error
	if err != nil {
		logrus.Errorf("Error in getting the template %v", err)
		return err
	}
	return err
}
