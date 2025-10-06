package auth

import "time"

type SignUpRequest struct {
	Email       string `json:"email" binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	Name        string `json:"name" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

type OtpRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required"`
}

type Auth struct {
	ID             uint   `gorm:"primaryKey;autoIncrement"`
	Email          string `gorm:"type:varchar(100);unique;not null"`
	HashedPassword string `gorm:"type:varchar(255);not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
	Email string `json:"email"`
}
