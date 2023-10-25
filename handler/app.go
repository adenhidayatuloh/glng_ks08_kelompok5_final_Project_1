package handler

import (
	"FinalProject1/infra/database"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitDatabase()
	db := database.GetDatabaseInstance()

	_ = db // sementara, agar tidak error

	r := gin.Default()

	r.Run(":8080")

}
