package routes

import (
	"accounta-go/controllers"
	"accounta-go/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	routes := router.Group("/user")
	routes.POST("/register", controllers.UserRegister)
	routes.POST("/login", controllers.UserLogin)
	routes.GET("/", middleware.IsAuthenticated, controllers.UserProfile)
}
