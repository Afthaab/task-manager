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

	return &ServerHTTP{
		engine: engine,
	}
}

func (s *ServerHTTP) Start() {
	s.engine.Run(os.Getenv("APP_PORT"))
}
