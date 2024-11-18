package routes

import (
	"log"
	"order_inventory_management/controllers"
	"order_inventory_management/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	admin      = []string{"admin"}
	user       = []string{"user"}
	admin_user = []string{"admin", "user"}
)

func StartRouter(controller *controllers.Controller, PORT string, GIN_MODE string) {
	gin.SetMode(GIN_MODE)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
	}))

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": false, "message": "Route not found"})
		c.Abort()
	})

	v1 := router.Group("api/v1")
	{
		v1.POST("signup", controller.SignUp)
		v1.POST("login", controller.Login)
		UserRoutes(v1, controller)       // User routes
		ProductRoutes(v1, controller)    // Products routes
		InventoryRoutes(v1, controller)  // Inventory routes
		OrderRoutes(v1, controller)      // Orders routes
		StatisticsRoutes(v1, controller) // Statistics routes
	}

	if err := router.Run(PORT); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func UserRoutes(v1 *gin.RouterGroup, controller *controllers.Controller) {
	userRoute := v1.Group("users")
	{
		userRoute.GET(":id", middleware.AuthMiddleware(admin), controller.GetUser)
		userRoute.GET("", middleware.AuthMiddleware(admin), controller.ListUser)
	}
}

func ProductRoutes(v1 *gin.RouterGroup, controller *controllers.Controller) {
	productRoute := v1.Group("products")
	{
		productRoute.GET(":id", middleware.AuthMiddleware(admin_user), controller.GetProduct)
		productRoute.GET("", middleware.AuthMiddleware(admin_user), controller.ListProduct)

		productRoute.POST("", middleware.AuthMiddleware(admin), controller.AddProduct)
		productRoute.PUT(":id", middleware.AuthMiddleware(admin), controller.UpdateProduct)
		productRoute.DELETE(":id", middleware.AuthMiddleware(admin), controller.DeleteProduct)
	}
}

func OrderRoutes(v1 *gin.RouterGroup, controller *controllers.Controller) {
	orderRoute := v1.Group("orders")
	{
		orderRoute.POST("", middleware.AuthMiddleware(admin_user), controller.PlaceOrder)
		orderRoute.GET(":user_id", middleware.AuthMiddleware(admin_user), controller.GetOrder)
		orderRoute.GET("", middleware.AuthMiddleware(admin_user), controller.ListOrder)
	}
}

func InventoryRoutes(v1 *gin.RouterGroup, controller *controllers.Controller) {
	inventoryRoute := v1.Group("inventory")
	{
		inventoryRoute.POST("", middleware.AuthMiddleware(admin), controller.AddInventory)
		inventoryRoute.PUT(":id", middleware.AuthMiddleware(admin), controller.UpdateInventory)
		inventoryRoute.DELETE(":id", middleware.AuthMiddleware(admin), controller.DeleteInventory)
		inventoryRoute.GET("", middleware.AuthMiddleware(admin), controller.ListInventory)
	}
}

func StatisticsRoutes(v1 *gin.RouterGroup, controller *controllers.Controller) {
	statsRoute := v1.Group("stats")
	{
		statsRoute.GET("customers", middleware.AuthMiddleware(admin), controller.GetCustomerStatistics)
		statsRoute.GET("orders", middleware.AuthMiddleware(admin), controller.GetOrderStatistics)
		statsRoute.GET("inventory", middleware.AuthMiddleware(admin), controller.GetInventoryStatistics)
	}
}
