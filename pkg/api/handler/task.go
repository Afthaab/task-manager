package handler

import (
	"strconv"

	"github.com/afthaab/task-manager/pkg/domain"
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
	"github.com/afthaab/task-manager/pkg/utility"
	"github.com/gin-gonic/gin"
)

type TaskHanlder struct {
	taskUseCase services.TaskUseCase
}

func (h *TaskHanlder) SignIn(c *gin.Context) {
	userData := domain.User{}
	err := c.Bind(&userData)
	if err != nil {
		utility.JsonValidationFailure(c, err)
		return
	}

	userData, status, err := h.taskUseCase.SignIn(userData)
	if err != nil {
		utility.FailureResponse(c, status, "User SignIn unsuccessfull", err)
		return
	}
	accessToken, err := h.taskUseCase.GenerateAccessToken(userData, "user")

	//setting response map
	responseMap := make(map[string]string)
	responseMap["accessToken"] = accessToken
	id := strconv.FormatUint(userData.Id, 10)
	responseMap["userId"] = id

	if err != nil {
		utility.FailureResponse(c, status, "User SignIn unsuccessfull", err)
		return
	} else {
		utility.SuccessResponse(c, status, "User SignIn successfull", responseMap)
	}
}

func (h *TaskHanlder) Signup(c *gin.Context) {
	userData := domain.User{}
	err := c.Bind(&userData)
	if err != nil {
		utility.JsonValidationFailure(c, err)
		return
	}

	email, status, err := h.taskUseCase.Signup(userData)

	//setting response map
	responseMap := make(map[string]string)
	responseMap["email"] = email

	if err != nil {
		utility.FailureResponse(c, status, "User Signup unsuccessfull", err)
		return
	} else {
		utility.SuccessResponse(c, status, "User Signup successfull", responseMap)
	}
}

func (h *TaskHanlder) UserVerification(c *gin.Context) {
	userData := domain.User{}
	err := c.Bind(&userData)
	if err != nil {
		utility.JsonValidationFailure(c, err)
		return
	}
	id, status, err := h.taskUseCase.UserVerification(userData)

	//setting response map
	responseMap := make(map[string]string)
	strId := strconv.FormatUint(id, 10)
	responseMap["userId"] = strId

	if err != nil {
		utility.FailureResponse(c, status, "User verification unsuccessfull", err)
		return
	} else {
		utility.SuccessResponse(c, status, "User verification successfull", responseMap)
	}
}

func NewTaskHandler(TaskUseCase services.TaskUseCase) *TaskHanlder {
	return &TaskHanlder{
		taskUseCase: TaskUseCase,
	}
}
