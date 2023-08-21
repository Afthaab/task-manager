package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/afthaab/task-manager/pkg/domain"
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
	"github.com/afthaab/task-manager/pkg/utility"
	"github.com/gin-gonic/gin"
)

type TaskHanlder struct {
	taskUseCase services.TaskUseCase
}

//////////////////////////////////// ---- MIDDLEWARES ---- ////////////////////////////////////

func (h *TaskHanlder) ValidateJWT(c *gin.Context) {
	tokenString, err := utility.GetTheBearerToken(c)
	if err != nil {
		utility.FailureResponse(c, http.StatusUnauthorized, "User not authorized", err)
	}

	ok, claims := h.taskUseCase.VerifyToken(tokenString)
	if !ok {
		utility.FailureResponse(c, http.StatusUnauthorized, "Token Verification Failed", errors.New("Token failed"))
	}

	status, err := h.taskUseCase.ValidateTheJwtUser(claims.Userid)
	if err != nil {
		utility.FailureResponse(c, status, "User authentication failed", err)
		return
	} else {
		uid := strconv.FormatUint(uint64(claims.Userid), 10)
		c.Set("userid", uid)
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
	strId := strconv.FormatUint(uint64(id), 10)
	responseMap["userId"] = strId

	if err != nil {
		utility.FailureResponse(c, status, "User verification unsuccessfull", err)
		return
	} else {
		utility.SuccessResponse(c, status, "User verification successfull", responseMap)
	}
}

//////////////////////////////////// ---- USER AUTHENTICATION ---- ////////////////////////////////////

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
	id := strconv.FormatUint(uint64(userData.Userid), 10)
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

//////////////////////////////////// ---- TASK MANAGEMENT ---- ////////////////////////////////////

func (h *TaskHanlder) AddTask(c *gin.Context) {
	userid := c.GetString("userid")
	taskData := domain.Task{}
	err := c.Bind(&taskData)
	if err != nil {
		utility.JsonValidationFailure(c, err)
		return
	}

	taskData, status, err := h.taskUseCase.AddTask(userid, taskData)

	//setting the response map
	responseMap := make(map[string]uint)
	responseMap["taskid"] = taskData.Taskid

	if err != nil {
		utility.FailureResponse(c, status, "Failed to add the Task", err)
		return
	} else {
		c.JSON(status, gin.H{
			"Success": true,
			"Message": "Task addedd successfully",
			"Data":    responseMap,
		})
	}
}

func (h *TaskHanlder) ViewAllTask(c *gin.Context) {
	userid := c.GetString("userid")

	taskDatas, status, err := h.taskUseCase.ViewAllTask(userid)

	//setting the response map
	responseMap := make(map[string][]domain.Task)
	responseMap["tasks"] = taskDatas

	if err != nil {
		utility.FailureResponse(c, status, "View all tasks unsuccessfull", err)
		return
	} else {
		c.JSON(status, gin.H{
			"Success": true,
			"Message": "View all tasks successfull",
			"Data":    responseMap,
		})
	}
}

func (h *TaskHanlder) EditTask(c *gin.Context) {
	userid := c.GetString("userid")
	taskData := domain.Task{}
	err := c.Bind(&taskData)
	if err != nil {
		utility.JsonValidationFailure(c, err)
		return
	}
	status, err := h.taskUseCase.EditTask(userid, taskData)

	//setting the response map
	uid := strconv.FormatUint(uint64(taskData.Taskid), 10)
	responseMap := make(map[string]string)
	responseMap["taskid"] = uid

	if err != nil {
		utility.FailureResponse(c, status, "Could not edit the task details", err)
		return
	} else {
		utility.SuccessResponse(c, status, "Successfully edited the task", responseMap)
	}
}

func (h *TaskHanlder) DeleteTask(c *gin.Context) {
	userid := c.GetString("userid")
	taskid := c.Query("taskid")

	status, err := h.taskUseCase.DeleteTask(userid, taskid)

	//setting the response map
	responseMap := make(map[string]string)
	responseMap["taskid"] = taskid

	if err != nil {
		utility.FailureResponse(c, status, "Could not delete the task", err)
		return
	} else {
		utility.SuccessResponse(c, status, "Successfully deleted the task", responseMap)
	}
}

func NewTaskHandler(TaskUseCase services.TaskUseCase) *TaskHanlder {
	return &TaskHanlder{
		taskUseCase: TaskUseCase,
	}
}
