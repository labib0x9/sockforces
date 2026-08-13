package github

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/labib0x9/sockforces/config"
)

func TestGithubRepo(t *testing.T) {
	cnf := config.GetConfig("../../../.env")

	client := NewClient(cnf.Github)
	repo := NewGithubRepo(client, cnf.Github)
	username := "labib0x9"
	repoName := fmt.Sprintf("tcp-echo-server-%s-intrigation-test", username)

	t.Run("Create Repository", func(t *testing.T) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := repo.CreateRepository(ctx, "tcp-echo-server-template", repoName)
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

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := client.Repositories.Delete(ctx, cnf.Github.Org, repoName)
		if err != nil {
			t.Fatalf("Cleanup Failed: %v\n", err)
		}
	})
}
