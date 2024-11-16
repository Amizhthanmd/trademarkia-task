package controllers

import (
	"net/http"
	"order_inventory_management/helpers"
	"order_inventory_management/middleware"
	"order_inventory_management/models"

	"github.com/gin-gonic/gin"
)

func (c *Controller) SignUp(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if !helpers.CheckValidEmail(user.Email) {
		ctx.JSON(400, gin.H{"status": false, "message": "Invalid email address"})
		return
	}
	user.Password = helpers.HashPassword(user.Password)

	err := c.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "User created successfully"})
}

func (c *Controller) Login(ctx *gin.Context) {
	var userLogin models.LoginUser
	if err := ctx.ShouldBindJSON(&userLogin); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if userLogin.Email == "" || userLogin.Password == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Missing credentials"})
		return
	}

	var existingUser models.User
	if err := c.UserService.GetUserByEmail(userLogin.Email, &existingUser); err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	if !helpers.VerifyPassword(existingUser.Password, userLogin.Password) {
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
		ctx.JSON(http.StatusFailedDependency, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": jwtToken, "message": "login successful", "status": true})
}
