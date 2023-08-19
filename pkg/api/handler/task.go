package handler

import (
	"net/http"

	"github.com/afthaab/task-manager/pkg/domain"
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
	"github.com/gin-gonic/gin"
)

type TaskHanlder struct {
	taskUseCase services.TaskUseCase
}

func (h *TaskHanlder) Signup(c *gin.Context) {
	userData := domain.User{}
	err := c.Bind(&userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Error":   err.Error(),
			"Message": "Could not process the given data",
		})
		return
	}

	email, status, err := h.taskUseCase.Signup(userData)
	if err != nil {
		c.JSON(status, gin.H{
			"Success": false,
			"Error":   err.Error(),
			"Message": "User Signup unsuccessfull",
		})
		return
	} else {
		c.JSON(status, gin.H{
			"Success": true,
			"Message": "User Signup successfull",
			"Data":    email,
		})
	}
}

func (h *TaskHanlder) UserVerification(c *gin.Context) {
	userData := domain.User{}
	err := c.Bind(&userData)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Error":   err.Error(),
			"Message": "Could not process the given data",
		})
		return
	}
	id, status, err := h.taskUseCase.UserVerification(userData)
	if err != nil {
		c.JSON(status, gin.H{
			"Success": false,
			"Error":   err.Error(),
			"Message": "User verification unsuccessfull",
		})
		return
	} else {
		c.JSON(status, gin.H{
			"Success": true,
			"Message": "User verification successfull",
			"Data":    id,
		})
	}
}

func NewTaskHandler(TaskUseCase services.TaskUseCase) *TaskHanlder {
	return &TaskHanlder{
		taskUseCase: TaskUseCase,
	}
}
