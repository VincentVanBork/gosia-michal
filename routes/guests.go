package routes

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
)

func AddGuestsURLs(r *gin.RouterGroup, controller *controllers.GuestController) {
	r.GET("guests/:token", controller.GetAll)
	r.POST("guests/create/:token", controller.Create)
	r.POST("guests/:token", controller.GetOne)
}
