package services

import (
	"app/datamodels"
	"app/repositories"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type IUserService interface {
	IsPwdSuccess(id string, pwd string) (*datamodels.User, bool)
	AddUser(user *datamodels.User) (int64, error)
}

type UserService struct {
	UserRepository repositories.IUserRepository
}

func (u UserService) IsPwdSuccess(name string, pwd string) (user *datamodels.User, isOk bool) {
	user, err := u.UserRepository.Select(name)
	if err != nil {
		fmt.Println(err)
		return &datamodels.User{}, false
	}
	isOk, err = ValidatePassword(pwd, user.HashPassword)
	if !isOk {
		fmt.Println(err)
		return &datamodels.User{}, false
	}
	return user, true
}

func ValidatePassword(userPassword, hashPassword string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(userPassword)); err != nil {
		return false, err
	}
	return true, nil
}

func (u UserService) AddUser(user *datamodels.User) (int64, error) {
	pwdByte, errPwd := GeneratePassword(user.HashPassword)
	if errPwd != nil {
		return 0, errPwd
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.Insert(user)
}

func GeneratePassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func NewUserService(repository repositories.IUserRepository) IUserService {
	return &UserService{repository}
}
