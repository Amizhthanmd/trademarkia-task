package controllers

import (
	"net/http"
	"order_inventory_management/helpers"
	"order_inventory_management/middleware"
	"order_inventory_management/models"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if !helpers.CheckValidEmail(user.Email) {
		c.logger.Error("Invalid email address")
		ctx.JSON(400, gin.H{"status": false, "message": "Invalid email address"})
		return
	}
	user.Password = helpers.HashPassword(user.Password)

	err := c.UserService.CreateUser(&user)
	if err != nil {
		c.logger.Error("Failed to create user :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Signup successful"})
}

func (c *Controller) Login(ctx *gin.Context) {
	var userLogin models.LoginUser
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if userLogin.Email == "" || userLogin.Password == "" {
		c.logger.Error("Missing credentials")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
		return
	}

	var existingUser models.User
	if err := c.UserService.GetUserByEmail(userLogin.Email, &existingUser); err != nil {
		c.logger.Error("User not found :", zap.Error(err))
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}
	if !helpers.VerifyPassword(existingUser.Password, userLogin.Password) {
		c.logger.Error("Password is incorrect")
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Password is incorrect"})
		return
	}

	jwtToken, err := middleware.GenerateToken(middleware.Claims{
		ID:        existingUser.ID,
		FirstName: existingUser.FirstName,
		LastName:  existingUser.LastName,
		Email:     existingUser.Email,
		Role:      existingUser.Role,
	})
	if err != nil {
		c.logger.Error("Failed to generate token :", zap.Error(err))
		ctx.JSON(http.StatusFailedDependency, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": jwtToken, "message": "login successful", "status": true})
}
