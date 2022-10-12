package entity

type User struct {
	Id       string `gorm:"primaryKey,autoIncrement"`
	Name     string `gorm:"size:256"`
	Email    string `gorm:"size:256"`
	Password string `gorm:"size:256"`
}
