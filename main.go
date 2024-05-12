package main

import (
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/muhhylmi/store-api/app"
	"github.com/muhhylmi/store-api/controller"
	"github.com/muhhylmi/store-api/utils/config"
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
	config := config.GetConfig()
	validate := validator.New()

	db, err := databases.InitPostgres(&databases.DBServiceVar{
		Logger:      logger,
		PostgresUri: &config.DB_URI,
	})
	if err != nil {
		l.Error(err)
		panic(err)
	}
	dbService := &databases.DBService{
		Gorm: db,
	}

	// product domain
	productRepository := repository.NewCategoryRepository(logger, dbService)
	productService := service.NewProductService(logger, productRepository, validate)
	productController := controller.NewProductController(logger, productService)

	// user domain
	userRepository := repository.NewUserRepository(logger, dbService)
	userService := service.NewUserService(logger, config, userRepository, validate)
	userController := controller.NewUserController(logger, userService)

	router := app.NewRouter(productController, userController)
	server := http.Server{
		Addr:    fmt.Sprintf("%1s:%2s", config.HOST, config.PORT),
		Handler: middleware.NewAuthMiddleware(router),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			l.Error(err)
		}
	}()

	l.Info("Starting HTTP server on port ", config.PORT)
	select {}
}
