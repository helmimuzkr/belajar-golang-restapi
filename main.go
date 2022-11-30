package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/helmimuzkr/belajar-golang-restapi/config"
	"github.com/helmimuzkr/belajar-golang-restapi/controller"
	"github.com/helmimuzkr/belajar-golang-restapi/middleware"
	"github.com/helmimuzkr/belajar-golang-restapi/repository"
	"github.com/helmimuzkr/belajar-golang-restapi/route"
	"github.com/helmimuzkr/belajar-golang-restapi/service"
	"github.com/helmimuzkr/belajar-golang-restapi/util"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Get Config
	appConfig := config.Get()

	// Register database
	db := util.NewDatabaseConnection(appConfig)

	// Register Repository
	categoryRepository := repository.NewCategoryRepo(db)

	// Register Service
	validator := validator.New()
	categoryService := service.NewCategoryService(validator, categoryRepository)

	// Register Controller
	categoryContoller := controller.NewCategoryController(categoryService)

	// Register Router
	router := httprouter.New()
	router = route.NewCategoryRouter(router, categoryContoller).CategoryRouter()

	// Register Middleware
	middleware := middleware.NewAuthApiMiddleware(router, appConfig.ApiKey.ApiKeyCategory)

	fmt.Println("test")

	// running server
	log.Println("running on", appConfig.App.Port)
	err := http.ListenAndServe(appConfig.App.BaseURL, middleware)
	if err != nil {
		log.Fatal(err)
		return
	}
}
