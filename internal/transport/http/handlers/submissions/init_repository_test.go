package submissions

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labib0x9/sockforces/config"
	subservice "github.com/labib0x9/sockforces/internal/app/submissions"
	"github.com/labib0x9/sockforces/internal/infra/github"
)

func TestInitRepository(t *testing.T) {
	cnf := config.GetConfig("../../../../../.env")

	validate := validator.New()

	gitClient := github.NewClient(cnf.Github)
	gitRepo := github.NewGithubRepo(gitClient, cnf.Github)
	subService := subservice.NewService(gitRepo)
	handler := NewHandler(subService, validate, nil)

	cases := []struct {
		req    initRequest
		code   int
		serial string
	}{
		{serial: "1", req: initRequest{}, code: 422},
		{serial: "2", req: initRequest{Username: "", Labid: "lab-01"}, code: http.StatusUnprocessableEntity},
		{serial: "3", req: initRequest{Username: "9xZer0", Labid: "lab-01"}, code: http.StatusCreated},
		{serial: "4", req: initRequest{Username: "9xZer0", Labid: "lab-01"}, code: http.StatusConflict},
	}

	for _, tc := range cases {
		t.Run(tc.serial, func(t *testing.T) {
			var buf bytes.Buffer
			if err := json.NewEncoder(&buf).Encode(tc.req); err != nil {
				t.Error("Unexpected json encoding errorr", err)
			}

			r := httptest.NewRequest(http.MethodPost, "/short", &buf)
			r.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()

			handler.InitRepository(w, r)

			resp := w.Result()
			if resp.StatusCode != tc.code {
				t.Errorf("Serial: %s, Wanted: %d, Got: %d", tc.serial, tc.code, resp.StatusCode)
			}
		})
	}

	t.Cleanup(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_, err := gitClient.Repositories.Delete(ctx, cnf.Github.Org, generateRepositoryName("9xZer0", gitRepo.GetTemplateRepository("lab-01")))
		if err != nil {
			t.Fatalf("Cleanup Failed: %v\n", err)
		}
	})

}

func generateRepositoryName(username, templateRepo string) string {
	return fmt.Sprintf("Submission-%s-%s", username, templateRepo)
}
