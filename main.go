package main

import (
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"main/controllers"
	"main/middleware"
	"main/models"
	"main/routes"
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

	r.LoadHTMLGlob("front/Grootoony/index.html")
	r.Static("/static", "front/Grootoony/")
	r.Static("/img", "front/Grootoony/img")
	r.Static("/css", "front/Grootoony/css")
	r.Static("/js", "front/Grootoony/js")
	r.Static("/mp3", "front/Grootoony/mp3")

	sqlDB, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	models.MigrateAll(db)

	invitations := r.Group("/guest")
	var invitationsController = controllers.InvitationsController{Objects: db}
	routes.AddFrontInvitation(invitations, &invitationsController)

	var auth = middleware.TokenAuth{Objects: db}
	secure := r.Group("/api")
	secure.Use(auth.CheckToken)
	routes.AddGuestsURLs(secure, &controllers.GuestController{Objects: db})
	routes.AddInvitationsURLs(secure, &invitationsController)

	r.GET("/:token", auth.CheckToken, invitationsController.RenderInvitation)
	//fmt.Printf(utils.HashPassword("L6A4YucGYKeDN5n5eKRHkMBtngDkAMV7"))
	log.Fatal(r.Run(":" + port))
}
