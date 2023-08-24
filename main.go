package main

import (
	"github.com/gofiber/fiber/v2"
)

func main() {
  app := fiber.New()

  app.Static("/", "./public")

  app.Post("/echo", func(c *fiber.Ctx) error {
    return c.Send(c.Body())
  })

  app.Listen(":3000")
}
