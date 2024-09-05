package routes

import (
	"accounta-go/controllers"
	"accounta-go/middleware"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(router *gin.RouterGroup) {
	routes := router.Group("/category")
	routes.Use(middleware.IsAuthenticated)
	routes.GET("/list", controllers.GetCategories)
	routes.GET("/:id", controllers.GetCategory)
	routes.POST("/", controllers.PostCategory)
	routes.PATCH("/:id", controllers.PatchCategory)
	routes.DELETE("/:id", controllers.DeleteCategory)

}
