package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-gin/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users/signup", controllers.SignUp())
	router.POST("/users/login", controllers.Login())
	router.POST("/admin/add-product", controllers.ProductViewerAdmin())
	router.POST("/admin/view-product", controllers.SearchProduct())
	router.GET("/users/search", controllers.SearchProductByQuery())
}