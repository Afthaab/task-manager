package interfaces

import (
	"github.com/afthaab/task-manager/pkg/domain"
)

type TaskUseCase interface {
	// Middleware
	VerifyToken(token string) (bool, *domain.JWTClaims)
	GenerateAccessToken(userData domain.User, role string) (string, error)
	ValidateTheJwtUser(userid uint) (int, error)

	// User Authentication
	Signup(userData domain.User) (string, int, error)
	UserVerification(userData domain.User) (uint, int, error)
	SignIn(user domain.User) (domain.User, int, error)

	// Task Management
	ViewAllTask(userid string) ([]domain.Task, int, error)
	AddTask(userid string, taskData domain.Task) (domain.Task, int, error)
	EditTask(userid string, taskData domain.Task) (int, error)
	DeleteTask(userid string, taskid string) (int, error)
}
