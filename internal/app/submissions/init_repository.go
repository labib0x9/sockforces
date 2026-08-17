package submissions

import (
	"context"
	"errors"
	"fmt"

	"github.com/labib0x9/sockforces/internal/domain/submissions"
)

func (s *service) InitRepository(ctx context.Context, username, labid string) (*submissions.InitResult, error) {
	templateRepo := s.gitRepo.GetTemplateRepository(labid)
	repoName := generateRepositoryName(username, templateRepo)
	repoUrl, err := s.gitRepo.CreateRepository(ctx, templateRepo, repoName)
	if err != nil {
		return nil, s.getError(err)
	}
	err = s.gitRepo.AddCollaborator(ctx, repoName, username)
	if err != nil {
		return nil, s.getError(err)
	}
	return &submissions.InitResult{RepoUrl: repoUrl}, nil
}

func generateRepositoryName(username, templateRepo string) string {
	return fmt.Sprintf("Submission_%s_%s", username, templateRepo)
}

func (s *service) getError(err error) error {
	if msg := s.gitRepo.GetErrorMsg(err); msg != "" {
		return errors.New(msg)
	}
	return err
}
