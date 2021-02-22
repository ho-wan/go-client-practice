package middleware

import (
	"log"
	"net/http"
)

// WithLogging wraps the client with a logging RoundTripper
func WithLogging(c *http.Client) *http.Client {
	c.Transport = &loggingRoundTripper{c.Transport}
	return c
}

type loggingRoundTripper struct {
	transport http.RoundTripper
}

// RoundTrip wraps the client transport and logs the time, request method and url
func (l loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Printf("%s %s", r.Method, r.URL)
	return l.transport.RoundTrip(r)
}
