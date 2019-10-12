package postgres

import (
	"github.com/YasukeXXX/clean-architecture-skeleton-go/registries"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/repositories"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/models"
)

type UserRepository struct {
	DB registries.Postgres
}

func (r UserRepository) Find(id uint) (user repositories.User, err error) {
	err = r.DB.Where(&repositories.User{ID: id}).Find(&user).Error
	return
}

func (r UserRepository) FindAll() (users []repositories.User, err error) {
	err = r.DB.Find(&users).Error
	return
}

func (r UserRepository) Create(user models.User) (record repositories.User, err error) {
	record = repositories.User{Name: user.Name, Email: user.Email, Password: user.Password}
	err = r.DB.Create(&record).Error
	return
}
