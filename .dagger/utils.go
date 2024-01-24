package main

import (
	"fmt"
	"os"
)


// GetHostEnv is like os.Getenv but ensures that the env var is set.
func GetHostEnv(name string) string {
	value, ok := os.LookupEnv(envPrefix + name)
	if !ok {
		fmt.Fprintf(os.Stderr, "env var %s must be set\n", name)
		os.Exit(1)
	}
	return value
}