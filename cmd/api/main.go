package main

import (
	"github.com/IraIvanishak/stack_app/internal/scrapper"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		jobs := scrapper.Scrap()
		return c.JSON(jobs)
	})
	app.Listen(":8080")
}
