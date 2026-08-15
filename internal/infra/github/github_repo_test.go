package github

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/labib0x9/sockforces/config"
)

func TestGithubRepo(t *testing.T) {
	cnf := config.GetConfig("../../../.env")

	client := NewClient(cnf.Github)
	repo := NewGithubRepo(client, cnf.Github)
	username := "labib0x9"
	repoName := fmt.Sprintf("tcp-echo-server-%s-intrigation-test", username)
	cloneRepo := "ZERO9xz/Submission-labib0x9-tcp-echo-server-template"
	id := uuid.New().String()
	var path string
	var err error

	t.Run("Create Repository", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := repo.CreateRepository(ctx, "tcp-echo-server-template", repoName)
		if err != nil {
			t.Fatalf("Create Repository Failed: %v\n", err)
		}
	})

	t.Run("Add Collborator", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := repo.AddCollaborator(ctx, repoName, username)
		if err != nil {
			t.Fatalf("Add Collborator Failed: %v\n", err)
		}
	})

	t.Run("Clone Repository", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		path, err = repo.CloneRepository(ctx, cloneRepo, id)
		if err != nil {
			t.Fatalf("Clone Repository Failed: %v\n", err)
		}
	})

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := client.Repositories.Delete(ctx, cnf.Github.Org, repoName)
		if err != nil {
			t.Fatalf("Cleanup Failed: %v\n", err)
		}
	})

	t.Cleanup(func() {
		err := os.RemoveAll(path)
		if err != nil {
			t.Fatalf("Cloned Repo (tmpDir) Cleanup Failed: %v\n", err)
		}
	})
}
