package test

import (
	"database/sql"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"kelas-golang-pzn/go-restful-api/app"
	"kelas-golang-pzn/go-restful-api/controller"
	"kelas-golang-pzn/go-restful-api/helper"
	"kelas-golang-pzn/go-restful-api/middleware"
	"kelas-golang-pzn/go-restful-api/repository"
	"kelas-golang-pzn/go-restful-api/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/go_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter() http.Handler {
	db := setupTestDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	router := setupRouter()

	requestBody := strings.NewReader(`{"name": "Laptop"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	fmt.Println(response)
	assert.Equal(t, 200, response.StatusCode)
}

func TestCreateCategoryFailed(t *testing.T) {
}

func TestUpdateCategorySuccess(t *testing.T) {
}

func TestUpdateCategoryFailed(t *testing.T) {
}

func TestGetCategorySuccess(t *testing.T) {
}

func TestGetCategoryFailed(t *testing.T) {
}

func TestDeleteCategorySuccess(t *testing.T) {
}

func TestDeleteCategoryFailed(t *testing.T) {
}

func TestListCategoriesSuccess(t *testing.T) {
}

func TestUnthorized(t *testing.T) {
}
