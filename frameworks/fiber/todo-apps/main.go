package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-restapi/todos/config"
	"github.com/golang-restapi/todos/controllers"
)

func main() {
	
	if err:= config.ConnectDB() ; err ==nil{

		fmt.Println("Connected to MONGODB ATLAS")
	}


	app := fiber.New()

	app.Get("/", controllers.GetTodos)

	log.Fatal(app.Listen(":5000"))
	
}
