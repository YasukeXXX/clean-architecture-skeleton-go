package interactors

import (
	"github.com/YasukeXXX/clean-architecture-skeleton-go/repositories"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/usecases"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/models"
)

type UserInteractor struct {
	UserRepository repositories.UserRepository
}

func (i UserInteractor) Find() (output usecases.FindUserOutput, err error) {
	users, err := i.UserRepository.FindAll()
	if err != nil {
		return
	}
	for _, user := range users {
		userOutput := models.User{Name: user.Name, Email: user.Email}
		output.Users = append(output.Users, userOutput)
	}
	return
}

func (i UserInteractor) Create(input usecases.CreateUserInput) (output usecases.CreateUserOutput, err error) {
	dto, err := i.UserRepository.Create(input.User)
	if err != nil {
		return
	}
	user := models.User{Name: dto.Name, Email: dto.Email}
	output = usecases.CreateUserOutput{user}
	return
}
