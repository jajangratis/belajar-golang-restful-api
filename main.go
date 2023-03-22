package main

import (
	_ "github.com/go-sql-driver/mysql"
	"jajangratis/belajar-golang-restful-api/helper"
	"jajangratis/belajar-golang-restful-api/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8000",
		Handler: authMiddleware,
	}
}

func main() {
	//db := app.NewDb()
	//validate := validator.New()
	//categoryRepository := repository.NewCategoryRepositoryImpl()
	//categoryService := service.NewCategoryService(categoryRepository, db, validate)
	//categoryController := controller.NewCategoryController(categoryService)
	//
	//router := app.NewRouter(categoryController)
	//
	//authMiddleware := middleware.NewAuthMiddleware(router)
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
