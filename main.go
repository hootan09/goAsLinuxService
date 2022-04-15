package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create new Fiber instance
	app := fiber.New()

	// Make path with some content
	app.Get("/hello", func(c *fiber.Ctx) error {
		// Return a string with a dummy text
		return c.SendString("Hello, World 👋!")
	})

	// Start server on http://localhost:3000 or error
	log.Fatal(app.Listen(":3000"))
}
