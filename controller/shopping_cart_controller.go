package controller

import (
	"net/http"

	"go-store-api/service"
	"go-store-api/utils/logger"

	"github.com/julienschmidt/httprouter"
)

type ShoppingCartControllerImpl struct {
	Logger  *logger.Logger
	Service service.ShoppingCartService
}

type ShoppingCartController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	List(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Checkout(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

func NewShoppingCartController(logger *logger.Logger, cart service.ShoppingCartService) ShoppingCartController {
	return &ShoppingCartControllerImpl{
		Logger:  logger,
		Service: cart,
	}
}
