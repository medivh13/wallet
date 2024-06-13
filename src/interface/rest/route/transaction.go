package route

import (
	"net/http"

	handlers "wallet/src/interface/rest/handlers/transaction"

	"github.com/go-chi/chi/v5"
)

// HealthRouter a completely separate router for health check routes
func TransactionRouter(h handlers.TransactionHandlerInterface) http.Handler {
	r := chi.NewRouter()

	r.Post("/transfer", h.Transfer)
	r.Get("/top_transactions_per_user", h.GetTopTen)
	r.Get("/top_users", h.GetOverallTopTransactions)
	return r
}
