package myTeams

import (
	"errors"
	// "github.com/google/go-github/v39/github"
	"os"
)

// need to do an oauth sign in
// use oauth to get team membership https://docs.github.com/en/rest/reference/teams#list-teams-for-the-authenticated-user
type OAuth struct {
	secret string
	id     string
}

func MakeOAuth() (*OAuth, error) {

	secret, ok := os.LookupEnv("GITHUB_CLIENT_SECRET")
	if !ok || secret == "" {
		return nil, errors.New("GITHUB_CLIENT_SECRET environment variable not available.")
	}

	id, ok := os.LookupEnv("GITHUB_CLIENT_ID")
	if !ok || id == "" {
		return nil, errors.New("GITHUB_CLIENT_ID environment variable not available.")
	}

	return &OAuth{secret: secret, id: id}, nil
}
