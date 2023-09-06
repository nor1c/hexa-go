package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, message string, err error) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	errResponse := ErrorResponse{
		Success: false,
		Message: message,
		Error:   err,
	}

	json.NewEncoder(w).Encode(errResponse)
}

func WriteSuccessResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	successResponse := SuccessResponse{
		Success: true,
		Data:    data,
	}

	json.NewEncoder(w).Encode(successResponse)
}
