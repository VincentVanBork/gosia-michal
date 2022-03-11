package controllers

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sethvargo/go-password/password"
	"main/models"
	"net/http"
	"net/mail"
	"net/url"
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
			if token == currentInvit.Token {
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
	res, err := password.Generate(16, 0, 0, false, true)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{
			"reason": "BAD PASS GEN",
		})
	}
	invitation.Token = res
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

func (u *InvitationsController) UpdateEmail(c *gin.Context) {
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
			if token == currentInvit.Token {
				*matchedInvite = currentInvit
			}
		}(&invitation, procInvite, token, &wg)
	}
	wg.Wait()
	u.Objects.Preload("Guests").Find(&invitation)
	email, isEmail := c.GetPostForm("Email")
	if isValidEmail(email) && isEmail {
		invitation.Email = sql.NullString{String: email, Valid: isValidEmail(email)}
	}

	hotel, isHotel := c.GetPostForm("Hotel")
	fmt.Println(hotel)
	if isHotel && hotel == "on" {
		invitation.Hotel = true
	}

	transport, isTransport := c.GetPostForm("Transport")
	fmt.Println(transport)
	if isTransport && transport == "on" {
		invitation.Transport = true
	}
	u.Objects.Save(invitation)
	q := url.Values{}
	q.Set("token", c.Param("token"))
	location := url.URL{Path: "/" + c.Param("token"), RawQuery: q.Encode()}
	c.Redirect(http.StatusSeeOther, location.RequestURI())
}

func (u *InvitationsController) RenderInvitation(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
