package controller

import (
	"net/http"

	"go-store-api/model/web"
	"go-store-api/service"
	"go-store-api/utils/logger"
	"go-store-api/utils/wrapper"

	"github.com/julienschmidt/httprouter"
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

func (controller *ProductControllerImpl) List(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	productListRequest := web.ProductListRequest{
		CategoryId: query.Get("category_id"),
		Keyword:    query.Get("q"),
	}
	ProductResponses := controller.ProductService.FindAll(request.Context(), productListRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponses,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
