package controller

import (
	"net/http"

	"go-store-api/model/web"
	"go-store-api/utils/wrapper"

	"github.com/julienschmidt/httprouter"
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

func (controller *ShoppingCartControllerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	status := "PENDING"
	if query.Get("status") != "" {
		status = query.Get("status")
	}

	listCartRequest := web.ListCartRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
		Status: status,
	}

	categoryResponse := controller.Service.FindAll(request.Context(), listCartRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   categoryResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartId := params.ByName("shopping_cart_id")

	updateCartRequest := web.UpdateCartRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
		ShoppingCartId: cartId,
	}
	wrapper.ReadJsonFromRequest(request, &updateCartRequest)

	response := controller.Service.Update(request.Context(), updateCartRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	cartId := params.ByName("shopping_cart_id")

	deleteCartRequest := web.DeleteCartRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
		ShoppingCartId: cartId,
	}

	response := controller.Service.Delete(request.Context(), deleteCartRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

func (controller *ShoppingCartControllerImpl) Checkout(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	checkoutRequest := web.CheckoutCartRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
	}
	wrapper.ReadJsonFromRequest(request, &checkoutRequest)

	response := controller.Service.Checkout(request.Context(), checkoutRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   response,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
