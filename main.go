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
	l := logger.LogWithContext("main", "init")

	uri := "postgres://postgres:password@localhost:5432/go-store"
	db, err := databases.InitPostgres(&databases.DBServiceVar{
		Logger:      logger,
		PostgresUri: &uri,
	})
	if err != nil {
		l.Error(err)
		panic(err)
	}
	validate := validator.New()

	// product domain
	productRepository := repository.NewCategoryRepository(logger, &databases.DBService{
		Gorm: db,
	})
	productService := service.NewProductService(logger, productRepository, validate)
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
