package main

import (
	"github.com/Swiddis/octodo/todo/handlers"
	"github.com/Swiddis/octodo/todo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/healthcheck", handlers.Healthcheck)
	r.GET("/todo/:id", func(c *gin.Context) {
		handlers.GetTodo(c, db)
	})
	r.POST("/todo", func(c *gin.Context) {
		handlers.PostTodo(c, db)
	})
	r.PUT("/todo/:id", func(c *gin.Context) {
		handlers.PutTodo(c, db)
	})
	r.DELETE("/todo/:id", func(c *gin.Context) {
		handlers.DeleteTodo(c, db)
	})
}

func main() {
	db, err := model.OpenDB("postgres")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	AddRoutes(r, db)
	r.Run()
}
