package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	//Handle Cors
	app.Use(cors.New())

	log.Fatal(app.Listen(":" + port))
}
