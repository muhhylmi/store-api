package wrapper

import (
	"encoding/json"
	"net/http"
)

func ReadJsonFromRequest(request *http.Request, result interface{}) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	if err != nil {
		panic(err)
	}
}

func WriteToResponseBody(writer http.ResponseWriter, response interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	enconder := json.NewEncoder(writer)
	err := enconder.Encode(response)
	if err != nil {
		panic(err)
	}
}
