package routes

import (
	"encoding/json"
	"net/http"
)

func SetupHealthRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"status":  "ok",
			"service": "user-service",
			"message": "User service is healthy",
		}

		json.NewEncoder(w).Encode(response)
	})
}
