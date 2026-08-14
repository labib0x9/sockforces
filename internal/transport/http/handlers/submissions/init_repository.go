package submissions

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strings"

	"github.com/labib0x9/sockforces/internal/domain/submissions"
	"github.com/labib0x9/sockforces/internal/utils"
)

type initRequest struct {
	Username string `json:"github_username" validate:"required"`
	Labid    string `json:"labid"`
}

func (h *Handler) InitRepository(w http.ResponseWriter, r *http.Request) {
	var req initRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendError(w, "Bad request", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(req); err != nil {
		utils.SendError(w, "field required", http.StatusUnprocessableEntity)
		return
	}

	res, err := h.srv.InitRepository(r.Context(), req.Username, req.Labid)
	if err != nil {
		slog.Error("Service Failed", "error", err)
		switch {
		case strings.Compare(err.Error(), submissions.ErrCreateRepoFailed.Error()) == 0:
			utils.SendError(w, "repo already exists", http.StatusConflict)
		default:
			utils.SendError(w, "internal server error", http.StatusInternalServerError)
		}
		return
	}

	utils.SendJson(w, map[string]any{
		"repo_url": res.RepoUrl,
		"msg":      "created",
	},
		http.StatusCreated,
	)
}
