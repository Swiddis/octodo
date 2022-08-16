package main

import (
	"github.com/Swiddis/octodo/todo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", handlers.Healthcheck)
	r.Run()
}
