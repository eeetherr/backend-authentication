package repositories

import (
	"ankit/authentication/database"
	"ankit/authentication/dto/users"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByEmailAndPhoneNumber(email, phoneNumber string) (*users.User, error) {
	db := database.GetDB()
	var user users.User
	result := db.Where("email = ? OR phone_number = ?", email, phoneNumber).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
