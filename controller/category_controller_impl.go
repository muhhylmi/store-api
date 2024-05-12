package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
	}
	wrapper.ReadJsonFromRequest(request, &categoryCreateRequest)

	categoryResponse := controller.Service.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   categoryResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (controller *CategoryControllerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryResponse := controller.Service.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
