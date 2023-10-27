package handler

import (
	"net/http"
	"strconv"

	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/dto"
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

// // CreateTodo godoc
// //
// //	@Summary		Create a todo
// //	@Description	Create a todo by json
// //	@Tags			todos
// //	@Accept			json
// //	@Produce		json
// //	@Param			todo	body		dto.NewTodoRequest	true	"Create todo request body"
// //	@Success		201		{object}	dto.NewTodoResponse
// //	@Failure		422		{object}	errs.MessageErrData
// //	@Failure		500		{object}	errs.MessageErrData
// //	@Router			/todos [post]
// func (t *TodoHandler) CreateTodo(ctx *gin.Context) {
// 	var requestBody dto.NewTodoRequest

// 	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
// 		newError := helper.NewUnprocessableEntity(err.Error())
// 		ctx.JSON(newError.StatusCode(), newError)
// 		return
// 	}

// 	createdTodo, err := t.todoService.CreateTodo(&requestBody)
// 	if err != nil {
// 		ctx.JSON(err.StatusCode(), err)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, createdTodo)
// }

// CreateTodo godoc
//
// @Summary Create a new todo
// @Description Create a new todo with the provided data
// @Tags todos
// @Accept json
// @Produce json
// @Param todo body dto.NewTodoRequest true "Todo data"
// @Success 201 {object} dto.GetTodoByIDResponse
// @Failure 400 {object} errs.MessageErrData
// @Router /todos [post]
func (t *TodoHandler) CreateTodo(ctx *gin.Context) {
	var newTodoRequest dto.NewTodoRequest

	if err := ctx.ShouldBindJSON(&newTodoRequest); err != nil {
		newError := helper.NewBadRequest("Invalid request body")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	createdTodo, err := t.todoService.CreateTodo(newTodoRequest)
	if err != nil {
		ctx.JSON(err.StatusCode(), err)
		return
	}

	response := &dto.GetTodoByIDResponse{
		Message: "success",
		Data: dto.DetailTodo{
			ID:        createdTodo.Data.ID,
			Title:     createdTodo.Data.Title,
			Completed: createdTodo.Data.Completed,
			CreatedAt: createdTodo.Data.CreatedAt,
			UpdatedAt: createdTodo.Data.UpdatedAt,
		},
	}

	ctx.JSON(http.StatusCreated, response)
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

// UpdateTodo godoc
//
//	@Summary		Update todo
//	@Description	Update a todo by json
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			id		path		uint				true	"Todo ID"
//	@Param			todo	body		dto.NewTodoRequest	true	"Update todo request body"
//	@Success		200		{object}	dto.GetTodoByID
//	@Failure		400		{object}	errs.MessageErrData
//	@Failure		422		{object}	errs.MessageErrData
//	@Failure		404		{object}	errs.MessageErrData
//	@Failure		500		{object}	errs.MessageErrData
//	@Router			/todos/{id} [put]
func (t *TodoHandler) UpdateTodo(ctx *gin.Context) {

	todoID := ctx.Param("id")

	ConvTodoID, err := strconv.Atoi(todoID)

	if err != nil {
		newError := helper.NewBadRequest("Error convert to int")
		ctx.JSON(newError.StatusCode(), newError)
		return
	}

	var newTodoRequest dto.NewTodoRequest

	err = ctx.ShouldBindJSON(&newTodoRequest)

	if err != nil {
		newError := helper.NewBadRequest("Error binding json")
		ctx.JSON(newError.StatusCode(), newError)
		return

	}

	errService := t.todoService.UpdateTodo(uint(ConvTodoID), newTodoRequest)

	if err != nil {

		ctx.JSON(errService.StatusCode(), errService)
		return

	}

	todo, errService := t.todoService.GetTodoByID(uint(ConvTodoID))
	if errService != nil {
		ctx.JSON(errService.StatusCode(), errService)
		return
	}

	ctx.JSON(http.StatusOK, todo)
}
