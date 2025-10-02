package repositories

import (
	//"ankit/authentication/dto"
	"ankit/authentication/dto/comms"
	"errors"
	"time"

	"gorm.io/gorm"
)

type CommsRepository struct {
	db *gorm.DB
}

func NewCommsRepository(db *gorm.DB) *CommsRepository {
	return &CommsRepository{db: db}
}

// Create a new verification record
func (r *CommsRepository) CreateVerification(userID int64, email, codeHash string, expiry time.Time) error {
	record := comms.Comms{
		UserID:    userID,
		Email:     email,
		CodeHash:  codeHash,
		ExpiresAt: expiry,
	}
	return r.db.Create(&record).Error
}

// Get latest verification for a user + channel
func (r *CommsRepository) GetLatestVerification(userID int64, channel string) (*comms.Comms, error) {
	var record comms.Comms
	err := r.db.Where("user_id = ? AND channel = ?", userID, channel).
		Order("created_at desc").
		First(&record).Error
	if err != nil {
		return nil, err
	}
	return &record, nil
}

func (r *CommsRepository) SaveVerificationCode(v comms.CommsVerification) error {
	result := r.db.Create(&v)
	return result.Error
}

func (r *CommsRepository) GetLatestVerificationByEmail(email string) (*comms.CommsVerification, error) {
	var verification comms.CommsVerification

	err := r.db.
		Where("email = ?", email).
		Order("created_at desc").
		First(&verification).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // No record found
		}
		return nil, err
	}

	return &verification, nil
}

func (r *CommsRepository) MarkCodeAsUsed(id uint) error {
	return r.db.Model(&comms.CommsVerification{}).
		Where("id = ?", id).
		Update("is_used", true).Error
}
