package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	Objects *gorm.DB
}

type Create interface {
	Create(c *gin.Context) error
}

type GetAll interface {
	GetAll(c *gin.Context)
}

type GetOne interface {
	GetOne(c *gin.Context)
}
