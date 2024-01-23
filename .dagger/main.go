package main

import (
	"fmt"
	"strconv"
)

const (
	defaultGoImageVersion = "1.21.6-alpine3.19"
	workdir               = "/app"
	envPrefix             = "AGATE_INDEXER_"
)

type Agate struct {
	// Image version of the Go image to use (default to: `1.21.6-alpine3.19`).
	GoImageVersion string
}

func New(goImageVersion Optional[string]) *Agate {
	return &Agate{
		GoImageVersion: goImageVersion.GetOr(defaultGoImageVersion),
	}
}

// CI gives access to CI commands.
func (m *Agate) CI() *CI {
	return &CI{
		Project: m.Project(),
	}
}

// Database returns a PostgreSQL database.
// Example usage: `dagger up --native database`
func (m *Agate) Database(
	username Optional[string],
	password Optional[string],
	dbName Optional[string],
	port Optional[int],
) *Service {
	var defaultPort int

	if !port.isSet {
		defaultPortStr := GetHostEnv("DATABASE_PORT")

		defaultPortTmp, err := strconv.Atoi(defaultPortStr)
		if err != nil {
			fmt.Println(fmt.Errorf("invalid port: %w", err))
			return nil
		}

		defaultPort = defaultPortTmp
	}

	return dag.
		Container().
		From("bitnami/postgresql:15.2.0").
		WithEnvVariable("POSTGRES_USER", username.GetOr(GetHostEnv("DATABASE_USERNAME"))).
		WithSecretVariable("POSTGRES_PASSWORD", dag.SetSecret("dbPassword", password.GetOr(GetHostEnv("DATABASE_PASSWORD")))).
		WithEnvVariable("POSTGRES_DATABASE", dbName.GetOr(GetHostEnv("DBNAME"))).
		WithMountedCache("/var/lib/postgresql/data", dag.CacheVolume("pg-data")).
		WithExposedPort(port.GetOr(defaultPort)).
		AsService()
}

// Binary returns Agate binary file to the host's current directory.
// Example usage: `dagger export binary`
func (m *Agate) Binary() *File {
	return m.CI().Build().File("agate")
}

// Project returns the container that will be used in further Dagger command.
// Open a shell inside it using `dagger shell project --entrypoint /bin/sh`.
func (m *Agate) Project() *Container {
	return dag.
		Container().
		From("golang:"+m.GoImageVersion).
		WithWorkdir(workdir).
		// Install build-base to build cgo packages.
		WithExec([]string{"apk", "add", "build-base"}).

		// Add cache for go modules and go build.
		WithMountedCache("/go/pkg/mod", dag.CacheVolume("gomod")).
		WithMountedCache("/root/.cache/go-build", dag.CacheVolume("gobuild")).

		// Install go dependencies
		WithFile("go.mod", source().File("go.mod")).
		WithFile("go.sum", source().File("go.sum")).
		WithExec([]string{"go", "mod", "download"}).

		// Add source code
		WithMountedDirectory(workdir, source())
}
