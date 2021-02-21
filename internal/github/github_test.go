package github_test

import (
	"testing"

	"github.com/ho-wan/go-client-practice/internal/github"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	gc := github.NewClient(nil)
	assert.NotNil(t, gc)
}
