package routes

import (
	"net/http"

	"ticket-booking/user-service/internal/handler"
)

func SetupRoutes(httpHandler *handler.HTTPHandler) *http.ServeMux {
	mux := http.NewServeMux()

	SetupHealthRoutes(mux)

	SetupAuthRoutes(mux, httpHandler)

	return mux
}
