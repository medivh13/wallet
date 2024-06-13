package route

import (
	"net/http"

	handlers "wallet/src/interface/rest/handlers/balance"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func BalanceRouter(h handlers.BalanceHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/balance_topup", h.TopUp)
	r.Get("/balance_read", h.Get)

	return r
}
