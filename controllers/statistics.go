package controllers

import (
	"fmt"
	"net/http"
	"order_inventory_management/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) GetCustomerStatistics(ctx *gin.Context) {
	emailID := ctx.DefaultQuery("email", "")
	role := ctx.DefaultQuery("role", "")
	fname := ctx.DefaultQuery("fname", "")
	lname := ctx.DefaultQuery("lname", "")
	sortBy := ctx.DefaultQuery("sort_by", "created_at")
	sortOrder := ctx.DefaultQuery("sort_order", "desc")

	var users []models.User

	query := c.UserService.DB

	if emailID != "" {
		query = query.Where("email = ?", emailID)
	}
	if role != "" {
		query = query.Where("role = ?", role)
	}
	if fname != "" {
		query = query.Where("first_name = ?", fname)
	}
	if lname != "" {
		query = query.Where("last_name = ?", lname)
	}

	if sortOrder == "asc" {
		query = query.Order(fmt.Sprintf("%s asc", sortBy))
	} else {
		query = query.Order(fmt.Sprintf("%s desc", sortBy))
	}

	if err := c.UserService.GetUserStats(&users, query); err != nil {
		c.logger.Error("Failed to get user stats", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get user stats"})
		return
	}

	if len(users) == 0 {
		c.logger.Error("Users not found")
		ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Users not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": users})
}

func (c *Controller) GetOrderStatistics(ctx *gin.Context) {
	productID := ctx.DefaultQuery("product_id", "")
	quantity := ctx.DefaultQuery("quantity", "")
	status := ctx.DefaultQuery("status", "")
	sortBy := ctx.DefaultQuery("sort_by", "created_at")
	sortOrder := ctx.DefaultQuery("sort_order", "desc")

	query := c.UserService.DB

	if productID != "" {
		query = query.Where("product_id = ?", productID)
	}
	if quantity != "" {
		query = query.Where("quantity = ?", quantity)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if sortOrder == "asc" {
		query = query.Order(fmt.Sprintf("%s asc", sortBy))
	} else {
		query = query.Order(fmt.Sprintf("%s desc", sortBy))
	}

	var orders []models.Order
	if err := c.UserService.GetOrderStats(&orders, query); err != nil {
		c.logger.Error("Failed to get order stats", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get order stats"})
		return
	}

	if len(orders) == 0 {
		c.logger.Error("Orders not found")
		ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Orders not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": orders})
}

func (c *Controller) GetInventoryStatistics(ctx *gin.Context) {
	productID := ctx.DefaultQuery("product_id", "")
	quantity := ctx.DefaultQuery("quantity", "")
	sortBy := ctx.DefaultQuery("sort_by", "created_at")
	sortOrder := ctx.DefaultQuery("sort_order", "desc")

	query := c.UserService.DB.Model(&models.Product{})

	if productID != "" {
		query = query.Where("products.id = ?", productID)
	}
	if quantity != "" {
		query = query.Where("inventories.quantity = ?", quantity)
	}

	if sortOrder == "asc" {
		query = query.Order(fmt.Sprintf("%s asc", sortBy))
	} else {
		query = query.Order(fmt.Sprintf("%s desc", sortBy))
	}

	var inventoryStats []models.Product
	if err := c.inventoryService.GetInventoryStats(&inventoryStats, query); err != nil {
		c.logger.Error("Failed to get inventory stats", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get inventory stats"})
		return
	}

	if len(inventoryStats) == 0 {
		c.logger.Error("Inventory not found")
		ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Inventory not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "data": inventoryStats})
}
