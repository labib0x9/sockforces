package submissions

import (
	"net/http"

	"github.com/labib0x9/sockforces/internal/transport/http/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middlewares.Manager) {
	// have to change later
	mux.Handle(
		"POST /submissions/hook",
		manager.With(
			http.HandlerFunc(h.Hook),
			h.middlewares.HookVerification,
		),
	)

	mux.Handle(
		"POST /submissions/init",

		http.HandlerFunc(h.InitRepository),
	)
}
