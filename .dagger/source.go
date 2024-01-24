package main

import (
	"os"
	"path/filepath"
)

func source() *Directory {
	return dag.Host().Directory(root(), HostDirectoryOpts{
		Exclude: []string{
			".dagger",
		},
	})
}

// Return the absolute path to the root of the project.
// This is required for now by Dagger but it should be fixed later.
// TODO: fix .. restriction when it's merge on dagger
func root() string {
	workdir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return filepath.Join(workdir, "..")
}