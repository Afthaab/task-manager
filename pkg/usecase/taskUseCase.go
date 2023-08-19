package usecase

import (
	"errors"
	"net/http"

	"github.com/afthaab/task-manager/pkg/domain"
	interfaces "github.com/afthaab/task-manager/pkg/repository/interface"
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
	"github.com/afthaab/task-manager/pkg/utility"
)

type TaskUseCase struct {
	taskRepo interfaces.TaskRepository
}

func (u *TaskUseCase) Signup(userData domain.User) (string, int, error) {
	err := utility.ValidateUser(userData)
	if err != nil {
		return userData.Email, http.StatusBadRequest, err
	}

	// searching if the user already exists
	userDetails, rows := u.taskRepo.FindUserByEmail(userData)
	if rows != 0 {
		// Deleting the user if he is not verified
		if userDetails.Isverified == false {
			rows := u.taskRepo.DeleteUser(userDetails)
			if rows == 0 {
				return userData.Email, http.StatusConflict, errors.New("Could not delete the unauthorized user")
			}
		} else {
			return userData.Email, http.StatusBadRequest, errors.New("User Email already exists")
		}
	}

	// Generating the OTP
	otp := utility.GenerateOtp(userData.Email)
	userData.Otp = otp

	// Hashing the Password
	userData.Password = utility.HashPassword(userData.Password)

	// Add Defualt Profile Picture
	userData.Profile = "https://images-for-deliveryapp.s3.ap-south-1.amazonaws.com/kindpng_786207.png"

	// Creating the user
	rows = u.taskRepo.CreateUser(userData)
	if rows == 0 {
		return userData.Email, http.StatusInternalServerError, err
	}

	return userData.Email, http.StatusOK, nil
}

func (u *TaskUseCase) UserVerification(userData domain.User) (uint, int, error) {
	// searching user with the otp and email
	userData, rows := u.taskRepo.VerifyOtp(userData)
	if rows == 0 {
		return userData.Id, http.StatusNotFound, errors.New("Provided OTP is wrong")
	}

	rows = u.taskRepo.VerifyTheUser(userData)
	if rows == 0 {
		return userData.Id, http.StatusInternalServerError, errors.New("Could not verify the user")
	}

	return userData.Id, http.StatusOK, nil

}

func NewTaskUseCase(TaskRepo interfaces.TaskRepository) services.TaskUseCase {
	return &TaskUseCase{
		taskRepo: TaskRepo,
	}
}
