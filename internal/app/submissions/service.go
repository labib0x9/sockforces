package submissions

import (
	"context"

	"github.com/labib0x9/sockforces/internal/domain/provider"
	"github.com/labib0x9/sockforces/internal/domain/submissions"
)

type Service interface {
	InitRepository(ctx context.Context, username string, labid string) (*submissions.InitResult, error)
}

type service struct {
	gitRepo provider.GithubRepo
}

func NewService(gitRepo provider.GithubRepo) Service {
	return &service{gitRepo: gitRepo}
}
