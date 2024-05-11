package main

import (
	"net/http"

	_ "github.com/lib/pq"
	"github.com/muhhylmi/store-api/app"
	"github.com/muhhylmi/store-api/controller"
	"github.com/muhhylmi/store-api/utils/databases"
	"github.com/muhhylmi/store-api/utils/logger"
	"github.com/muhhylmi/store-api/utils/middleware"

	"github.com/muhhylmi/store-api/repository"
	"github.com/muhhylmi/store-api/service"

	"github.com/go-playground/validator/v10"
)

func main() {
	logger := logger.Newlogger()
	db := databases.NewDB(logger)
	validate := validator.New()
	l := logger.LogWithContext("main", "init")

	// product domain
	productRepository := repository.NewCategoryRepository(logger)
	productService := service.NewProductService(logger, productRepository, db, validate)
	productController := controller.NewProductController(logger, productService)

	router := app.NewRouter(productController)
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Error(err)
		}
	}()

	l.Info("Starting HTTP server on port 3000")
	select {}
}
