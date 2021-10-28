package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
)

type GuestController Controller

func (u *GuestController) GetAll(c *gin.Context) {
	var guests []models.Guest
	u.Objects.Find(&guests)
	c.JSON(200, guests)
}

func (u *GuestController) GetOne(c *gin.Context) {
	var guest models.Guest

	err := c.Bind(&guest)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"reason": "BAD BIND",
		})
	}
	u.Objects.Find(&guest)
	c.JSON(200, guest)
}

func (u *GuestController) Create(c *gin.Context) {
	var guest models.Guest
	err := c.Bind(&guest)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{
			"reason": "BAD BIND",
		})
	}

	u.Objects.Create(&guest)

	c.JSON(200, guest)
}
