package submissions

import "net/http"

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	// have to change later
	mux.Handle(
		"POST /submissions/hook",
		http.HandlerFunc(h.Hook),
	)

	mux.Handle(
		"POST /submissions/init",
		http.HandlerFunc(h.InitRepository),
	)
}
