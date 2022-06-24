package controllers

import (
	"github.com/gin-gonic/gin"
	"main/models"
)

type TableController Controller

func (u *TableController) GetAll(c *gin.Context) {
	var tables []models.TableInfo
	u.Objects.Find(&tables)
	c.JSON(200, tables)
}
