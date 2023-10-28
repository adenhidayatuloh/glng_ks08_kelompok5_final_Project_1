package service

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/dto"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/entity"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
	repository "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/repository/todoRepository"
)

type todoServiceImpl struct {
	repo repository.TodoRepository
}

// CreateTodo implements TodoService.
func (t *todoServiceImpl) CreateTodo(newTodoRequest dto.NewTodoRequest) (*dto.GetTodoByIDResponse, helper.MessageErr) {
	todoPayload := entity.Todo{
		Title:     newTodoRequest.Title,
		Completed: newTodoRequest.Completed,
	}

	createdTodo, err := t.repo.CreateTodo(todoPayload)
	if err != nil {
		return nil, err
	}

	response := &dto.GetTodoByIDResponse{
		Message: "success",
		Data: dto.DetailTodo{
			ID:        createdTodo.Todo_Id,
			Title:     createdTodo.Title,
			Completed: createdTodo.Completed,
			CreatedAt: createdTodo.Created_At,
			UpdatedAt: createdTodo.Updated_At,
		},
	}

	return response, nil
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

func (t *todoServiceImpl) UpdateTodo(todo_id uint, newTodoRequest dto.NewTodoRequest) helper.MessageErr {
	todoPayload := entity.Todo{
		Todo_Id:   todo_id,
		Title:     newTodoRequest.Title,
		Completed: newTodoRequest.Completed,
	}

	err := t.repo.UpdateTodo(todoPayload)

	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo implements TodoService.
func (t *todoServiceImpl) DeleteTodo(id uint) helper.MessageErr {
	_, err := t.repo.GetTodoByID(id)
	if err != nil {
		return err
	}

	err = t.repo.DeleteTodo(id)
	if err != nil {
		return helper.NewInternalServerError("Failed to delete todo")
	}

	return nil
}

func NewTodoServiceImpl(r repository.TodoRepository) TodoService {
	return &todoServiceImpl{
		repo: r,
	}
}
