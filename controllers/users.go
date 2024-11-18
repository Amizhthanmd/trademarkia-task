package controllers

import (
	"net/http"
	"order_inventory_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) ListUser(ctx *gin.Context) {
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

	var users []models.User

	if err := c.UserService.ListUsers(&users, limitInt, offsetInt); err != nil {
		c.logger.Error("Failed to list users :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to list users"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Users fetched successfully", "data": users})
}

func (c *Controller) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	if err := c.UserService.GetUserById(&user, id); err != nil {
		c.logger.Error("Failed to get user :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get user"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "User fetched successfully", "data": user})
}
