//go:build wireinject
// +build wireinject

package di

import (
	"github.com/afthaab/task-manager/pkg/api"
	"github.com/afthaab/task-manager/pkg/api/handler"
	"github.com/afthaab/task-manager/pkg/db"
	"github.com/afthaab/task-manager/pkg/repository"
	"github.com/afthaab/task-manager/pkg/usecase"
	"github.com/google/wire"
)

func InitializeApi() (*api.ServerHTTP, error) {
	wire.Build(db.ConnectToDatabase,
		repository.NewTaskRepository,
		usecase.NewTaskUseCase,
		handler.NewTaskHandler,
		api.NewServerHTTP)
	return &api.ServerHTTP{}, nil
}
