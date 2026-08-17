package submissions

import (
	"context"
	"log/slog"
	"strings"

	"github.com/google/uuid"
	"github.com/labib0x9/sockforces/internal/domain/queue"
	"github.com/labib0x9/sockforces/internal/domain/submissions"
)

func (s *service) Publish(ctx context.Context, req submissions.PushEvent) error {
	if req.Before == "0000000000000000000000000000000000000000" {
		slog.Info("", "Initial commit", "no need to publish")
		return nil
	}
	ubmissionId := uuid.New().String()

	msg := queue.PushMessage{
		Id:       ubmissionId,
		RepoName: req.Repository.Name,
		Username: getUserName(req.Repository.Name),
	}
	_ = msg
	return s.queue.Publish(ctx, msg)
}

func getUserName(repo string) string {
	repo = strings.TrimPrefix(repo, "Submission_")
	return strings.SplitN(repo, "_", 2)[0]
}
