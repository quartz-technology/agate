package main

import (
	"context"

	"golang.org/x/sync/errgroup"
)

type CI struct {
	Project *Container
}

// Run runs all CI steps in parallel.
// It runs: unit tests, lint and build.
// Example usage: `dagger call ci run`
func (c *CI) Run(ctx context.Context) error {
	wg, gctx := errgroup.WithContext(ctx)

	wg.Go(func() error {
		_, err := c.UnitTest(gctx)

		return err
	})

	wg.Go(func() error {
		_, err := c.Lint(gctx)

		return err
	})

	wg.Go(func() error {
		_, err := c.Build().Stdout(gctx)

		return err
	})

	return wg.Wait()
}

// UnitTest runs go unit tests on the project.
// Example usage: `dagger call ci unit-test`
func (c *CI) UnitTest(ctx context.Context) (string, error) {
	return c.
		Project.
		Pipeline("unit-test").
		WithExec([]string{"go", "test", "-v", "-race", "./...",}).
		Stdout(ctx)
}

// Lint runs golangci-lint on the project.
// Example usage: `dagger call ci lint`
func (c *CI) Lint(ctx context.Context) (string, error) {
	// We cannot use golangci-lint official image because it doesn't support go
	// 1.21.6 for now
	return c.
		Project.
		Pipeline("lint").
		
		// Add binutils-gold for golang
		// See: https://github.com/golang/go/issues/52399
		WithExec([]string{"apk", "add", "binutils-gold"}).

		// Manually install golangci-lint binary
		WithExec([]string{"go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2"}).
		WithExec([]string{"golangci-lint", "run", "-v", "--timeout", "5m"}).
		Stdout(ctx)
}

// Build returns a container with Agate binary built in it.
// Example usage: `dagger call ci build`
func (c *CI) Build() *Container {
	return c.
		Project.
		Pipeline("build").
		WithExec([]string{"go", "build", "-o", "agate", "main.go"})
}
