package main

import (
	"github.com/Swiddis/octodo/todo/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/healthcheck", handlers.Healthcheck)
	r.GET("/todo/:id", handlers.GetTodo)
	r.POST("/todo", handlers.PostTodo)
	r.PUT("/todo/:id", handlers.PutTodo)
	r.DELETE("/todo/:id", handlers.DeleteTodo)
	r.Run()
}
