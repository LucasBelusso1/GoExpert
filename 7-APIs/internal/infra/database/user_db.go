package database

import (
	"github.com/LucasBelusso1/GoExpert/7-APIS/internal/entity"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func (u *User) Create(user *entity.User) error {
	return u.DB.Create(user).Error
}

func (u *User) FindByEmail(email string) (*entity.User, error) {
	var user entity.User

	err := u.DB.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func NewUser(db *gorm.DB) *User {
	return &User{DB: db}
}
