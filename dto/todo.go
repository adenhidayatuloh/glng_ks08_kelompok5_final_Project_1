package dto

import (
	"time"

	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/entity"
)

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

type NewTodoRequest struct {
	Title     string `json:"title" example:"Belajar Flutter"`
	Completed bool   `json:"completed" example:"false"`
}

func (t *NewTodoRequest) TodoRequestToEntity() *entity.Todo {
	return &entity.Todo{
		Title:     t.Title,
		Completed: t.Completed,
	}
}

type NewTodoResponse struct {
	Message string         `json:"message" example:"Todo with id 69 has been successfully created"`
	Data    NewTodoRequest `json:"data"`
}
