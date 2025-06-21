package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func FindProjectRoot() string {
	// Start with the current file's directory
	_, filename, _, _ := runtime.Caller(0)
	dir := filepath.Dir(filename)

	// Navigate up until you find go.mod
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}

		// Move up one directory
		parent := filepath.Dir(dir)
		if parent == dir {
			// Reached filesystem root without finding go.mod
			// As a fallback, use the current working directory
			cwd, _ := os.Getwd()
			return cwd
		}
		dir = parent
	}
}
