package api

import (
	"os"

	"github.com/afthaab/task-manager/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(taskHandler *handler.TaskHanlder) *ServerHTTP {
	engine := gin.New()

	engine.Use(gin.Logger())

	taskManager := engine.Group("/taskmanager")
	{
		user := taskManager.Group("/user")
		{
			user.POST("/signup", taskHandler.Signup)
			user.POST("/verify", taskHandler.UserVerification)
			user.POST("/signin", taskHandler.SignIn)
		}

		task := taskManager.Group("/task")
		task.Use(taskHandler.ValidateJWT)
		{
			task.GET("/view/all", taskHandler.ViewAllTask)
			task.POST("/add", taskHandler.AddTask)
			task.PUT("/edit", taskHandler.EditTask)
			task.DELETE("/delete", taskHandler.DeleteTask)
		}

	}

	return &ServerHTTP{
		engine: engine,
	}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(os.Getenv("APP_PORT"))
}
