package controllers

import (
	"net/http"
	"order_inventory_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) PlaceOrder(ctx *gin.Context) {
	var order models.OrderDetails
	if err := ctx.ShouldBindJSON(&order); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	var product models.Product
	if err := c.productService.GetProductById(&product, order.ProductID); err != nil {
		c.logger.Error("Failed to get product :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get product"})
		return
	}

	if product.Inventory.Quantity < order.Quantity {
		c.logger.Error("Stocks not available")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Stocks not available"})
		return
	}

	totalAmount := float64(order.Quantity) * product.Price
	orderDetails := models.Order{
		TotalAmount: totalAmount,
		Status:      "Order Placed",
		Quantity:    order.Quantity,
		UserID:      ctx.GetString("user_id"),
		ProductID:   order.ProductID,
		Products:    []models.Product{product},
	}

	if err := c.UserService.PlaceOrder(&orderDetails); err != nil {
		c.logger.Error("Failed to create order :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to place order"})
		return
	}

	// product.Inventory.Quantity -= orderDetails.Quantity

	// if err := c.inventoryService.UpdateInventoryQty(product.Inventory.Quantity, product.InventoryID); err != nil {
	// 	c.logger.Error("Failed to update inventory quantity :", zap.Error(err))
	// 	ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to update inventory quantity"})
	// 	return
	// }

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Order placed successfully", "data": orderDetails})
}

func (c *Controller) GetOrder(ctx *gin.Context) {
	var orders []models.Order
	id := ctx.Param("id")
	if err := c.UserService.GetOrderById(&orders, id); err != nil {
		c.logger.Error("Failed to get order :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get order"})
		return
	}

	if len(orders) == 0 {
		c.logger.Error("Orders not found")
		ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Orders not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Order fetched successfully", "data": orders})
}

func (c *Controller) ListOrder(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.logger.Error("Invalid limit :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid limit"})
		return
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		c.logger.Error("Invalid offset :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid offset"})
		return
	}

	var orders []models.Order
	if err := c.UserService.GetOrders(&orders, limitInt, offsetInt); err != nil {
		c.logger.Error("Failed to list orders :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to list orders"})
		return
	}

	if len(orders) == 0 {
		c.logger.Error("Orders not found")
		ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "Orders not found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Orders listed successfully", "data": orders})
}
