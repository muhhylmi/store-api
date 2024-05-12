package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/service"
	"github.com/muhhylmi/store-api/utils/logger"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

type ProductControllerImpl struct {
	Logger         *logger.Logger
	ProductService service.ProductService
}

func NewProductController(logger *logger.Logger, ProductService service.ProductService) ProductController {
	return &ProductControllerImpl{
		Logger:         logger,
		ProductService: ProductService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductCreateRequest := web.ProductCreateRequest{
		AuthData: web.AuthData{
			Role:   request.Header.Get("role"),
			UserId: request.Header.Get("userId"),
		},
	}
	wrapper.ReadJsonFromRequest(request, &ProductCreateRequest)

	ProductResponse := controller.ProductService.Create(request.Context(), ProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

// func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	ProductUpdateRequest := web.ProductUpdateRequest{}
// 	wrapper.ReadJsonFromRequest(request, &ProductUpdateRequest)

// 	productId := params.ByName("productId")
// 	ProductUpdateRequest.Id = productId

// 	ProductResponse := controller.ProductService.Update(request.Context(), ProductUpdateRequest)
// 	webResponse := web.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   ProductResponse,
// 	}

// 	wrapper.WriteToResponseBody(writer, webResponse)
// }

// func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	ProductId := params.ByName("productId")

// 	controller.ProductService.Delete(request.Context(), productId)
// 	webResponse := web.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 	}

// 	wrapper.WriteToResponseBody(writer, webResponse)
// }

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	productId := params.ByName("productId")

	ProductResponse := controller.ProductService.FindById(request.Context(), productId)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}

// func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
// 	ProductResponses := controller.ProductService.FindAll(request.Context())
// 	webResponse := web.WebResponse{
// 		Code:   200,
// 		Status: "OK",
// 		Data:   ProductResponses,
// 	}

// 	wrapper.WriteToResponseBody(writer, webResponse)
// }
