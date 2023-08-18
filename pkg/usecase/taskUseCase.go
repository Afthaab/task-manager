package usecase

import (
	interfaces "github.com/afthaab/task-manager/pkg/repository/interface"
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
)

type TaskUseCase struct {
	taskRepo interfaces.TaskRepository
}

func NewTaskUseCase(TaskRepo interfaces.TaskRepository) services.TaskUseCase {
	return &TaskUseCase{
		taskRepo: TaskRepo,
	}
}
