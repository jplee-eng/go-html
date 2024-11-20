package handlers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Message: "active",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
