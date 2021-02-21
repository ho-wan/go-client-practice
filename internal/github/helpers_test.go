package github_test

import "time"

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Int64 is a helper routine that allocates a new int64 value
// to store v and returns a pointer to it.
func Int64(v int64) *int64 { return &v }

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// Time is a helper routine that returns a single time value
// to allow use inline
func Time(t string) *time.Time {
	tm, _ := time.Parse("2006-01-02T15:04:05Z", t)
	return &tm
}
