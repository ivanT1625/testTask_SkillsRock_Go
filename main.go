package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"todo-api/db"
	"todo-api/handlers"
)

func main() {
	db.InitDB("postgres://postgres:1Qwerty@localhost:5432/todo_db")

	app := fiber.New()

	api := app.Group("/api/tasks")
	api.Post("/", handlers.CreateTask)
	api.Get("/", handlers.GetTasks)
	api.Put("/:id", handlers.UpdateTask)
	api.Delete("/:id", handlers.DeleteTask)

	log.Println("Сервер работает на http://localhost:3000")
	log.Fatal(app.Listen(":3000"))
}
