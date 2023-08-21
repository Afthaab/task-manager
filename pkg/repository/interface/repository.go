package interfaces

import "github.com/afthaab/task-manager/pkg/domain"

type TaskRepository interface {

	// User Authentication
	FindUserByEmail(userData domain.User) (domain.User, int64)
	DeleteUser(userData domain.User) int64
	CreateUser(userData domain.User) int64
	VerifyOtp(userData domain.User) (domain.User, int64)
	VerifyTheUser(userData domain.User) int64
	FindTheUserById(userid uint) int64

	// Task Management
	GetAllTasks(userid string) ([]domain.Task, int64)
	AddTask(taskData domain.Task) (domain.Task, int64)
}
