package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func (controller *ShoppingCartControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartCreateRequest := web.ShopingCartCreateRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
	}
	wrapper.ReadJsonFromRequest(request, &cartCreateRequest)

	categoryResponse := controller.Service.Create(request.Context(), cartCreateRequest)
	webResponse := web.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   categoryResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

// func (controller *ShoppingCartControllerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	categoryResponse := controller.Service.FindAll(request.Context())
// 	webResponse := web.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   categoryResponse,
// 	}

// 	wrapper.WriteToResponseBody(writer, webResponse)
// }
