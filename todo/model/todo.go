package model

type Todo struct {
	Id          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"done"`
}

func GetDemoTodo() Todo {
	return Todo{
		Id:          1,
		Title:       "Implement REST for Octodo",
		Description: "We have a solid idea, we just need to implement it!",
		IsDone:      false,
	}
}
