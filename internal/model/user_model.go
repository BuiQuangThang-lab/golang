package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id          string    `json:"id"`
	UserName    string    `json:"user_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
	FullName    string    `json:"full_name"`
	Role        int       `json:"role"`
	Gender      int       `json:"gender"`
	Birthday    string    `json:"birthday"`
	IsDelete    int       `json:"is_delete"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"created_date"`
}

func (User) TableName() string { return "qlnv.user" }

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return
	}
}
