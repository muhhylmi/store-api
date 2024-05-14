package controller

import (
	"net/http"

	"go-store-api/service"
	"go-store-api/utils/logger"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	Logger      *logger.Logger
	UserService service.UserService
}

type UserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	TopUp(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

func NewUserController(logger *logger.Logger, userService service.UserService) UserController {
	return &UserControllerImpl{
		Logger:      logger,
		UserService: userService,
	}
}
