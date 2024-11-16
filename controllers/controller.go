package controllers

import (
	"order_inventory_management/services"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Controller struct {
	UserDB           *gorm.DB
	UserService      *services.UserService
	productService   *services.ProductService
	inventoryService *services.InventoryService
	logger           *zap.Logger
}

func NewController(
	UserDB *gorm.DB,
	UserService *services.UserService,
	productService *services.ProductService,
	inventoryService *services.InventoryService,
	logger *zap.Logger,
) *Controller {
	return &Controller{
		UserDB:           UserDB,
		UserService:      UserService,
		productService:   productService,
		inventoryService: inventoryService,
		logger:           logger,
	}
}
