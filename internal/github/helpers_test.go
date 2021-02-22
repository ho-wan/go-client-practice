package github_test

import (
	"testing"
	"time"
)

// Bool is a helper function that returns a pointer to a bool
func Bool(v bool) *bool { return &v }

// Int is a helper function that returns a pointer to an int
func Int(v int) *int { return &v }

// Int64 is a helper function that returns a pointer to an int64
func Int64(v int64) *int64 { return &v }

// String is a helper function that returns a pointer to a string
func String(v string) *string { return &v }

// Time is a helper function that parses and returns a time with an error
func Time(t *testing.T, v string) *time.Time {
	t.Helper()
	tm, err := time.Parse("2006-01-02T15:04:05Z", v)
	if err != nil {
		t.Fatalf("failed to parse time: %v", err)
	}
	return &tm
}
