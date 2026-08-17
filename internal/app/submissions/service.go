package submissions

import (
	"context"

	"github.com/labib0x9/sockforces/internal/domain/provider"
	"github.com/labib0x9/sockforces/internal/domain/queue"
	"github.com/labib0x9/sockforces/internal/domain/submissions"
)

type Service interface {
	InitRepository(ctx context.Context, username string, labid string) (*submissions.InitResult, error)
	Clone(req submissions.PushEvent) error
	Publish(ctx context.Context, msg submissions.PushEvent) error
}

type service struct {
	gitRepo provider.GithubRepo
	queue   queue.Queue
}

func NewService(gitRepo provider.GithubRepo, queue queue.Queue) Service {
	return &service{gitRepo: gitRepo, queue: queue}
}
