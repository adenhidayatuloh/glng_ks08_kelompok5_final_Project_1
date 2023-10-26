package service

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/dto"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
	repository "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/repository/todoRepository"
)

type todoServiceImpl struct {
	repo repository.TodoRepository
}

// GetAllTodos implements TodoService.
func (t *todoServiceImpl) GetAllTodos() (*dto.GetAllTodosResponse, helper.MessageErr) {
	var todosData []dto.Todos
	// Panggil Repository
	todos, err := t.repo.GetAllTodos()
	if err != nil {
		return nil, err
	}

	for _, todo := range todos {
		todosData = append(todosData, dto.Todos{
			ID:        todo.Todo_Id,
			Title:     todo.Title,
			Completed: todo.Completed,
		})
	}

	response := &dto.GetAllTodosResponse{
		Message: "success",
		Data:    todosData,
	}

	return response, nil
}

// GetTodoByID implements TodoService.
func (t *todoServiceImpl) GetTodoByID(id uint) (*dto.GetTodoByIDResponse, helper.MessageErr) {
	// Panggil Repository
	todo, err := t.repo.GetTodoByID(id)
	if err != nil {
		return nil, err
	}
	// Membuat Response
	response := &dto.GetTodoByIDResponse{
		Message: "success",
		Data: dto.DetailTodo{
			ID:        todo.Todo_Id,
			Title:     todo.Title,
			Completed: todo.Completed,
			CreatedAt: todo.Created_At,
			UpdatedAt: todo.Updated_At,
		},
	}
	// Return response
	return response, nil
}

func NewTodoServiceImpl(r repository.TodoRepository) TodoService {
	return &todoServiceImpl{
		repo: r,
	}
}
