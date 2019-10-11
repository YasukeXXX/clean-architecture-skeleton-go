package repositories

import (
	"github.com/YasukeXXX/clean-architecture-skeleton-go/usecases"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID             uint      `gorm:"not null"`
	Name           string    `gorm:"not null"`
	Email          string    `gorm:"not null"`
	Password       string    `gorm:"-"`
	PasswordDigest string    `gorm:"not null`
	CreatedAt      time.Time `gorm:"not null`
	UpdatedAt      time.Time `gorm:"not null`
}

func (u *User) BeforeCreate() (err error) {
	now := time.Now()
	passwordDigest, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	u.PasswordDigest = string(passwordDigest)
	u.CreatedAt = now
	u.UpdatedAt = now
	return
}

func (u *User) BeforeUpdate() (err error) {
	u.UpdatedAt = time.Now()
	return
}

type UserRepository interface {
	Find(id uint) (User, error)
	FindAll() ([]User, error)
	Create(usecases.User) (User, error)
}
