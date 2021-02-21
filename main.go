package main

import (
	"context"
	"log"

	"github.com/ho-wan/go-client-practice/internal/config"
	"github.com/ho-wan/go-client-practice/internal/github"
	"golang.org/x/oauth2"
)

const (
	configFilename = "config.yaml"
)

func main() {
	ctx := context.Background()
	gc := getAuthenticatedGithubClient(ctx)

	repos, err := gc.GetRepos(ctx)
	if err != nil {
		log.Fatalln("failed to get repos", err)
	}

	log.Println("got repos:", len(repos))
}

func getAuthenticatedGithubClient(ctx context.Context) *github.Client {
	cfg, err := config.LoadConfig(configFilename)
	if err != nil {
		log.Fatalln("failed to load config", err)
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Github.AccessToken},
	)

	tc := oauth2.NewClient(ctx, ts)

	gc := github.NewClient(tc)
	return gc
}
