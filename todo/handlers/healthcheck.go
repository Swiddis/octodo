package handlers

import (
	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) {
	c.String(200, "HEALTHY")
}
