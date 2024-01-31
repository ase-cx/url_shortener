package main

import (
	"github.com/gofiber/fiber/v2"

	"ase.cx/url-shortener/database"
	"ase.cx/url-shortener/handlers"
	"ase.cx/url-shortener/middlewares"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Post("/login", handlers.Login)
	app.Post("/register", handlers.Register)

	app.Get("/:s", handlers.Redirect)

	app.Use(middlewares.Protected())
	app.Post("/shorten", handlers.Shorten)

	app.Listen("0.0.0.0:3000")
}
