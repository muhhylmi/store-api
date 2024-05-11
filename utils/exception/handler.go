package exception

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/muhhylmi/store-api/model/web"
	"github.com/muhhylmi/store-api/utils/wrapper"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if NotFoundDataError(writer, request, err) {
		return
	}

	if ValidationError(writer, request, err) {
		return
	}

	InternalServerError(writer, request, err)
}

func ValidationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST ERROR",
			Data:   exception.Error(),
		}

		wrapper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func NotFoundDataError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND ERROR",
			Data:   exception.Error,
		}

		wrapper.WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   err,
	}

	wrapper.WriteToResponseBody(writer, webResponse)
}
