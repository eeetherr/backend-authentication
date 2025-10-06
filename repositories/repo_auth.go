package repositories

import (
	"ankit/authentication/database"
	"ankit/authentication/dto/auth"
)

type AuthRepository struct{}

func (r *AuthRepository) CreateAuth(auth *auth.Auth) error {
	return database.DB.Create(auth).Error
}

func (r *AuthRepository) GetAuthByEmail(email string) (*auth.Auth, error) {
	var auth auth.Auth
	result := database.DB.Where("email = ?", email).First(&auth)
	if result.Error != nil {
		return nil, result.Error
	}
	return &auth, nil
}
