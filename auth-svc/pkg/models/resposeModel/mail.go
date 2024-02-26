package resposemodel

type MailUserDetils struct {
	Email  string `gorm:"column:email"`
	UserID string `gorm:"column:id"`
}
