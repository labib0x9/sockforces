package provider

import "context"

type GithubRepo interface {
	CreateRepository(ctx context.Context, templateRepo, repo string) error
	AddCollaborator(ctx context.Context, repo, username string) error
}
