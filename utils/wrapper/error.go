package wrapper

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/muhhylmi/store-api/model/web"
)

type NotFoundError struct {
	Error string
}

type StatusConflictError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}

func NewValidationError(error error) validator.ValidationErrors {
	return validator.ValidationErrors{}
}

func NewStatuConflictError(error string) StatusConflictError {
	return StatusConflictError{Error: error}
}

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if NotFoundDataError(writer, request, err) {
		return
	}

	if ValidationError(writer, request, err) {
		return
	}

	if DataConflictError(writer, request, err) {
		return
	}

	InternalServerError(writer, request, err.(error))
}

func ValidationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	var message string
	if ok {
		for _, err := range exception {
			fieldName := err.Field()
			tagName := err.Tag()
			paramValue := err.Param()
			message = fmt.Sprintf("Field '%s' failed validation for tag '%s' with parameter '%s'", fieldName, tagName, paramValue)
			// Anda dapat menggunakan informasi ini untuk melakukan penanganan error lebih lanjut
		}

		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST ERROR",
			Data:   message,
		}

		WriteToResponseBody(writer, webResponse)
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

		WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func DataConflictError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(StatusConflictError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code:   http.StatusConflict,
			Status: "CONFLICT DATA ERROR",
			Data:   exception.Error,
		}

		WriteToResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}

}

func InternalServerError(writer http.ResponseWriter, request *http.Request, err error) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)
	message := err.Error()
	if message == "EOF" {
		message = "error in reading input"
	}

	webResponse := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL SERVER ERROR",
		Data:   message,
	}

	WriteToResponseBody(writer, webResponse)
}
