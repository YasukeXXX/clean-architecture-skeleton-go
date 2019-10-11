package usecases

type User struct {
	Name     string
	Email    string
	Password string
}

type FindUserOutput struct {
	Users []User
}

type CreateUserOutput struct {
	User User
}

type CreateUserInput struct {
	User User
}

type UserUsecase interface {
	Find() (FindUserOutput, error)
	Create(CreateUserInput) (CreateUserOutput, error)
}
