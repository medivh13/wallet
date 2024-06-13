package route

import (
	"net/http"

	handlers "wallet/src/interface/rest/handlers/health"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func HealthRouter(h handlers.IHealthHandler) http.Handler {
	r := chi.NewRouter()

	r.Get("/ping", h.Ping)

	return r
}
