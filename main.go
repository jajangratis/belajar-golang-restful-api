package main

import (
	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
	"jajangratis/belajar-golang-restful-api/app"
	"jajangratis/belajar-golang-restful-api/controller"
	"jajangratis/belajar-golang-restful-api/helper"
	"jajangratis/belajar-golang-restful-api/middleware"
	"jajangratis/belajar-golang-restful-api/model/repository"
	"jajangratis/belajar-golang-restful-api/service"
	"net/http"
)

func main() {
	db := app.NewDb()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepositoryImpl()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:8000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
