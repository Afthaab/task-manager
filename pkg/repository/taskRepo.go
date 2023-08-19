package repository

import (
	"github.com/afthaab/task-manager/pkg/domain"
	interfaces "github.com/afthaab/task-manager/pkg/repository/interface"
	"gorm.io/gorm"
)

type TaskDatabase struct {
	db *gorm.DB
}

func (r *TaskDatabase) VerifyTheUser(userData domain.User) int64 {
	result := r.db.Exec("update users set isverified = true where email LIKE ? AND otp = ?", userData.Email, userData.Otp)
	return result.RowsAffected
}

func (r *TaskDatabase) VerifyOtp(userData domain.User) (domain.User, int64) {
	result := r.db.Raw("SELECT * FROM users WHERE email LIKE ? AND otp = ?", userData.Email, userData.Otp).Scan(&userData)
	return userData, result.RowsAffected
}

func (r *TaskDatabase) DeleteUser(userData domain.User) int64 {
	result := r.db.Exec("DELETE FROM users WHERE email LIKE ?", userData.Email)
	return result.RowsAffected
}

func (r *TaskDatabase) FindUserByEmail(userData domain.User) (domain.User, int64) {
	result := r.db.Raw("select * from users where email LIKE ?", userData.Email).Scan(&userData)
	return userData, result.RowsAffected
}

func (r *TaskDatabase) CreateUser(userData domain.User) int64 {
	result := r.db.Create(&userData)
	return result.RowsAffected
}

func NewTaskRepository(DB *gorm.DB) interfaces.TaskRepository {
	return &TaskDatabase{
		db: DB,
	}
}
