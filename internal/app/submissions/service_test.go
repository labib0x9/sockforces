package submissions

import (
	"context"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"

	"github.com/labib0x9/sockforces/internal/domain/provider/mocks"
)

func TestService_InitRepository(t *testing.T) {
	ctx := context.Background()
	username := "9xZer0"
	labid := "lab-01"
	templateRepo := "tcp-echo-server-template"
	repoName := generateRepositoryName(username, templateRepo)

	t.Run("success", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		gitRepo := mocks.NewMockGithubRepo(ctrl)

		gitRepo.EXPECT().
			GetTemplateRepository(labid).
			Return(templateRepo)

		gitRepo.EXPECT().
			CreateRepository(ctx, templateRepo, repoName).
			Return("https://github.com/org/"+repoName, nil)

		gitRepo.EXPECT().
			AddCollaborator(ctx, repoName, username).
			Return(nil)

		svc := NewService(gitRepo)
		result, err := svc.InitRepository(ctx, username, labid)

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		if result.RepoUrl != "https://github.com/org/"+repoName {
			t.Errorf("unexpected RepoUrl: %s", result.RepoUrl)
		}
	})

	t.Run("create repository fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		gitRepo := mocks.NewMockGithubRepo(ctrl)
		wantErr := errors.New("template not found")

		gitRepo.EXPECT().
			GetTemplateRepository(labid).
			Return(templateRepo)

		gitRepo.EXPECT().
			CreateRepository(ctx, templateRepo, repoName).
			Return("", wantErr)

		// // AddCollaborator must NOT be called if CreateRepository fails.
		// gitRepo.EXPECT().
		// 	AddCollaborator(gomock.Any(), gomock.Any(), gomock.Any()).
		// 	Times(0)

		svc := NewService(gitRepo)
		result, err := svc.InitRepository(ctx, username, labid)

		if !errors.Is(err, wantErr) {
			t.Fatalf("expected %v, got %v", wantErr, err)
		}
		if result != nil {
			t.Errorf("expected nil result on error, got %+v", result)
		}
	})

	t.Run("add collaborator fails", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		gitRepo := mocks.NewMockGithubRepo(ctrl)
		wantErr := errors.New("user not found")

		gitRepo.EXPECT().
			GetTemplateRepository(labid).
			Return(templateRepo)

		gitRepo.EXPECT().
			CreateRepository(ctx, templateRepo, repoName).
			Return("https://github.com/org/"+repoName, nil)

		gitRepo.EXPECT().
			AddCollaborator(ctx, repoName, username).
			Return(wantErr)

		svc := NewService(gitRepo)
		result, err := svc.InitRepository(ctx, username, labid)

		if !errors.Is(err, wantErr) {
			t.Fatalf("expected %v, got %v", wantErr, err)
		}
		if result != nil {
			t.Errorf("expected nil result on error, got %+v", result)
		}
	})
}
