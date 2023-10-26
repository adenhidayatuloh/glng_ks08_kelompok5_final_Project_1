package dto

import "time"

type GetAllTodosResponse struct {
	Message string  `json:"message"`
	Data    []Todos `json:"data"`
}

type GetTodoByIDResponse struct {
	Message string     `json:"message"`
	Data    DetailTodo `json:"data"`
}

type Todos struct {
	ID        uint   `json:"id" example:"69"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type DetailTodo struct {
	ID        uint      `json:"id" example:"69"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
