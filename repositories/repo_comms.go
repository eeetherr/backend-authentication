package repositories

import (
	"ankit/authentication/database"

	"ankit/authentication/dto/comms"

	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

var DB = &gorm.DB{}

type CommsRepository struct {
}

func NewCommsRepository() *CommsRepository {
	return &CommsRepository{}
}

func (r *CommsRepository) GetTemplateUsingName(templateName string) (comms.Template, error) {
	db := database.GetDB()
	var template comms.Template
	db = db.Debug()
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

func (r *CommsRepository) GetUserByEventTypeAndEmail(email, eventType string) (*comms.CommunicationLog, error) {
	var log comms.CommunicationLog
	err := database.DB.Where("destination = ? AND event_type = ?", email, eventType).
		Order("created_at desc").
		First(&log).Error

	if err != nil {
		return nil, err
	}
	return &log, nil
}
