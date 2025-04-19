package service

import (
	"go-api-with-gin/dto"
	"go-api-with-gin/model"
	"go-api-with-gin/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

type UserService interface {
	FindAllUsers() ([]model.User, error)
	FindUserById(id int64) (*model.User, error)
	FindUserByUsername(username string) (*model.User, error)
	CreateUser(user dto.CreateUserDTO) error
	UpdateUser(id int64, user dto.UpdateUserDTO) error
	DeleteUserById(id int64) error
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (userRepository *userService) FindAllUsers() ([]model.User, error) {
	return userRepository.userRepository.FindAll()
}

func (userRepository *userService) FindUserById(id int64) (*model.User, error) {
	return userRepository.userRepository.FindById(id)
}

func (userRepository *userService) FindUserByUsername(username string) (*model.User, error) {
	return userRepository.userRepository.FindByUsername(username)
}

func (userRepository *userService) CreateUser(newUser dto.CreateUserDTO) error {
	user := model.User{
		Name:     newUser.Name,
		Email:    newUser.Email,
		Username: newUser.Username,
		Password: newUser.Password, // TODO: Hash password
	}
	return userRepository.userRepository.Create(&user)
}

func (userRepository *userService) UpdateUser(id int64, updatingUser dto.UpdateUserDTO) error {
	user := model.User{
		Name:     updatingUser.Name,
		Email:    updatingUser.Email,
		Password: updatingUser.Password, // TODO: Hash password
	}
	return userRepository.userRepository.Update(id, &user)
}

func (userRepository *userService) DeleteUserById(id int64) error {
	return userRepository.userRepository.DeleteById(id)
}
