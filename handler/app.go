package handler

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/infra/database"
	repository "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/repository/todoRepository"
	service "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/service/todoService"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitDatabase()
	db := database.GetDatabaseInstance()

	todoRepository := repository.NewTodoRepositoryImpl(db)
	todoService := service.NewTodoServiceImpl(todoRepository)
	todoHandler := NewTodoHandler(todoService)

	r := gin.Default()

	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos", todoHandler.GetAllTodos)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.PUT("todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	r.Run(":8080")

}
