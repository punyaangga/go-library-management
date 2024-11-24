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
		protected.GET("/profile", func(c *gin.Context) {
			username, _ := c.Get("username")
			c.JSON(200, gin.H{"message": "Hello, " + username.(string)})
		})
		protected.POST("/logout", controllers.Logout)
	}
}
