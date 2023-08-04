package movie_handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  bool        `json:"status"`
	Message int         `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(w http.ResponseWriter, status bool, code int, data interface{}) {
	response := Response{
		Status:  status,
		Message: code,
		Data:    data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(response)
}

func BuldErrorResponse(w http.ResponseWriter, status bool, code int, message string) {
	BuildResponse(w, status, code, map[string]string{"error": message})
}
