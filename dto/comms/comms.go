package comms

import (
	"encoding/json"
	"time"
)

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

type Template struct {
	TemplateID      int             `gorm:"column:template_id;primaryKey;not null"`
	TemplateName    string          `gorm:"column:template_name;size:50;not null"`
	TemplateContent json.RawMessage `gorm:"column:template_content"` // jsonb
	CreatedAt       *time.Time      `gorm:"column:created_at"`
	UpdatedAt       *time.Time      `gorm:"column:updated_at"`
	DeletedAt       *time.Time      `gorm:"column:deleted_at"` // soft delete (optional GORM's soft delete)
}

// TableName sets the name of the table in the database
func (Template) TableName() string {
	return "templates"
}

type TemplateContent struct {
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

type CommunicationLog struct {
	ID               int             `json:"id" gorm:"primaryKey;autoIncrement;column:id"`
	EventType        string          `json:"event_type" gorm:"column:event_type;type:varchar(50);not null"`
	Status           string          `json:"status" gorm:"column:status;type:varchar(50);not null"`
	AdditionalData   json.RawMessage `json:"additional_data,omitempty" gorm:"column:additional_data;type:jsonb"`
	Destination      string          `json:"destination" gorm:"column:destination;type:varchar(50);not null"`
	CreatedAt        *time.Time      `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt        *time.Time      `json:"updated_at,omitempty" gorm:"column:updated_at"`
	DeletedAt        *time.Time      `json:"deleted_at,omitempty" gorm:"column:deleted_at"`
	VerificationCode string          `json:"verification_code" gorm:"column:verification_code;type:varchar(50);not null"`
}

// TableName overrides default GORM table name
func (CommunicationLog) TableName() string {
	return "communication_logs"
}
