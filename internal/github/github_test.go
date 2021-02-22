package github_test

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ho-wan/go-client-practice/internal/github"
	mockhttp "github.com/ho-wan/go-client-practice/internal/github/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	gc := github.NewClient(nil)
	assert.NotNil(t, gc)
}

func TestClient_GetRepos_Success(t *testing.T) {
	tests := []struct {
		name string
		want []github.Repository
	}{
		{
			name: "get repos",
			want: []github.Repository{
				{
					ID:        Int64(1),
					Name:      String("some-repo-name"),
					CreatedAt: Time(t, "2021-01-02T15:04:05Z"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRoundTripper := mockhttp.NewMockRoundTripper(ctrl)

			mockRoundTripper.EXPECT().RoundTrip(gomock.Any()).
				DoAndReturn(func(*http.Request) (*http.Response, error) {
					body, err := json.Marshal(tt.want)
					if err != nil {
						t.Fatalf("failed to marshall json: %v", err)
					}
					return &http.Response{Body: ioutil.NopCloser(bytes.NewReader(body))}, nil
				}).Times(1)

			mockClient := &http.Client{
				Transport: mockRoundTripper,
			}

			gc := github.NewClient(mockClient).WithLogging()
			require.NotNil(t, gc)

			got, err := gc.GetRepos(ctx)
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
