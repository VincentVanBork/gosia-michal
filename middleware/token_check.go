package middleware

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
	"main/models"
	"main/utils"
	"os"
	"sync"
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
	if !isAuthorized {
		c.AbortWithStatusJSON(401, gin.H{
			"reason": "no hacky hacky",
		})
	}
}

func (t *TokenAuth) CheckAnyToken(c *gin.Context) {
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
	var wg sync.WaitGroup
	for _, m := range invitations {
		procInvite := m
		wg.Add(1)
		go func(authStatus *bool, currentInvit models.Invitation, token string, wgroup *sync.WaitGroup) {
			defer wg.Done()
			if utils.CheckPasswordHash(token, currentInvit.Token) {
				*authStatus = true
			}
		}(&isAuthorized, procInvite, token, &wg)
	}
	wg.Wait()
	if !isAuthorized {
		c.AbortWithStatusJSON(401, gin.H{
			"reason": "no hacky hacky",
		})
	} else {
		c.Next()
	}

}
