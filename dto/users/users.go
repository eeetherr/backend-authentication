package users

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Email    string `gorm:"uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"-"`  // Hidden from JSON responses
	Name     string `gorm:"not null" json:"name"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
}