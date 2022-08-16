package model

type Todo struct {
	Id          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsDone      bool   `json:"done"`
}

func GetDemoTodo(id uint64) Todo {
	return Todo{
		Id:          id,
		Title:       "Implement REST for Octodo",
		Description: "We have a solid idea, we just need to implement it!",
		IsDone:      false,
	}
}
