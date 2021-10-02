package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/keyauth/v2"
	"main/models"
	"main/routes"

	"log"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	app := fiber.New()
	routes.ServeFront(app)
	routes.SetupAuthRoutes(app)
	//Handle Cors
	app.Use(cors.New())
	app.Use(keyauth.New(keyauth.Config{KeyLookup: "query:token", Validator: models.ValidateGuestToken}))
	log.Fatal(app.Listen(":" + port))
}
