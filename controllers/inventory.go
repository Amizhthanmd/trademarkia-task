package controllers

import (
	"net/http"
	"order_inventory_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) AddInventory(ctx *gin.Context) {
	var inventory models.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if err := c.inventoryService.CreateInventory(&inventory); err != nil {
		c.logger.Error("Failed to create inventory :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to create inventory"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Inventory created successfully"})
}

func (c *Controller) ListInventory(ctx *gin.Context) {
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

	var inventory []models.Inventory
	if err := c.inventoryService.ListInventory(&inventory, limitInt, offsetInt); err != nil {
		c.logger.Error("Failed to list inventory :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to list inventory"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Inventory listed successfully", "data": inventory})
}

func (c *Controller) UpdateInventory(ctx *gin.Context) {
	id := ctx.Param("id")
	var inventory models.Inventory
	if err := ctx.ShouldBindJSON(&inventory); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if err := c.inventoryService.UpdateInventory(&inventory, id); err != nil {
		c.logger.Error("Failed to update inventory :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to update inventory"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Inventory updated successfully", "data": inventory})
}

func (c *Controller) DeleteInventory(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.inventoryService.DeleteInventory(id); err != nil {
		c.logger.Error("Failed to delete inventory :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to delete inventory"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Inventory deleted successfully"})
}
