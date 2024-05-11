package helper

import (
	"encoding/json"
	"net/http"
)

func ReadJsonFromRequest(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	PanicIfError(err)
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	enconder := json.NewEncoder(writer)
	err := enconder.Encode(response)
	PanicIfError(err)
}
