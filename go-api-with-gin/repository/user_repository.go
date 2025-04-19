package repository

import (
	"go-api-with-gin/model"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	FindAll() ([]model.User, error)
	FindById(id int64) (*model.User, error)
	FindByUsername(username string) (*model.User, error)
	Create(user *model.User) error
	Update(id int64, user *model.User) error
	DeleteById(id int64) error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) FindAll() ([]model.User, error) {
	var users []model.User
	err := userRepository.db.Find(&users).Error
	return users, err
}

func (userRepository *userRepository) FindById(id int64) (*model.User, error) {
	var user model.User
	err := userRepository.db.Where("user_id = ?", id).First(&user).Error
	return &user, err
}

func (userRepository *userRepository) FindByUsername(username string) (*model.User, error) {
	var user model.User
	err := userRepository.db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func (userRepository *userRepository) Create(user *model.User) error {
	return userRepository.db.Create(user).Error
}

func (userRepository *userRepository) Update(id int64, user *model.User) error {
	return userRepository.db.Where("user_id = ?", id).Updates(user).Error
}

func (userRepository *userRepository) DeleteById(id int64) error {
	return userRepository.db.Where("user_id = ?", id).Delete(&model.User{}).Error
}
