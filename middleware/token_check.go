package middleware

import (
	"github.com/gin-gonic/gin"
	"main/controllers"
	"main/models"
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

	if token == "L6A4YucGYKeDN5n5eKRHkMBtngDkAMV7" {
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

	if token == "L6A4YucGYKeDN5n5eKRHkMBtngDkAMV7" {
		isAuthorized = true
		c.Next()
	}
	var wg sync.WaitGroup
	for _, m := range invitations {
		procInvite := m
		wg.Add(1)
		go func(authStatus *bool, currentInvit models.Invitation, token string, wgroup *sync.WaitGroup) {
			defer wg.Done()
			if token == currentInvit.Token {
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
