package cmd

import (
	"archive/zip"
	"bufio"
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
		cleanCssFile(f)
	}

	return nil
}

func cleanCssFile(f *zip.File) error {
	reader, err := f.Open()
	if err != nil {
		return fmt.Errorf("failed to open CSS file %s: %v", f.Name, err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	//buffer := bytes.NewBuffer(make([]byte, 0))

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("%s\n", line)
	}

	return nil
}
