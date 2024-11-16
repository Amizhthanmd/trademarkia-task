package services

import (
	"order_inventory_management/models"

	"gorm.io/gorm"
)

type UserService struct {
	DB *gorm.DB
}

func InitializeUserService(UserDB *gorm.DB) *UserService {
	return &UserService{
		DB: UserDB,
	}
}

func (u *UserService) CreateUser(data *models.User) error {
	return u.DB.Create(&data).Error
}

func (u *UserService) GetUserByEmail(email string, data *models.User) error {
	return u.DB.Where("email = ?", email).First(&data).Error
}
