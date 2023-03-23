package github_api

import (
	"context"

	"github.com/google/go-github/v50/github"
	"github.com/pkg/errors"
)

func GetStarredRepos(user string) ([]*github.StarredRepository, error) {
	ctx := context.Background()

	client := github.NewClient(nil)

	repos, _, err := client.Activity.ListStarred(ctx, user, nil)
	if err != nil {
		if _, ok := err.(*github.RateLimitError); ok {
			err = errors.New("failed to get github starred repos: hit github rate limit")
			return nil, err
		} else if _, ok := err.(*github.AcceptedError); ok {
			err = errors.New("failed to get github starred repos: scheduled on GitHub side")
			return nil, err
		} else {
			return nil, errors.Wrap(err, "failed to get github starred repos")
		}
	}

	return repos, nil
}
