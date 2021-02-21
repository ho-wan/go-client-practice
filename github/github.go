package github

import "context"

// Client ...
type Client struct{}

// NewClient creates a new instance of Client
func NewClient() *Client {
	return &Client{}
}

// GetList ...
func (c *Client) GetList(ctx context.Context) error {
	return nil
}
