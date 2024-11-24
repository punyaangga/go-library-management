package routes

import (
	"libraryManagement/controllers"
	"libraryManagement/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Public routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected routes
	protected := router.Group("/api")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.POST("/stock", controllers.AddStock)
		protected.PUT("/stock/:id", controllers.UpdateStock)
		protected.GET("/stock", controllers.GetStocks)

		protected.POST("/category", controllers.AddCategory)
		protected.PUT("/category/:id", controllers.UpdateCategory)

		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(200, gin.H{"message": "Hello, " + username.(string)})
		})

		protected.POST("/logout", controllers.Logout)
	}
}
