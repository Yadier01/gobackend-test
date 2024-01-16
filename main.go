package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/api/", func(c *fiber.Ctx) {
		msg := "test"
		c.SendString(msg)
	})

	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}
