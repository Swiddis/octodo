package handlers

import (
	"fmt"
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
	if Check(err, fmt.Sprintf("Unable to parse id '%v', expected unsigned int", rawId), c) {
		return
	}
	todo := model.GetDemoTodo(id)
	c.JSON(200, todo)
}
