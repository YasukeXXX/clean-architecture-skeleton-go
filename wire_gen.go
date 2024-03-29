// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/YasukeXXX/clean-architecture-skeleton-go/controllers"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/handlers"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/interactors"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/registries"
	"github.com/YasukeXXX/clean-architecture-skeleton-go/repositories/postgres"
)

// Injectors from wire.go:

func InitializeHandler() handlers.AppHandler {
	registriesPostgres := registries.NewPostgres()
	userRepository := &postgres.UserRepository{
		DB: registriesPostgres,
	}
	userInteractor := &interactors.UserInteractor{
		UserRepository: userRepository,
	}
	userController := controllers.UserController{
		Interactor: userInteractor,
	}
	userHandler := handlers.UserHandler{
		Controller: userController,
	}
	appHandler := handlers.AppHandler{
		UserHandler: userHandler,
	}
	return appHandler
}
