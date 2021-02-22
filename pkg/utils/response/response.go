package response

import (
	"encoding/json"
	"net/http"
)

func ResponseWithError(response http.ResponseWriter, statusCode int, msg string) {
	ResponseWithJSON(response, statusCode, map[string]string{
		"error": msg,
	})
}

func ResponseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
