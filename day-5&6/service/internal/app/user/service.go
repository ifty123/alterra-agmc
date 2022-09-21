package user

import (
	"context"
	"day2-task1/internal/dto"
	"day2-task1/internal/factory"
	mid "day2-task1/internal/middleware"
	"day2-task1/internal/repository"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserUseCase interface {
	GetUserControllers(ctx context.Context) interface{}
	CreateUserController(ctx context.Context, u *dto.User) error
	UpdateUserController(ctx context.Context, u *dto.User, id int) error
	GetUserById(ctx context.Context, id string) (interface{}, error)
	DeleteUserById(ctc context.Context, id int) error
	LoginController(ctx context.Context, u *dto.User) (string, error)
}

type UserController struct {
	UserRepo repository.UserDatabase
}

func NewUserController(f *factory.Factory) *UserController {
	return &UserController{
		UserRepo: f.UserRepository,
	}
}

func (uc *UserController) GetUserControllers(ctx context.Context) interface{} {
	users, e := uc.UserRepo.GetUsers()

	if e != nil {
		return nil
	} else {
		return users
	}
}

func (uc *UserController) CreateUserController(ctx context.Context, u *dto.User) error {

	//change into DTO
	userDTO := dto.ToUserDTO(u)

	//change to models
	userModels := dto.AssembUserDTO(userDTO)

	//save user
	if err := uc.UserRepo.SaveUser(userModels); err != nil {
		return err
	}

	return nil
}

func (uc *UserController) UpdateUserController(ctx context.Context, u *dto.User, id int) error {

	//change into DTO
	userDTO := dto.ToUserDTO(u)

	//change to models
	userModels := dto.AssembUserDTO(userDTO)

	//save user
	if err := uc.UserRepo.UpdateUsers(userModels, id); err != nil {
		return err
	}
	return nil
}

func (uc *UserController) GetUserById(ctx context.Context, id string) (interface{}, error) {

	//get user
	getUser, err := uc.UserRepo.GetUserById(id)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return getUser, nil
}

func (uc *UserController) DeleteUserById(c context.Context, id int) error {

	//get user
	err := uc.UserRepo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UserController) LoginController(ctx context.Context, u *dto.User) (string, error) {

	//change into DTO
	userDTO := dto.ToUserDTO(u)

	//change to models
	userModels := dto.AssembUserDTO(userDTO)

	login, err := uc.UserRepo.LoginUser(userModels)
	if err != nil || login == nil {
		return "", errors.New("Cant login, please check username / password")
	}

	//create token
	token, errToken := mid.CreateToken(strconv.Itoa(userDTO.Id))
	if errToken != nil {
		return "", errToken
	}

	return token, nil
}
