package controller

import (
	"net/http"

	"go-store-api/service"
	"go-store-api/utils/logger"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	Logger  *logger.Logger
	Service service.CategoryService
}

type CategoryController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	List(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

func NewCategoryController(logger *logger.Logger, category service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		Logger:  logger,
		Service: category,
	}
}
