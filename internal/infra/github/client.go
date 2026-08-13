package github

import (
	"net/http"
	"time"

	gh "github.com/bradleyfalzon/ghinstallation/v2"
	git "github.com/google/go-github/v62/github"
	"github.com/labib0x9/sockforces/config"
)

type Client struct {
	*git.Client
}

func NewClient(cnf *config.Github) *Client {
	transport, err := gh.NewKeyFromFile(
		http.DefaultTransport,
		cnf.AppId,
		cnf.InstalationId,
		cnf.PrivateKeyPath,
	)
	if err != nil {
		panic(err)
	}

	return &Client{
		Client: git.NewClient(&http.Client{
			Transport: transport,
			Timeout:   10 * time.Second,
		}),
	}
}
