package main

import (
	"github.com/Swiddis/octodo/todo/handlers"
	"github.com/Swiddis/octodo/todo/model"
	"github.com/gin-gonic/gin"
)

func AddRoutes(r *gin.Engine) {
	r.GET("/healthcheck", handlers.Healthcheck)
	r.GET("/todo/:id", handlers.GetTodo)
	r.POST("/todo", handlers.PostTodo)
	r.PUT("/todo/:id", handlers.PutTodo)
	r.DELETE("/todo/:id", handlers.DeleteTodo)
}

func main() {
	db, _ := model.OpenDB("sqlite")
	handlers.Use(db)
	r := gin.Default()
	AddRoutes(r)
	r.Run()
}
