package model

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Title       string `json:"Title"`
	Description string `json:"Description"`
	IsDone      bool   `json:"Done"`
}

func GetDemoTodo(id uint64) Todo {
	return Todo{
		Title:       "Implement REST for Octodo",
		Description: "We have a solid idea, we just need to implement it!",
		IsDone:      false,
	}
}
