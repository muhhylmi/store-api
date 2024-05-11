package controller

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/muhhylmi/store-api/helper"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/service"
)

type ProductControllerImpl struct {
	ProductService service.ProductService
}

func NewProductController(ProductService service.ProductService) ProductController {
	return &ProductControllerImpl{
		ProductService: ProductService,
	}
}

func (controller *ProductControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductCreateRequest := web.ProductCreateRequest{}
	helper.ReadJsonFromRequest(request, &ProductCreateRequest)

	ProductResponse := controller.ProductService.Create(request.Context(), ProductCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductUpdateRequest := web.ProductUpdateRequest{}
	helper.ReadJsonFromRequest(request, &ProductUpdateRequest)

	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)
	ProductUpdateRequest.Id = id

	ProductResponse := controller.ProductService.Update(request.Context(), ProductUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	controller.ProductService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductId := params.ByName("ProductId")
	id, err := strconv.Atoi(ProductId)
	helper.PanicIfError(err)

	ProductResponse := controller.ProductService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ProductControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	ProductResponses := controller.ProductService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   ProductResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
