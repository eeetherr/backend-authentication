package services

import (
	"ankit/authentication/constants"
	"ankit/authentication/dto/auth"
	"ankit/authentication/dto/comms"
	"ankit/authentication/repositories"
	"ankit/authentication/utils"
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
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
		EventType:      constants.SignUpComms,
		Status:         status,
		Destination:    req.Email,
		AdditionalData: additionalData,
	}
	errComms := commsRepo.SaveCommunicationLogs(commsData)
	if errComms != nil {
		logrus.Errorf("Unable to save comms")
		return err
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
