package github

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/google/go-github/v52/github"
	"github.com/gregjones/httpcache"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

type PullRequest struct {
	Owner      string
	Repository string
	Number     int

	client      *github.Client
	pullRequest *github.PullRequest
}

func NewPullRequest(logger *logrus.Logger, owner, repository string, number int) (*PullRequest, error) {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		return nil, fmt.Errorf("the environment variable GITHUB_TOKEN is missing")
	}

	// Setup in-memory HTTP caching engine
	cache := &http.Client{
		Transport: &oauth2.Transport{
			Base:   httpcache.NewMemoryCacheTransport(),
			Source: oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token}),
		},
	}

	client := github.NewClient(cache)

	pull, _, err := client.PullRequests.Get(context.Background(), owner, repository, number)
	if err != nil {
		return nil, fmt.Errorf("unable to fetch %s/%s#%d: %w", owner, repository, number, err)
	}

	pullRequest := &PullRequest{
		Owner:       owner,
		Repository:  repository,
		Number:      number,
		client:      client,
		pullRequest: pull,
	}

	return pullRequest, nil
}

func (p *PullRequest) GetTitle() string {
	if p == nil || p.pullRequest == nil {
		return ""
	}

	return *p.pullRequest.Title
}

func (p *PullRequest) GetLabels() []*github.Label {
	if p == nil || p.pullRequest == nil {
		return nil
	}

	return p.pullRequest.Labels
}
