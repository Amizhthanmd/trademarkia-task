package controllers

import (
	"net/http"
	"order_inventory_management/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (c *Controller) AddProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if err := c.productService.CreateProduct(&product); err != nil {
		c.logger.Error("Failed to create product :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to create product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "product created successfully"})

}

func (c *Controller) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product

	if err := c.productService.GetProductById(&product, id); err != nil {
		c.logger.Error("Failed to get product :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to get product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Product fetched successfully", "data": product})
}

func (c *Controller) ListProduct(ctx *gin.Context) {
	limit := ctx.DefaultQuery("limit", "10")
	offset := ctx.DefaultQuery("offset", "0")
	searchName := ctx.Query("name")

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

	var products []models.Product
	if searchName != "" {
		if err := c.productService.SearchProducts(&products, searchName); err != nil {
			c.logger.Error("Failed to search products :", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to search products"})
			return
		}
	} else {
		if err := c.productService.ListProducts(&products, limitInt, offsetInt); err != nil {
			c.logger.Error("Failed to list products :", zap.Error(err))
			ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to list products"})
			return
		}
	}
	if len(products) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No products found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Products listed successfully", "data": products})
}

func (c *Controller) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		c.logger.Error("Invalid request payload :", zap.Error(err))
		ctx.JSON(http.StatusBadRequest, gin.H{"status": false, "message": "Invalid request payload"})
		return
	}

	if err := c.productService.UpdateProduct(&product, id); err != nil {
		c.logger.Error("Failed to update product :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to update product"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Product updated successfully"})
}

func (c *Controller) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := c.productService.DeleteProduct(id); err != nil {
		c.logger.Error("Failed to delete product :", zap.Error(err))
		ctx.JSON(http.StatusInternalServerError, gin.H{"status": false, "message": "Failed to delete product"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": true, "message": "Product deleted successfully"})
}
