package handlers

/*
 * Author      : Jody (jody.almaida@gmail.com)
 * Modifier    :
 * Domain      : wallet
 */
import (
	"net/http"

	"wallet/src/interface/rest/response"
)

type IHealthHandler interface {
	Ping(w http.ResponseWriter, r *http.Request)
}

type healthHandler struct {
	response response.IResponseClient
}

func NewHealthHandler(r response.IResponseClient) IHealthHandler {
	return &healthHandler{
		response: r,
	}
}

func (h *healthHandler) Ping(w http.ResponseWriter, r *http.Request) {
	h.response.JSON(w, "Pong", nil, nil)
}
