// +build wireinject

package main

import (
	"github.com/YasukeXXX/clean-architecture-skeleton-go/controllers"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/handlers"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/interactors"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/registries"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/repositories"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/repositories/postgres"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/usecases"
	"github.com/google/wire"
)

func InitializeHandler() handlers.AppHandler {
	wire.Build(
		registries.NewPostgres,
		wire.Struct(new(postgres.UserRepository), "*"),
		wire.Bind(new(repositories.UserRepository), new(*postgres.UserRepository)),
		wire.Bind(new(usecases.UserUsecase), new(*interactors.UserInteractor)),
		wire.Struct(new(interactors.UserInteractor), "*"),
		wire.Struct(new(controllers.UserController), "*"),
		wire.Struct(new(handlers.UserHandler), "*"),
		wire.Struct(new(handlers.AppHandler), "*"),
	)
	return handlers.AppHandler{}
}
