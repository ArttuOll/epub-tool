package cmd

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CleanupE(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("failed to open epub file: %w", err)
	}

	defer r.Close()

	// TODO: the file name has to be configurable in the command
	archive, err := os.Create("archive.zip")
	if err != nil {
		return fmt.Errorf("failed to create zip archive: %w", err)
	}

	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	for _, f := range r.File {
		buffer := bytes.NewBuffer(make([]byte, 0))
		if filepath.Ext(f.Name) == ".css" {
			fmt.Printf("found CSS file: %s\n", f.Name)
			buffer, err = cleanCssFile(f)
			if err != nil {
				return fmt.Errorf("failed to clean CSS file %s: %w", f.Name, err)
			}
		}

		fileWriter, err := zipWriter.Create(f.Name)
		if err != nil {
			return fmt.Errorf("failed to create a writer to add file %s to zip archive: %w", f.Name, err)
		}

		_, err = fileWriter.Write(buffer.Bytes())
		if err != nil {
			return fmt.Errorf("failed to write file %s to the new archive: %w", f.Name, err)
		}
	}

	err = zipWriter.Close()
	if err != nil {
		return fmt.Errorf("failed to write zip archive %s", err)
	}

	return nil
}

func cleanCssFile(f *zip.File) (*bytes.Buffer, error) {
	reader, err := f.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open CSS file %s: %v", f.Name, err)
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
			return nil, fmt.Errorf("error writing line to buffer: %s", line)
		}
	}

	return buffer, nil
}

func isColorDeclaration(line string) bool {
	property := strings.Split(line, ":")[0]
	return property == "color"
}

func isFontSizeDeclaration(line string) bool {
	property := strings.Split(line, ":")[0]
	return property == "font-size"
}
