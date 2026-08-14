package provider

import "context"

type GithubRepo interface {
	CreateRepository(ctx context.Context, templateRepo, repo string) (string, error)
	AddCollaborator(ctx context.Context, repo, username string) error
	GetTemplateRepository(labid string) string
	GetErrorMsg(err error) string
}
