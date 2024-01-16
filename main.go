package main

import (
	"net/http"

	"kelas-golang-pzn/go-restful-api/app"
	"kelas-golang-pzn/go-restful-api/controller"
	"kelas-golang-pzn/go-restful-api/helper"
	"kelas-golang-pzn/go-restful-api/middleware"
	"kelas-golang-pzn/go-restful-api/repository"
	"kelas-golang-pzn/go-restful-api/service"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
