// Package version holds build metadata injected via -ldflags at compile time.
package version

import "fmt"

var (
	// Version is the semantic release tag or "dev" for local builds.
	Version = "dev"
	// Commit is the VCS revision.
	Commit = "none"
	// BuildTime is an RFC3339 or opaque build timestamp.
	BuildTime = "unknown"
)

// String returns a short version suitable for APIs and logs.
func String() string {
	return Version
}

// Full returns version, commit, and build time for startup diagnostics.
func Full() string {
	return fmt.Sprintf("%s (commit=%s built=%s)", Version, Commit, BuildTime)
}
