package routes

import (
	"github.com/gofiber/fiber/v2"
)

func ServeFront(app *fiber.App) {
	app.Static("/", "./front/build")
}

func SetupAuthRoutes(app *fiber.App) {
	app.Get("/auth", func(c *fiber.Ctx) error {
		token, _ := c.Locals("token").(string)
		return c.SendString(token)
	})
}
