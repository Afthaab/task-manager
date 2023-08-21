package usecase

import (
	"errors"
	"net/http"
	"os"
	"strconv"
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

//////////////////////////////////// ---- MIDDLEWARE ---- ////////////////////////////////////

func (r *TaskUseCase) ValidateTheJwtUser(userid uint) (int, error) {
	rows := r.taskRepo.FindTheUserById(userid)
	if rows == 0 {
		return http.StatusUnauthorized, errors.New("User not found in the database")
	} else {
		return http.StatusOK, nil
	}
}

func (u *TaskUseCase) VerifyToken(token string) (bool, *domain.JWTClaims) {
	claims := &domain.JWTClaims{}
	tkn, err := utility.GetTokenFromString(token, claims)
	if err != nil {
		return false, claims
	}
	if tkn.Valid {
		if err := claims.Valid(); err != nil {
			return false, claims
		}
	}
	return true, claims
}

func (u *TaskUseCase) GenerateAccessToken(userData domain.User, role string) (string, error) {
	claims := domain.JWTClaims{
		Userid: userData.Userid,
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

//////////////////////////////////// ---- USER AUTHENTICATION ---- ////////////////////////////////////

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
		return userData.Userid, http.StatusNotFound, errors.New("Provided OTP is wrong")
	}

	rows = u.taskRepo.VerifyTheUser(userData)
	if rows == 0 {
		return userData.Userid, http.StatusInternalServerError, errors.New("Could not verify the user")
	}

	return userData.Userid, http.StatusOK, nil

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

//////////////////////////////////// ---- TASK MANAGEMENT ---- ////////////////////////////////////

func (u *TaskUseCase) ViewAllTask(userid string) ([]domain.Task, int, error) {
	taskDatas, rows := u.taskRepo.GetAllTasks(userid)
	if rows == 0 {
		return taskDatas, http.StatusNotFound, errors.New("Could not find the user with his tasks")
	}
	return taskDatas, http.StatusOK, nil
}

func (u *TaskUseCase) AddTask(userid string, taskData domain.Task) (domain.Task, int, error) {
	uid, _ := strconv.ParseUint(userid, 10, 0)
	taskData.Uid = uint(uid)
	taskData, rows := u.taskRepo.AddTask(taskData)
	if rows == 0 {
		return taskData, http.StatusInternalServerError, errors.New("Could not add the task to the Database")
	}
	return taskData, http.StatusOK, nil
}
func (u *TaskUseCase) EditTask(userid string, taskData domain.Task) (int, error) {
	rows := u.taskRepo.EditTask(userid, taskData)
	if rows == 0 {
		return http.StatusNotFound, errors.New("Could not update the task in the database")
	}
	return http.StatusOK, nil
}

func (u *TaskUseCase) DeleteTask(userid string, taskid string) (int, error) {
	rows := u.taskRepo.DeleteTask(userid, taskid)
	if rows == 0 {
		return http.StatusInternalServerError, errors.New("Could not delete the task")
	}
	return http.StatusOK, nil
}

func NewTaskUseCase(TaskRepo interfaces.TaskRepository) services.TaskUseCase {
	return &TaskUseCase{
		taskRepo: TaskRepo,
	}
}
