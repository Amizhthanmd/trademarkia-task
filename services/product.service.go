package services

import (
	"fmt"
	"order_inventory_management/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type ProductService struct {
	DB     *gorm.DB
	logger *zap.Logger
}

func InitializeProductService(UserDB *gorm.DB, logger *zap.Logger) *ProductService {
	return &ProductService{
		DB:     UserDB,
		logger: logger,
	}
}

func (p *ProductService) CreateProduct(data *models.Product) error {
	p.logger.Info("Create product")
	return p.DB.Create(&data).Error
}

func (p *ProductService) GetProductById(data *models.Product, id string) error {
	p.logger.Info("Get product by id")
	return p.DB.Where("id = ?", id).First(data).Error
}

func (p *ProductService) ListProducts(data *[]models.Product, limit, offset int) error {
	p.logger.Info("List products")
	return p.DB.Limit(limit).Offset(offset).Find(data).Error
}

func (p *ProductService) UpdateProduct(data *models.Product, id string) error {
	var product models.Product
	p.logger.Info("Update product")
	if err := p.DB.Where("id = ?", id).First(&product).Error; err != nil {
		p.logger.Error("Product not found", zap.Error(err))
		return fmt.Errorf("Product not found: %v", err)
	}
	return p.DB.Where("id = ?", id).Updates(&data).Error
}

func (p *ProductService) DeleteProduct(id string) error {
	p.logger.Info("Delete product")
	return p.DB.Where("id = ?", id).Delete(&models.Product{}).Error
}
