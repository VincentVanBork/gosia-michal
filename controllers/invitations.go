package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"main/models"
	"main/utils"
	"net/http"
	"sync"
)

type InvitationsController Controller

func (u *InvitationsController) GetAll(c *gin.Context) {
	var invitations []models.Invitation
	u.Objects.Find(&invitations)
	c.JSON(200, invitations)
}

func (u *InvitationsController) GetOne(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.AbortWithStatusJSON(400, gin.H{
			"reason": "BAD BIND",
		})
	}
	var invitation models.Invitation
	var invitations []models.Invitation
	u.Objects.Find(&invitations)
	var wg sync.WaitGroup
	for _, m := range invitations {
		procInvite := m
		wg.Add(1)
		go func(matchedInvite *models.Invitation, currentInvit models.Invitation, token string, wgroup *sync.WaitGroup) {
			defer wg.Done()
			if utils.CheckPasswordHash(token, currentInvit.Token) {
				*matchedInvite = currentInvit
			}
		}(&invitation, procInvite, token, &wg)
	}
	wg.Wait()
	u.Objects.Preload("Guests").Find(&invitation)
	c.JSON(200, invitation)
}

func (u *InvitationsController) Create(c *gin.Context) {

	var invitation models.Invitation
	err := c.Bind(&invitation)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"reason": "BAD BIND",
		})
	}
	// Generate a password that is 64 characters long with 10 digits, 10 symbols,
	// allowing upper and lower case letters, disallowing repeat characters.
	res, err := password.Generate(64, 0, 0, false, true)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"reason": "BAD PASS GEN",
		})
	}
	hash, err := utils.HashPassword(res)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"reason": "BAD HASH PASS",
		})
	}
	invitation.Token = hash
	u.Objects.Create(&invitation)

	c.JSON(200, gin.H{
		"invitation": invitation,
		"password":   res,
	})

}

func (u *InvitationsController) UpdateGuests(c *gin.Context) {
	var newInvitation models.Invitation
	err := c.Bind(&newInvitation)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"reason": "BAD BIND",
		})
	}
	var oldInvitation models.Invitation
	u.Objects.First(&oldInvitation, newInvitation.ID)
	oldInvitation.Guests = newInvitation.Guests
	u.Objects.Save(oldInvitation)
}

func (u *InvitationsController) RenderInvitation(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
