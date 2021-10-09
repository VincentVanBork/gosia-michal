package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("front/build/index.html")
	router.Static("/static", "front/build/static")
	router.StaticFile("/manifest.json", "front/build/manifest.json")
	router.StaticFile("/logo512.png", "front/build/logo512.png")
	router.StaticFile("/logo192.png", "front/build/logo192.png")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	log.Fatal(router.Run(":" + port))
}
