package handlers

import (
	"strconv"

	"github.com/Swiddis/octodo/todo/model"
	"github.com/gin-gonic/gin"
)

func Check(err error, msgIfError string, ctx *gin.Context) bool {
	if err == nil {
		return false
	}
	ctx.JSON(400, gin.H{
		"error": msgIfError,
	})
	return true
}

func GetTodo(c *gin.Context) {
	rawId := c.Param("id")
	id, err := strconv.ParseUint(rawId, 10, 0)
	if Check(err, "unable to parse id", c) {
		return
	}
	todo := model.GetDemoTodo(id)
	c.JSON(200, todo)
}

func PostTodo(c *gin.Context) {
	todo := model.Todo{}
	err := c.ShouldBindJSON(&todo)
	if Check(err, "unable to parse Todo", c) {
		return
	}
	c.JSON(201, todo)
}

func DeleteTodo(c *gin.Context) {
	rawId := c.Param("id")
	_, err := strconv.ParseUint(rawId, 10, 0)
	if Check(err, "unable to parse id", c) {
		return
	}
	c.Status(204)
}
