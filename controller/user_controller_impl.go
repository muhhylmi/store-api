package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func (controller *UserControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := web.UserCreateRequest{}
	wrapper.ReadJsonFromRequest(request, &userCreateRequest)

	ProductResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserControllerImpl) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginRequest{}
	wrapper.ReadJsonFromRequest(request, &loginRequest)

	ProductResponse := controller.UserService.Login(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
