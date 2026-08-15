package github

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

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

func (g *githubRepo) getToken() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	return g.client.t.Token(ctx)
}

// id is the submission id, unique for each submission
func (g *githubRepo) CloneRepository(ctx context.Context, fullname string, id string) (string, error) {
	token, err := g.getToken()
	if err != nil {
		return "", err
	}

	repoPath := filepath.Join(os.TempDir(), id)
	cloneURL := fmt.Sprintf("https://x-access-token:%s@github.com/%s.git", token, fullname)

	cmd := exec.CommandContext(ctx, "git", "clone", cloneURL, repoPath)
	out, err := cmd.CombinedOutput()
	if err != nil {
		_ = out
		return "", err
	}

	return repoPath, nil
}

func RemoveClonedRepository(path string) error {
	return os.RemoveAll(path)
}
