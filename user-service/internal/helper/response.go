package helper

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}

func WriteSuccess(w http.ResponseWriter, message string, data interface{}) error {
	response := Response{
		Success: true,
		Message: message,
		Data:    data,
	}
	return WriteJSON(w, http.StatusOK, response)
}

func WriteError(w http.ResponseWriter, statusCode int, message string, err error) error {
	response := Response{
		Success: false,
		Message: message,
	}

	if err != nil {
		response.Error = err.Error()
	}

	return WriteJSON(w, statusCode, response)
}

func WriteBadRequest(w http.ResponseWriter, message string, err error) error {
	return WriteError(w, http.StatusBadRequest, message, err)
}

func WriteUnauthorized(w http.ResponseWriter, message string, err error) error {
	return WriteError(w, http.StatusUnauthorized, message, err)
}

func WriteInternalError(w http.ResponseWriter, message string, err error) error {
	return WriteError(w, http.StatusInternalServerError, message, err)
}
