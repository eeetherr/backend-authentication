package comms

import "time"

type Comms struct {
	ID        int64     `gorm:"primary_key;auto_increment" json:"id"`
	UserID    int64     `gorm:"not null;index" json:"user_id"`
	Email     string    `gorm:"not null" json:"email"`
	CodeHash  string    `gorm:"not null" json:"-"` // store hashed OTP
	ExpiresAt time.Time `gorm:"not null" json:"expires_at"`
	Attempts  int       `gorm:"default:0" json:"attempts"` // number of failed tries
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type EmailData struct {
	Name string
	Code string
}

type CommsVerification struct {
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	Email            string    `gorm:"index;not null"`
	VerificationCode string    `gorm:"not null"`
	CreatedAt        time.Time `gorm:"autoCreateTime"`
	IsUsed           bool      `gorm:"default:false"`
}

type SignUpRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

type VerifyRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required"`
}

//type CommsService struct {
//	SMTPHost string
//	SMTPPort string
//	Username string
//	Password string
//	From     string
//}
