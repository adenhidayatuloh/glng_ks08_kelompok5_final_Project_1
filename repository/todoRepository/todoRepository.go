package repository

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/entity"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
)

type TodoRepository interface {
	CreateTodo(todoPayload entity.Todo) (*entity.Todo, helper.MessageErr)
	GetAllTodos() ([]entity.Todo, helper.MessageErr)
	GetTodoByID(id uint) (*entity.Todo, helper.MessageErr)
	UpdateTodo(todoPayload entity.Todo) helper.MessageErr
	DeleteTodo(id uint) helper.MessageErr
}
