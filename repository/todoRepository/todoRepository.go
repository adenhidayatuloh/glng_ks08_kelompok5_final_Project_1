package repository

import (
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/entity"
	"github.com/adenhidayatuloh/glng_ks08_kelompok5_final_Project_1/helper"
)

type TodoRepository interface {
	GetAllTodos() ([]entity.Todo, helper.MessageErr)
	GetTodoByID(id uint) (*entity.Todo, helper.MessageErr)
}
