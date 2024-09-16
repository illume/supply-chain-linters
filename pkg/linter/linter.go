package linter

import (
	"fmt"
	"os"
	"path/filepath"
)

// Run executes the linter on the given path.
func Run(path string) error {
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			fmt.Printf("Scanning file: %s\n", path)
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the path: %w", err)
	}

	return nil
}
