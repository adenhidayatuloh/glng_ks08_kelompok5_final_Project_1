package handler

import (
	"os"

	_ "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/docs"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/infra/database"
	repository "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/repository/todoRepository"
	service "github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/service/todoService"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// title Todo API
// @version 1.0
// @description this is a sample service for managing todos
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email asdf@gmail.com
// @license.name Apache 2.0
// @license.urhttp://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func StartApp() {
	database.InitDatabase()
	db := database.GetDatabaseInstance()

	todoRepository := repository.NewTodoRepositoryImpl(db)
	todoService := service.NewTodoServiceImpl(todoRepository)
	todoHandler := NewTodoHandler(todoService)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos", todoHandler.GetAllTodos)
	r.GET("/todos/:id", todoHandler.GetTodoByID)
	r.PUT("todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	port := os.Getenv("PORT")

	r.Run(":" + port)

}
