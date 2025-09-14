package service

import (
	"errors"
	"go_task_4/internal/model"
	"go_task_4/internal/repository"
	"go_task_4/pkg/utils"
	"time"
)

type UserService struct {
	userrepository *repository.UserRepository
}

func NewUserService(userrepository *repository.UserRepository) *UserService {
	return &UserService{userrepository: userrepository}
}

func (us *UserService) Register(user *model.User) error {
	if err := utils.CheckNotEmpty("Username", user.Username); err != nil {
		return err
	}
	if err := utils.CheckNotEmpty("Password", user.Password); err != nil {
		return err
	}
	if user.Email != "" {
		if err := utils.CheckEmail("Email", user.Email); err != nil {
			return err
		}
	}
	user.Password = utils.GeneratePassword(user.Password)
	err := us.userrepository.Create(*user)
	return err
}

func (us *UserService) Login(user *model.User) (string, error) {
	if err := utils.CheckNotEmpty("Username", user.Username); err != nil {
		return "", err
	}
	if err := utils.CheckNotEmpty("Password", user.Password); err != nil {
		return "", err
	}
	userdb, err := us.userrepository.GetByUsername(user.Username)
	if err != nil {
		return "", err
	}
	result, err := utils.VerifyPassword(user.Password, userdb.Password)
	if err != nil {
		return "", err
	}
	if !result {
		return "", errors.New("invalid password")
	}
	token, err := utils.GenerateToken(userdb.ID, userdb.Username, time.Hour*24)
	if err != nil {
		return "", err
	}
	return token, nil
}
