package usecase

import (
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/afthaab/task-manager/pkg/domain"
	interfaces "github.com/afthaab/task-manager/pkg/repository/interface"
	services "github.com/afthaab/task-manager/pkg/usecase/interface"
	"github.com/afthaab/task-manager/pkg/utility"
	"github.com/golang-jwt/jwt"
)

type TaskUseCase struct {
	taskRepo interfaces.TaskRepository
}

func (u *TaskUseCase) GenerateAccessToken(userData domain.User, role string) (string, error) {
	claims := domain.JWTClaims{
		Userid: userData.Id,
		Role:   role,
		Source: "AccessToken",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().AddDate(1, 0, 0).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessToken, err := token.SignedString([]byte(os.Getenv("SECRET_KET")))

	return accessToken, err
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

func (u *TaskUseCase) UserVerification(userData domain.User) (uint64, int, error) {
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
func (u *TaskUseCase) SignIn(user domain.User) (domain.User, int, error) {
	// checking if user credential is valid ?
	userData, rows := u.taskRepo.FindUserByEmail(user)
	if rows == 0 {
		return userData, http.StatusNotFound, errors.New("Entered email does not exist")
	}

	// checking if the user is Otp verfied
	if userData.Isverified == false {
		rows := u.taskRepo.DeleteUser(userData)
		if rows == 0 {
			return userData, http.StatusConflict, errors.New("Could not delete the authenticated user")
		}
		return userData, http.StatusAccepted, errors.New("User is not Authenticated, SignUp again")
	}

	// checking and comparing the passwords
	if !utility.VerifyPassword(user.Password, userData.Password) {
		return userData, http.StatusOK, errors.New("Passwords did not match")
	}

	return userData, http.StatusOK, nil
}

func NewTaskUseCase(TaskRepo interfaces.TaskRepository) services.TaskUseCase {
	return &TaskUseCase{
		taskRepo: TaskRepo,
	}
}
