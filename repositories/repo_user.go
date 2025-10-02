package repositories

import (
	"ankit/authentication/database"
	"ankit/authentication/dto/users"
	"errors"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	db := database.GetDB()
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *users.User) error {
	result := r.DB.Create(user)
	return result.Error
}

func (r *UserRepository) GetUserByEmailAndPhoneNumber(email, phoneNumber string) (*users.User, error) {
	var user users.User
	result := r.DB.Where("email = ? OR phone_number = ?", email, phoneNumber).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}
	return &user, nil
}
