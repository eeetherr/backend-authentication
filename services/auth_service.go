package services

import (
	"ankit/authentication/constants"
	"ankit/authentication/dto/auth"
	"ankit/authentication/dto/comms"
	"ankit/authentication/dto/users"
	"ankit/authentication/repositories"
	"ankit/authentication/utils"
	"encoding/json"
	"errors"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	UserRepo  *repositories.UserRepository
	CommsRepo *repositories.CommsRepository
	Comms     *commsService // for sending email
}

// SignUp handles user registration and sends OTP
func (s *AuthService) SignUp(req auth.SignUpRequest) error {
	// Check if user already exists
	existingUser, _ := s.UserRepo.GetUserByEmailAndPhoneNumber(req.Email, req.PhoneNumber)
	if existingUser != nil {
		return errors.New("user already exists with this email or phone number")
	}

	// Generate OTP
	verificationCode := utils.GenerateVerificationCode(constants.VerificationCodeLength)

	data := map[string]string{
		"Name": req.Name,
		"Code": verificationCode,
	}

	var status string
	// Send email with OTP
	err := s.Comms.SendEmail(req.Email, data, constants.SignUpComms)
	if err != nil {
		status = "FAILED"
		logrus.Errorf("Failed to send the comms while signup")
	} else {
		status = "SUCCESS"
	}

	additionalData, err := json.Marshal(req)
	if err != nil {
		logrus.Errorf("Unable to unmarshal the data")
		return err
	}
	commsRepo := repositories.CommsRepository{}
	commsData := &comms.CommunicationLog{
		EventType:        constants.SignUpComms,
		Status:           status,
		Destination:      req.Email,
		AdditionalData:   additionalData,
		VerificationCode: verificationCode,
	}
	errComms := commsRepo.SaveCommunicationLogs(commsData)
	if errComms != nil {
		logrus.Errorf("Unable to save comms")
		return errComms
	}

	return err
}

//
//// VerifyEmail handles OTP verification
//func (s *AuthService) VerifyEmail(email string, otpInput string) error {
//	record, err := s.CommsRepo.GetLatestVerificationByEmail(email)
//	if err != nil {
//		return errors.New("database error")
//	}
//	if record == nil {
//		return errors.New("no OTP found for this email")
//	}
//	if record.IsUsed {
//		return errors.New("OTP already used")
//	}
//	if record.VerificationCode != otpInput {
//		return errors.New("invalid OTP")
//	}
//
//	// Mark OTP as used
//	err = s.CommsRepo.MarkCodeAsUsed(record.ID)
//	if err != nil {
//		return errors.New("failed to update OTP status")
//	}
//
//	return nil
//}

func (s *AuthService) VerifyAuth(req auth.OtpRequest) error {
	commsRepo := repositories.CommsRepository{}
	commsData, err := commsRepo.GetUserByEventTypeAndEmail(req.Email, constants.SignUpComms)
	if err != nil || commsData == nil {
		logrus.Errorf("Unable to find the user with verification code")
		return errors.New("invalid verification code")
	}

	if commsData.VerificationCode != req.Code {
		logrus.Errorf("verification failed")
		return errors.New("invalid verification code")
	}

	var signUpReq auth.SignUpRequest
	err = json.Unmarshal([]byte(commsData.AdditionalData), &signUpReq)
	if err != nil {
		logrus.Errorf("Failed to parse additional data: %v", err)
		return errors.New("internal server error")
	}

	hashedPassword, err := utils.HashPassword(signUpReq.Password)
	if err != nil {
		logrus.Error("Password hashing failed")
		return errors.New("internal server error")
	}

	user := &users.User{
		Name:        signUpReq.Name,
		PhoneNumber: signUpReq.PhoneNumber,
		Email:       signUpReq.Email,
	}

	userRepo := repositories.UserRepository{}
	err = userRepo.SaveUser(user)
	if err != nil {
		logrus.Error("Unable to save the user")
		return errors.New("internal server error")
	}

	auth := &auth.Auth{
		Email:          signUpReq.Email,
		HashedPassword: hashedPassword,
	}
	authRepo := repositories.AuthRepository{}

	err = authRepo.CreateAuth(auth)
	if err != nil {
		logrus.Error("Unable to save the auth")
		return errors.New("internal server error")
	}

	return nil
}

func (s *AuthService) Login(req auth.LoginRequest) (auth.LoginResponse, error) {
	authRepo := repositories.AuthRepository{}
	authData, err := authRepo.GetAuthByEmail(req.Email)
	if err != nil || authData == nil {
		return auth.LoginResponse{}, errors.New("invalid email or password")
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(authData.HashedPassword), []byte(req.Password))
	if err != nil {
		return auth.LoginResponse{}, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := utils.GenerateJWT(authData.Email)
	if err != nil {
		return auth.LoginResponse{}, errors.New("failed to generate token")
	}

	resToken := auth.LoginResponse{
		Token: token,
		Email: req.Email,
	}

	return resToken, nil
}
