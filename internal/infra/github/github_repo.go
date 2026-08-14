package github

import (
	"context"
	"errors"

	"github.com/google/go-github/v62/github"
	"github.com/labib0x9/sockforces/config"
	"github.com/labib0x9/sockforces/internal/domain/provider"
)

type githubRepo struct {
	client *Client
	org    string
}

func NewGithubRepo(client *Client, cnf *config.Github) provider.GithubRepo {
	return &githubRepo{client: client, org: cnf.Org}
}

func (g *githubRepo) CreateRepository(ctx context.Context, templateRepo, repo string) (string, error) {
	trepo, _, err := g.client.Repositories.CreateFromTemplate(
		ctx,
		g.org,
		templateRepo,
		&github.TemplateRepoRequest{
			Name:    new(repo),
			Owner:   new(g.org),
			Private: new(false),
		},
	)
	if err != nil {
		return "", err
	}
	return trepo.GetCloneURL(), nil
}

func (g *githubRepo) AddCollaborator(ctx context.Context, repo, username string) error {
	_, _, err := g.client.Repositories.AddCollaborator(
		ctx,
		g.org,
		repo,
		username,
		&github.RepositoryAddCollaboratorOptions{}, // default value is push
	)
	return err
}

func (g *githubRepo) GetTemplateRepository(labid string) string {
	return "tcp-echo-server-template"
}

func (g *githubRepo) GetErrorMsg(err error) string {
	var ghErr *github.ErrorResponse
	if errors.As(err, &ghErr) {
		return ghErr.Message
	}
	return ""
}
