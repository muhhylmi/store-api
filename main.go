package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/muhhylmi/store-api/app"
	"github.com/muhhylmi/store-api/controller"
	"github.com/muhhylmi/store-api/helper"
	"github.com/muhhylmi/store-api/middleware"
	"github.com/muhhylmi/store-api/repository"
	"github.com/muhhylmi/store-api/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	db := helper.NewDB()
	validate := validator.New()

	// product domain
	productRepository := repository.NewCategoryRepository()
	productService := service.NewProductService(productRepository, db, validate)
	productController := controller.NewProductController(productService)

	router := app.NewRouter(productController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
