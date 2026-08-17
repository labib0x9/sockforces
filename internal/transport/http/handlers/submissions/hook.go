package submissions

import (
	"encoding/json"
	"net/http"

	"github.com/labib0x9/sockforces/internal/domain/submissions"
	"github.com/labib0x9/sockforces/internal/utils"
)

// push events to message queue with everything is needed to run the tests with a unique submission id, which is unique for this event
func (h *Handler) Hook(w http.ResponseWriter, r *http.Request) {
	// body, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	utils.SendError(w, "internal service error", http.StatusInternalServerError)
	// 	return
	// }

	var req submissions.PushEvent
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "bad request", http.StatusBadRequest)
		return
	}

	err := h.srv.Publish(r.Context(), req)
	if err != nil {
		utils.SendError(w, "internal service error", http.StatusInternalServerError)
		return
	}

	utils.SendJson(w, "", http.StatusAccepted)
}
