package usecases

import (
	"github.com/YasukeXXX/clean-architecture-skeleton-go/models"
)

type FindUserOutput struct {
	Users []models.User
}

type CreateUserOutput struct {
	User models.User
}

type CreateUserInput struct {
	User models.User
}

type UserUsecase interface {
	Find() (FindUserOutput, error)
	Create(CreateUserInput) (CreateUserOutput, error)
}
