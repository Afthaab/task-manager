package interfaces

import "github.com/afthaab/task-manager/pkg/domain"

type TaskRepository interface {
	FindUserByEmail(userData domain.User) (domain.User, int64)
	DeleteUser(userData domain.User) int64
	CreateUser(userData domain.User) int64
	VerifyOtp(userData domain.User) (domain.User, int64)
	VerifyTheUser(userData domain.User) int64
}
