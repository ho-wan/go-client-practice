package main

import (
	"context"
	"log"
	"net/http"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/ho-wan/go-client-practice/internal/config"
	"github.com/ho-wan/go-client-practice/internal/github"
	"github.com/ho-wan/go-client-practice/internal/middleware"
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

	// use oauth2 to create a Transport with a static token
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: cfg.Github.AccessToken},
	)
	authTransport := &oauth2.Transport{
		Source: ts,
		Base:   http.DefaultTransport,
	}

	// initialise a new retryable client
	rc := retryablehttp.NewClient()
	rc.RetryMax = 10

	// add token to the transport, and wrap with a logging transport
	sc := rc.StandardClient()
	sc.Transport = authTransport
	sc = middleware.WithLogging(sc)

	gc := github.NewClient(sc)
	return gc
}
