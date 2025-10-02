package services

import (
	"ankit/authentication/constants"
	"ankit/authentication/dto"
	"ankit/authentication/dto/comms"
	"ankit/authentication/dto/users"
	"ankit/authentication/repositories"
	"ankit/authentication/utils"
	"errors"
	"time"
)

type AuthService struct {
	UserRepo  *repositories.UserRepository
	CommsRepo *repositories.CommsRepository
	Comms     *commsService // for sending email
}

// SignUp handles user registration and sends OTP
func (s *AuthService) SignUp(req dto.SignUpRequest) error {
	// Check if user already exists
	existingUser, _ := s.UserRepo.GetUserByEmailAndPhoneNumber(req.Email, req.PhoneNumber)
	if existingUser != nil {
		return errors.New("user already exists with this email or phone number")
	}

	// Generate OTP
	verificationCode := utils.GenerateVerificationCode(constants.VerificationCodeLength)

	// Send email with OTP
	err := s.Comms.SendEmail(req.Email, verificationCode)
	if err != nil {
		return errors.New("could not send verification email")
	}

	// Save user (optional: you can delay until after verification)
	user := users.User{
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password, // consider hashing this
	}
	err = s.UserRepo.CreateUser(&user)
	if err != nil {
		return errors.New("could not create user")
	}

	// Save OTP in comms table
	otp := comms.CommsVerification{
		Email:            req.Email,
		VerificationCode: verificationCode,
		IsUsed:           false,
		CreatedAt:        time.Now(),
	}
	err = s.CommsRepo.SaveVerificationCode(otp)
	if err != nil {
		return errors.New("could not save OTP")
	}

	return nil
}

// VerifyEmail handles OTP verification
func (s *AuthService) VerifyEmail(email string, otpInput string) error {
	record, err := s.CommsRepo.GetLatestVerificationByEmail(email)
	if err != nil {
		return errors.New("database error")
	}
	if record == nil {
		return errors.New("no OTP found for this email")
	}
	if record.IsUsed {
		return errors.New("OTP already used")
	}
	if record.VerificationCode != otpInput {
		return errors.New("invalid OTP")
	}

	// Mark OTP as used
	err = s.CommsRepo.MarkCodeAsUsed(record.ID)
	if err != nil {
		return errors.New("failed to update OTP status")
	}

	return nil
}
