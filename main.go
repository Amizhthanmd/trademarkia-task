package main

import (
	"fmt"
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

	// Intialize user database connection
	userDB := db.InitializeUserDB(POSTGRES_DB_URL, USER_DB)

	// Intialize user services
	userService := services.InitializeUserService(userDB)

	// Intialize Controller
	controller := controllers.NewController(userDB, userService)

	// Initialize Routes
	routes.StartRouter(controller, PORT, GIN_MODE)

}
