package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Healthcheck(c *gin.Context) {
	c.String(200, "HEALTHY")
}

func Use(database *gorm.DB) {
	db = database
}
