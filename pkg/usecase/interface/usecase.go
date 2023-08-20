package interfaces

import (
	"github.com/afthaab/task-manager/pkg/domain"
)

type TaskUseCase interface {
	Signup(userData domain.User) (string, int, error)
	UserVerification(userData domain.User) (uint64, int, error)
	SignIn(user domain.User) (domain.User, int, error)
	GenerateAccessToken(userData domain.User, role string) (string, error)
}
