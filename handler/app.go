package handler

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/infra/database"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	database.InitDatabase()
	db := database.GetDatabaseInstance()

	_ = db // sementara, agar tidak error

	r := gin.Default()

	r.Run(":8080")

}
