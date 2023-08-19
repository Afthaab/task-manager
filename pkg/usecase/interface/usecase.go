package interfaces

import "github.com/afthaab/task-manager/pkg/domain"

type TaskUseCase interface {
	Signup(userData domain.User) (string, int, error)
	UserVerification(userData domain.User) (uint, int, error)
}
