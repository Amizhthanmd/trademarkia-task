package controllers

import (
	"order_inventory_management/services"

	"gorm.io/gorm"
)

type Controller struct {
	UserDB      *gorm.DB
	UserService *services.UserService
}

func NewController(
	UserDB *gorm.DB,
	UserService *services.UserService,
) *Controller {
	return &Controller{
		UserDB:      UserDB,
		UserService: UserService,
	}
}
