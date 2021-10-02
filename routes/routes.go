package routes

import (
	"github.com/gofiber/fiber/v2"
)

func ServeFront(app *fiber.App) {
	app.Static("/", "./front/build/")
}
