package routes

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
)

func AddInvitationsURLs(r *gin.RouterGroup, controller *controllers.InvitController) {
	r.POST("invitations/create/:token", controller.Create)
	r.POST("invitations/update/:token", controller.UpdateGuests)
	r.GET("invitations/:token", controller.GetAll)
}

func AddFrontInvitation(r *gin.RouterGroup, controller *controllers.InvitController) {
	r.GET("invitations/get/:token", controller.GetOne)
	r.POST("invitations/update/:token", controller.UpdateEmail)
	//r.POST("invitations/create/:token", controller.Create)
	//r.POST("invitations/update/:token", controller.UpdateGuests)
}
