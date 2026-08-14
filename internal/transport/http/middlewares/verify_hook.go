package middlewares

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"strings"

	"github.com/labib0x9/sockforces/internal/utils"
)

func (m *Middlewares) HookVerification(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sig := r.Header.Get("X-Hub-Signature-256")
		if sig == "" || !strings.HasPrefix(sig, "sha256") {
			utils.SendError(w, "github signature header invalid", http.StatusBadRequest)
			return
		}
		body, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			utils.SendError(w, "internal server error", http.StatusInternalServerError)
			return
		}
		if !utils.VerifySignature([]byte(sig), body, []byte(m.Cnf.Github.HookSecret)) {
			utils.SendError(w, "invalid signature", http.StatusUnprocessableEntity)
			return
		}

		event := r.Header.Get("X-GitHub-Event")
		if event != "push" {
			slog.Info("Webhook verification middleware", "event", event)
			utils.SendError(w, "we dont want anything other than push event", http.StatusUnprocessableEntity)
			return
		}

		r.Body = io.NopCloser(bytes.NewReader(body))
		next.ServeHTTP(w, r)
	})
}
