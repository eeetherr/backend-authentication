package users

type User struct {
	ID          int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Email       string `gorm:"uniqueIndex;not null" json:"email"`
	Name        string `gorm:"not null" json:"name"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
}
