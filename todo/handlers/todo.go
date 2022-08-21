package handlers

import (
	"strconv"

	"github.com/Swiddis/octodo/todo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func GetTodo(c *gin.Context, db *gorm.DB) {
	rawId := c.Param("id")
	if id, err := strconv.Atoi(rawId); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		var todo model.Todo
		if err := db.Take(&todo, id).Error; err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "not found",
			})
		} else {
			c.JSON(200, todo)
		}
	}
}

func PostTodo(c *gin.Context, db *gorm.DB) {
	todo := model.Todo{}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(400, gin.H{
			"error": "unable to parse todo",
		})
		return
	}
	db.Create(&todo)
	c.JSON(201, todo)
}

func PutTodo(c *gin.Context, db *gorm.DB) {
	rawId := c.Param("id")
	if id, err := strconv.Atoi(rawId); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		var todo, newTodo model.Todo
		if err := c.ShouldBindJSON(&newTodo); err != nil {
			c.JSON(400, gin.H{
				"error": "unable to parse todo",
			})
			return
		}
		if err := db.Where("ID = ?", id).First(&todo).Error; err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "not found",
			})
		} else {
			db.Model(&todo).Update("Title", newTodo.Title)
			db.Model(&todo).Update("Description", newTodo.Description)
			db.Model(&todo).Update("IsDone", newTodo.IsDone)
			c.JSON(200, todo)
		}
	}
}

func DeleteTodo(c *gin.Context, db *gorm.DB) {
	rawId := c.Param("id")
	if id, err := strconv.Atoi(rawId); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid id",
		})
		return
	} else {
		var todo model.Todo
		if err := db.Delete(&todo, id).Error; err == gorm.ErrRecordNotFound {
			c.JSON(404, gin.H{
				"error": "not found",
			})
		} else {
			c.Status(204)
		}
	}
}
