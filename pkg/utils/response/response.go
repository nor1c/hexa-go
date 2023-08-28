package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	errResponse := ErrorResponse{
		Error: err.Error(),
	}
	json.NewEncoder(w).Encode(errResponse)
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(SuccessResponse{
		Success: true,
		Data:    data,
	})
}
