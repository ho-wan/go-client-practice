//go:generate mockgen -destination=./mocks/mockhttp.go -package=mockhttp net/http RoundTripper

package github

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"

	"github.com/pkg/errors"
)

const (
	defaultBaseURL = "https://api.github.com/"
	defaultTimeout = 10 * time.Second
)

// Client holds the github client
type Client struct {
	client  *http.Client
	baseURL *url.URL

	accessToken string
}

// NewClient creates a new instance of Client
func NewClient(client *http.Client) *Client {
	if client == nil {
		client = &http.Client{
			Timeout: defaultTimeout,
		}
	}

	baseURL, err := url.Parse(defaultBaseURL)
	if err != nil {
		log.Println("failed to parse base url", err)
	}

	return &Client{
		client:  client,
		baseURL: baseURL,
	}
}

// Repository ...
type Repository struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// GetRepos gets repos for user
func (c *Client) GetRepos(ctx context.Context) ([]Repository, error) {
	route := "user/repos"

	u := c.baseURL
	u.Path = path.Join(u.Path, route)

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new request")
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch repos")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}
	defer resp.Body.Close()

	var repos []Repository
	if err := json.Unmarshal(body, &repos); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal response")
	}

	return repos, nil
}
