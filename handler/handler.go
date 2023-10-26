package handler

import (
	"net/http"
	"strconv"

	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
	service "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/service/todoService"
	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService service.TodoService
}

func NewTodoHandler(todoService service.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

// GetAllTodos godoc
//
//	@Summary		Get all todos
//	@Description	List todos
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	dto.GetAllTodosResponse
//	@Failure		500	{object}	errs.MessageErrData
//	@Router			/todos [get]
func (t *TodoHandler) GetAllTodos(ctx *gin.Context) {
	// Panggil Service
	todos, errService := t.todoService.GetAllTodos()
	if errService != nil {
		ctx.JSON(errService.StatusCode(), errService)
		return
	}
	// API Response
	ctx.JSON(http.StatusOK, todos) // GetTodoByID godoc
}

// GetTodoByID godoc
//
//	@Summary		Get a todo
//	@Description	Get a todo by id
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint	true	"Todo ID"
//	@Success		200	{object}	dto.GetTodoByIDResponse
//	@Failure		400	{object}	errs.MessageErrData
//	@Failure		404	{object}	errs.MessageErrData
//	@Router			/todos/{id} [get]
func (t *TodoHandler) GetTodoByID(ctx *gin.Context) {
	// Ambil Params
	id := ctx.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		newError := helper.NewBadRequest("ID should be an unsigned integer")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	todo, errService := t.todoService.GetTodoByID(uint(idUint))
	if errService != nil {
		ctx.JSON(errService.StatusCode(), errService)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}
