package repository

import (
	"database/sql"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/entity"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
)

type todoRepositoryImpl struct {
	db *sql.DB
}

const (
	getAllTodosQuery = `SELECT * FROM todo;`
	getTodoByIDQuery = `SELECT * FROM todo WHERE todo_id = $1`
)

// GetAllTodos implements TodoRepository.
func (t *todoRepositoryImpl) GetAllTodos() ([]entity.Todo, helper.MessageErr) {
	var todos []entity.Todo
	// Query
	rows, errQuery := t.db.Query(getAllTodosQuery)
	if errQuery != nil {
		return nil, helper.NewInternalServerError("Failed to get all todos")
	}
	// Mengisi slice Todos
	for rows.Next() {
		var todo entity.Todo

		errScan := rows.Scan(&todo.Todo_Id, &todo.Title, &todo.Completed, &todo.Created_At, &todo.Updated_At)
		if errScan != nil {
			return nil, helper.NewInternalServerError("Failed to get all todos")
		}
		todos = append(todos, todo)
	}
	// Return
	return todos, nil
}

// GetTodoByID implements TodoRepository.
func (t *todoRepositoryImpl) GetTodoByID(id uint) (*entity.Todo, helper.MessageErr) {
	var todo entity.Todo
	// Query
	errQuery := t.db.QueryRow(getTodoByIDQuery, id).Scan(&todo.Todo_Id, &todo.Title, &todo.Completed, &todo.Created_At, &todo.Updated_At)
	if errQuery == sql.ErrNoRows {
		return nil, helper.NewNotFound("Todos Not Found")
	}
	// Return
	return &todo, nil
}

func NewTodoRepositoryImpl(db *sql.DB) TodoRepository {
	return &todoRepositoryImpl{
		db: db,
	}
}
