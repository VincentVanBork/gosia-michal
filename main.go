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
	var auth = middleware.TokenAuth{Objects: db}

	invitations := r.Group("/guest")
	//invitations.Use(auth.CheckAnyToken)
	var InvitationsController = controllers.InvitController{Objects: db}
	routes.AddFrontInvitation(invitations, &InvitationsController)

	secure := r.Group("/api")
	secure.Use(auth.CheckToken)
	routes.AddGuestsURLs(secure, &controllers.GuestController{Objects: db})
	routes.AddInvitationsURLs(secure, &InvitationsController)

	r.GET("/:token", auth.CheckAnyToken, InvitationsController.RenderInvitation)
	r.POST("/:token", InvitationsController.UpdateEmail)

	var TablesController = controllers.TableController{Objects: db}
	guestTables := r.Group("/lista", gin.BasicAuth(gin.Accounts{
		"michaljuras": "P@ndaExpress",
	}))
	guestTables.Static("/miejsca", "front/Grootoony/spis")
	guestTables.GET("/dane", TablesController.GetAll)

	//fmt.Printf(utils.HashPassword("L6A4YucGYKeDN5n5eKRHkMBtngDkAMV7"))
	log.Fatal(r.Run(":" + port))
}
