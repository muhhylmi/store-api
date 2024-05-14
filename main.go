package main

import (
	"fmt"
	"net/http"

	app "go-store-api/bin"
	"go-store-api/controller"
	"go-store-api/utils/config"
	"go-store-api/utils/databases"
	"go-store-api/utils/logger"
	"go-store-api/utils/middleware"

	_ "github.com/lib/pq"

	"go-store-api/repository"
	"go-store-api/service"

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

	// category domain
	categoryRepo := repository.NewCategoryRepository(logger, dbService)
	categoryService := service.NewCategoryService(logger, categoryRepo, validate)
	categoryController := controller.NewCategoryController(logger, categoryService)
	// product domain
	productRepository := repository.NewProductRepository(logger, dbService)
	productService := service.NewProductService(logger, productRepository, categoryRepo, validate)
	productController := controller.NewProductController(logger, productService)
	// user domain
	userRepository := repository.NewUserRepository(logger, dbService)
	userService := service.NewUserService(logger, config, userRepository, validate)
	userController := controller.NewUserController(logger, userService)
	// cart domain
	cartRepo := repository.NewShoppingCartRepository(logger, dbService)
	cartService := service.NewShoppingCartService(logger, cartRepo, productRepository, userRepository, validate)
	cartController := controller.NewShoppingCartController(logger, cartService)

	// running application
	router := app.NewRouter(productController, userController, categoryController, cartController)
	server := http.Server{
		Addr:    fmt.Sprintf("%1s:%2s", config.HOST, config.PORT),
		Handler: middleware.NewAuthMiddleware(router, config),
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
