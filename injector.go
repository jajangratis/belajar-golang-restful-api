//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"jajangratis/belajar-golang-restful-api/app"
	"jajangratis/belajar-golang-restful-api/controller"
	"jajangratis/belajar-golang-restful-api/middleware"
	"jajangratis/belajar-golang-restful-api/model/repository"
	"jajangratis/belajar-golang-restful-api/service"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDb,
		validator.New,
		service.NewCategoryService,
		controller.NewCategoryController,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
