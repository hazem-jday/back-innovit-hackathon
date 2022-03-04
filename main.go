// @/main.go
package main

import (
	"back/config"
	"back/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	config.Connect()

	app.Get("/users", handlers.GetUsers)
	app.Get("/user/:id", handlers.GetUser)
	app.Post("/users", handlers.AddUser)
	app.Put("/users/:id", handlers.UpdateUser)
	app.Delete("/users/:id", handlers.RemoveUser)

	app.Get("/posts", handlers.GetPosts)
	app.Get("/post/:id", handlers.GetPost)
	app.Post("/posts", handlers.AddPost)
	app.Put("/posts/:id", handlers.UpdatePost)
	app.Delete("/posts/:id", handlers.RemovePost)

	log.Fatal(app.Listen(":8080"))
}
