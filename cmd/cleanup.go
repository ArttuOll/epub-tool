package cmd

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
)

func CleanupE(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("failed to open epub file: %w", err)
	}

	defer r.Close()

	for _, f := range r.File {
		if filepath.Ext(f.Name) == ".css" {
			fmt.Printf("found CSS file: %s\n", f.Name)
			cleanCssFile(f)
		}
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
	buffer := bytes.NewBuffer(make([]byte, 0))

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if isColorDeclaration(trimmed) {
			fmt.Printf("removing color declaration: %s\n", trimmed)
			continue
		}

		if isFontSizeDeclaration(trimmed) {
			fmt.Printf("removing font size declaration: %s\n", trimmed)
			continue
		}

		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("error writing line to buffer: %s", line)
		}
	}

	return nil
}

func isColorDeclaration(line string) bool {
	property := strings.Split(line, ":")[0]
	return property == "color"
}

func isFontSizeDeclaration(line string) bool {
	property := strings.Split(line, ":")[0]
	return property == "font-size"
}
