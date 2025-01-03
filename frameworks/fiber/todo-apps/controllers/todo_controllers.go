package controllers

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-restapi/todos/config"
	"github.com/golang-restapi/todos/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)
var db *mongo.Collection


func GetTodos (ctx *fiber.Ctx) error {
	var todos []models.Todo

	if err:= config.ConnectDB(); err != nil {
		log.Fatal("Error connecting to database", err)
	}

	cursor, err := db.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo models.Todo
		if err := cursor.Decode(&todo); err != nil {
			return err

		}
		todos = append(todos, todo)
	}

	return ctx.JSON(todos)
}