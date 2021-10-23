package main

import (
	"github.com/gin-contrib/cors"
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
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())

	r.LoadHTMLGlob("front/Grootoony/Grootoony/index.html")
	r.Static("/static", "front/Grootoony/Grootoony")
	r.Static("/img", "front/Grootoony/Grootoony/img")
	r.Static("/css", "front/Grootoony/Grootoony/css")
	r.Static("/js", "front/Grootoony/Grootoony/js")
	r.Static("/mp3", "front/Grootoony/Grootoony/mp3")

	//r.LoadHTMLGlob("front/build/index.html")
	//r.Static("/static", "front/build/static/")
	//r.StaticFile("/manifest.json", "front/build/manifest.json")
	//r.StaticFile("/logo512.png", "front/build/logo512.png")
	//r.StaticFile("/logo192.png", "front/build/logo192.png")

	//sqlDB, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//db, err := gorm.Open(postgres.New(postgres.Config{
	//	Conn: sqlDB,
	//}), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//models.MigrateAll(db)

	authorized := r.Group("/api")
	authorized.Use()
	{
		r.GET("/guests", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", nil)
		})
	}
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	log.Fatal(r.Run(":" + port))
}
