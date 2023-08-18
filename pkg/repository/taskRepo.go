package repository

import (
	interfaces "github.com/afthaab/task-manager/pkg/repository/interface"
	"gorm.io/gorm"
)

type TaskDatabase struct {
	db *gorm.DB
}

func NewTaskRepository(DB *gorm.DB) interfaces.TaskRepository {
	return &TaskDatabase{
		db: DB,
	}
}
