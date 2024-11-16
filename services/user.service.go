package services

import (
	"order_inventory_management/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserService struct {
	DB     *gorm.DB
	logger *zap.Logger
}

func InitializeUserService(UserDB *gorm.DB, logger *zap.Logger) *UserService {
	return &UserService{
		DB:     UserDB,
		logger: logger,
	}
}

func (u *UserService) CreateUser(data *models.User) error {
	u.logger.Info("Create user")
	return u.DB.Create(&data).Error
}

func (u *UserService) GetUserByEmail(email string, data *models.User) error {
	u.logger.Info("Get user by email")
	return u.DB.Where("email = ?", email).First(&data).Error
}
