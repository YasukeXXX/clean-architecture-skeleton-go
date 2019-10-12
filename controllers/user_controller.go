package controllers

import (
	"encoding/json"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/models"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/usecases"
	"log"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(user models.User) User {
	return User{user.Name, user.Email}
}

type UserController struct {
	Interactor usecases.UserUsecase
}

func (c UserController) Index() (string, error) {
	output, err := c.Interactor.Find()
	if err != nil {
		log.Print(err)
		return "", err
	}
	var users []User
	for _, user := range output.Users {
		users = append(users, NewUser(user))
	}
	res, err := json.Marshal(users)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(res), nil
}

func (c UserController) Create(name string, email string, password string) (string, error) {
	input := usecases.CreateUserInput{models.User{Name: name, Email: email, Password: password}}
	output, err := c.Interactor.Create(input)
	if err != nil {
		log.Print(err)
		return "", err
	}
	user := NewUser(output.User)
	res, err := json.Marshal(user)
	if err != nil {
		log.Print(err)
		return "", err
	}
	return string(res), nil
}
