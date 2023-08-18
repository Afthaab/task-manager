package handler

import (
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
)

type TaskHanlder struct {
	taskUseCase services.TaskUseCase
}

func NewTaskHandler(TaskUseCase services.TaskUseCase) *TaskHanlder {
	return &TaskHanlder{
		taskUseCase: TaskUseCase,
	}
}
