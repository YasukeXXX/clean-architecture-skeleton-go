package controllers

import (
	"encoding/json"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/models"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/usecases"
	"log"
)

type UserController struct {
	Interactor usecases.UserUsecase
}

func (c UserController) Index() (string, error) {
	users, err := c.Interactor.Find()
	if err != nil {
		log.Print(err)
		return "", err
	}
	res, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(res), nil
}

func (c UserController) Create(name string, email string, password string) (string, error) {
	input := usecases.CreateUserInput{usecases.User{Name: name, Email: email, Password: password}}
	output, err := c.Interactor.Create(input)
	if err != nil {
		log.Print(err)
		return "", err
	}
	user := models.User{Name: output.User.Name, Email: output.User.Email}
	res, err := json.Marshal(user)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(res), nil
}
