package middleware

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
	"main/models"
	"main/utils"
	"os"
)

type TokenAuth controllers.Controller

func (t *TokenAuth) CheckToken(c *gin.Context) {
	isAuthorized := false
	token := c.Param("token")
	if token == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"reason": "BAD BIND",
		})
	}
	var invitations []models.Invitation
	t.Objects.Find(&invitations)

	hash := os.Getenv("GOSIA_MICHAL")
	if utils.CheckPasswordHash(token, hash) {
		isAuthorized = true
		c.Next()
	}
	//for _, m := range invitations {
	//	if utils.CheckPasswordHash(token, m.Token) {
	//		isAuthorized = true
	//		c.Next()
	//	}
	//}
	if !isAuthorized {
		c.AbortWithStatusJSON(401, gin.H{
			"reason": "no hacky hacky",
		})
	}

}
