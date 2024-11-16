package services

import (
	"fmt"
	"order_inventory_management/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type InventoryService struct {
	DB     *gorm.DB
	logger *zap.Logger
}

func InitializeInventoryService(UserDB *gorm.DB, logger *zap.Logger) *InventoryService {
	return &InventoryService{
		DB:     UserDB,
		logger: logger,
	}
}

func (i *InventoryService) CreateInventory(data *models.Inventory) error {
	i.logger.Info("Create inventory")
	return i.DB.Create(&data).Error
}

func (i *InventoryService) ListInventory(data *[]models.Inventory, limit, offset int) error {
	i.logger.Info("List inventory")
	return i.DB.Limit(limit).Offset(offset).Find(data).Error
}

func (i *InventoryService) UpdateInventory(data *models.Inventory, id string) error {
	var inventory models.Inventory
	i.logger.Info("Update inventory")
	if err := i.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		i.logger.Error("Inventory not found", zap.Error(err))
		return fmt.Errorf("Inventory not found: %v", err)
	}
	return i.DB.Where("id = ?", id).Updates(&data).Error
}

func (i *InventoryService) UpdateInventoryQty(data int, id string) error {
	return i.DB.Model(&models.Inventory{}).Where("id = ?", id).Update("quantity", data).Error
}

func (i *InventoryService) DeleteInventory(id string) error {
	i.logger.Info("Delete inventory")
	return i.DB.Where("id = ?", id).Delete(&models.Inventory{}).Error
}
