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
	createTodoQuery  = `INSERT INTO "todo" ("title", "completed", "created_at", "updated_at") VALUES ($1, $2, NOW(), NOW()) RETURNING "todo_id"`
	getAllTodosQuery = `SELECT * FROM todo;`
	getTodoByIDQuery = `SELECT * FROM todo WHERE todo_id = $1`
	updateTodoQuery  = `update "todo" set "title" = $1 , "completed" = $2,"updated_at" = NOW() where "todo_id" = $3 `
	deleteTodoQuery  = `DELETE FROM todo WHERE todo_id = $1`
)

func NewTodoRepositoryImpl(db *sql.DB) TodoRepository {
	return &todoRepositoryImpl{
		db: db,
	}
}

// CreateTodo implements TodoRepository.
func (t *todoRepositoryImpl) CreateTodo(todoPayload entity.Todo) (*entity.Todo, helper.MessageErr) {
	var todoID uint
	err := t.db.QueryRow(createTodoQuery, todoPayload.Title, todoPayload.Completed).Scan(&todoID)
	if err != nil {
		return nil, helper.NewInternalServerError("Failed to create new todo")
	}

	createdTodo := &entity.Todo{
		Todo_Id:    todoID,
		Title:      todoPayload.Title,
		Completed:  todoPayload.Completed,
		Created_At: todoPayload.Created_At,
		Updated_At: todoPayload.Updated_At,
	}

	return createdTodo, nil
}

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

func (t *todoRepositoryImpl) UpdateTodo(todoPayload entity.Todo) helper.MessageErr {
	tx, err := t.db.Begin()

	if err != nil {
		return helper.NewInternalServerError("Error in database")
	}

	_, err = tx.Exec(updateTodoQuery, todoPayload.Title, todoPayload.Completed, todoPayload.Todo_Id)

	if err != nil {
		tx.Rollback()
		return helper.NewInternalServerError("Error in executing query")

	}

	err = tx.Commit()

	if err != nil {
		tx.Rollback()
		return helper.NewInternalServerError("Error in commit database")
	}

	return nil
}

// DeleteTodo implements TodoRepository.
func (t *todoRepositoryImpl) DeleteTodo(id uint) helper.MessageErr {
	tx, err := t.db.Begin()
	if err != nil {
		return helper.NewInternalServerError("Error in database")
	}

	_, err = tx.Exec(deleteTodoQuery, id)
	if err != nil {
		tx.Rollback()
		return helper.NewInternalServerError("Error in executing query")
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return helper.NewInternalServerError("Error in commit database")
	}

	return nil
}
