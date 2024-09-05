package routes

import "github.com/gin-gonic/gin"

func ApiRoutes(r *gin.Engine) {
	api := r.Group("/api")
	UserRoutes(api)
}
