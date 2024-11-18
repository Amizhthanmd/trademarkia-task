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

func (u *UserService) PlaceOrder(data *models.Order) error {
	u.logger.Info("Place order")
	return u.DB.Create(&data).Error
}

func (u *UserService) GetOrderByUserId(data *[]models.Order, id string) error {
	u.logger.Info("Get order by id")
	return u.DB.Find(&data, "user_id = ?", id).Error
}

func (u *UserService) GetOrders(data *[]models.Order, query *gorm.DB, limit, offset int) error {
	u.logger.Info("List orders with filters")
	query = query.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("users.id, users.first_name, users.last_name, users.email")
	}).Preload("Products", func(db *gorm.DB) *gorm.DB {
		return db.Select("products.id, products.name, products.price, products.description")
	})
	return query.Limit(limit).Offset(offset).Find(&data).Error
}
