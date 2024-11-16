package main

import (
	"fmt"
	"order_inventory_management/appInit"
	"order_inventory_management/controllers"
	"order_inventory_management/db"
	"order_inventory_management/helpers"
	"order_inventory_management/routes"
	"order_inventory_management/services"
	"os"
)

func main() {
	// Load ENV
	helpers.LoadEnv()

	var (
		PORT            = fmt.Sprintf(":%s", os.Getenv("PORT"))
		GIN_MODE        = os.Getenv("GIN_MODE")
		POSTGRES_DB_URL = os.Getenv("POSTGRES_DB_URL")
		USER_DB         = os.Getenv("DB")
	)

	// Initialize logger
	logger := appInit.ZapLogger(GIN_MODE)
	defer logger.Sync()

	// Intialize user database connection
	userDB := db.InitializeUserDB(POSTGRES_DB_URL, USER_DB)

	// Intialize user services
	userService := services.InitializeUserService(userDB, logger)

	productService := services.InitializeProductService(userDB, logger)
	InitializeInventoryService := services.InitializeInventoryService(userDB, logger)

	// Intialize Controller
	controller := controllers.NewController(userDB, userService, productService, InitializeInventoryService, logger)

	// Initialize Routes
	routes.StartRouter(controller, PORT, GIN_MODE)

}
