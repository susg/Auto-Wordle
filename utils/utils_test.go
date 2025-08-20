package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindProjectRoot(t *testing.T) {
	// Test that FindProjectRoot returns a directory containing go.mod
	root := FindProjectRoot()

	// Check if the returned path exists
	if _, err := os.Stat(root); os.IsNotExist(err) {
		t.Errorf("FindProjectRoot() returned non-existent path: %s", root)
	}

	// Check if go.mod exists in the returned directory
	goModPath := filepath.Join(root, "go.mod")
	if _, err := os.Stat(goModPath); os.IsNotExist(err) {
		t.Errorf("go.mod not found in returned root directory: %s", root)
	}

	// Verify it returns an absolute path
	if !filepath.IsAbs(root) {
		t.Errorf("FindProjectRoot() should return absolute path, got: %s", root)
	}
}

func TestFindProjectRootConsistency(t *testing.T) {
	// Test that multiple calls return the same result
	root1 := FindProjectRoot()
	root2 := FindProjectRoot()

	if root1 != root2 {
		t.Errorf("FindProjectRoot() returned inconsistent results: %s vs %s", root1, root2)
	}
}
