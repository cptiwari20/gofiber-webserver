package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	getHandler void()= func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Chandra and this is new handler")
	}

	app.Get("/", getHandler)

	app.Listen(":3000")
}
