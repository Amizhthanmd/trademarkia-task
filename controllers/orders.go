package controllers

import (
	"net/http"
	"order_inventory_management/models"

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
		c.logger.Error("Insufficient quantity")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Insufficient quantity"})
		return
	}

	totalAmount := float64(order.Quantity) * product.Price
	orderDetails := models.Order{
		TotalAmount: totalAmount,
		Status:      "Order Placed",
		Quantity:    order.Quantity,
		UserID:      ctx.GetString("user_id"),
		Products:    []models.Product{product},
	}

	if err := c.UserService.PlaceOrder(&orderDetails); err != nil {
		c.logger.Error("Failed to create order :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to place order"})
		return
	}

	product.Inventory.Quantity -= orderDetails.Quantity

	if err := c.inventoryService.UpdateInventoryQty(product.Inventory.Quantity, product.InventoryID); err != nil {
		c.logger.Error("Failed to update inventory quantity :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to update inventory quantity"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Order placed successfully", "data": orderDetails})
}

func (c *Controller) GetOrder(ctx *gin.Context) {}

func (c *Controller) ListOrder(ctx *gin.Context) {}
