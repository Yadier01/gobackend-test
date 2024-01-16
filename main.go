package main

import (
	"log"

	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	// GET /api/register
	app.Get("/api/", func(c *fiber.Ctx) {
		msg := "test"
		c.SendString(msg)
	})

	log.Fatal(app.Listen(":3000"))
}
