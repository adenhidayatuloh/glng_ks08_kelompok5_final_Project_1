package service

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/dto"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
)

type TodoService interface {
	CreateTodo(newTodoRequest dto.NewTodoRequest) (*dto.GetTodoByIDResponse, helper.MessageErr)
	GetAllTodos() (*dto.GetAllTodosResponse, helper.MessageErr)
	GetTodoByID(id uint) (*dto.GetTodoByIDResponse, helper.MessageErr)
	UpdateTodo(todo_id uint, newOrderRequest dto.NewTodoRequest) helper.MessageErr
	DeleteTodo(id uint) helper.MessageErr
}
