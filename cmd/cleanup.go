package cmd

import (
	"archive/zip"
	"fmt"
)

func CleanupE(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("failed to open epub file: %w", err)
	}

	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for _, f := range r.File {
		fmt.Printf("Files contained by the given epub %s:\n", f.Name)
	}

	return nil
}
