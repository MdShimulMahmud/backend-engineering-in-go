package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-gin/controllers"
	"github.com/go-gin/database"
	"github.com/go-gin/middleware"
	"github.com/go-gin/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == ""{
		port=8000
	}

	app := controllers.NewApplication(database.ProductData(database.Client, "Products"),database.UserData(database.Client, "Users") )


	router:= gin.New()

	router.Use(gin.Logger())

	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	router.GET("/add-to-cart", app.AddToCart())
	router.GET("/checkout", app.Checkout())
	router.GET("/remove-item", app.RemoveItem())
	router.GET("/instant-buy", app.InstantBuy())

	log.Fatal(router.Run(":"+port))
}
