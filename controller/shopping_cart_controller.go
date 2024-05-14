package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/service"
	"github.com/muhhylmi/store-api/utils/logger"
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
