package cmd

import (
	"archive/zip"
	"fmt"
	"path/filepath"
)

func CleanupE(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("failed to open epub file: %w", err)
	}

	defer r.Close()

	cssFiles := make([]*zip.File, 0)
	for _, f := range r.File {
		if filepath.Ext(f.Name) == ".css" {
			cssFiles = append(cssFiles, f)
		}
	}

	for _, f := range cssFiles {
		fmt.Printf("found CSS file: %s\n", f.Name)
	}

	return nil
}
