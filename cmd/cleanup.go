package cmd

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ArttuOll/epub-tool/util"
	"github.com/spf13/cobra"
)

func CleanupE(cmd *cobra.Command, path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return fmt.Errorf("failed to open epub file: %w", err)
	}

	defer r.Close()

	archive, err := os.Create(getOutputFileName(cmd, path))
	if err != nil {
		return fmt.Errorf("failed to create zip archive: %w", err)
	}

	defer archive.Close()

	zipWriter := zip.NewWriter(archive)

	for _, f := range r.File {
		buffer := new(bytes.Buffer)
		if filepath.Ext(f.Name) == ".css" {

			util.LogVerbose(cmd, fmt.Sprintf("found CSS file: %s\n", f.Name))

			buffer, err = cleanCssFile(cmd, f)
			if err != nil {
				return fmt.Errorf("failed to clean CSS file %s: %w", f.Name, err)
			}
		} else {
			readFileToBuffer(f, buffer)
		}

		writeFileToArchive(buffer, zipWriter, f)
	}

	err = zipWriter.Close()
	if err != nil {
		return fmt.Errorf("failed to write zip archive %s", err)
	}

	return nil
}

func cleanCssFile(cmd *cobra.Command, f *zip.File) (*bytes.Buffer, error) {
	reader, err := f.Open()
	if err != nil {
		return nil, fmt.Errorf("failed to open CSS file %s: %v", f.Name, err)
	}

	defer reader.Close()

	scanner := bufio.NewScanner(reader)
	buffer := new(bytes.Buffer)

	for scanner.Scan() {
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		if isColorDeclaration(trimmed) {
			util.LogVerbose(cmd, fmt.Sprintf("removing color declaration: %s\n", trimmed))
			continue
		}

		if isFontSizeDeclaration(trimmed) {
			util.LogVerbose(cmd, fmt.Sprintf("removing font size declaration: %s\n", trimmed))
			continue
		}

		_, err := buffer.WriteString(line + "\n")
		if err != nil {
			return nil, fmt.Errorf("error writing line to buffer: %s", line)
		}
	}

	return buffer, nil
}

func readFileToBuffer(f *zip.File, buffer *bytes.Buffer) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", f.Name, err)
	}
	defer rc.Close()

	_, err = io.Copy(buffer, rc)
	if err != nil {
		return fmt.Errorf("failed to copy file %s into buffer: %w", f.Name, err)
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

func writeFileToArchive(buffer *bytes.Buffer, zipWriter *zip.Writer, f *zip.File) error {
	header := &zip.FileHeader{
		Name:           f.Name,
		Comment:        f.Comment,
		Modified:       f.Modified,
		Method:         f.Method,
		NonUTF8:        f.NonUTF8,
		CreatorVersion: f.CreatorVersion,
		ExternalAttrs:  f.ExternalAttrs,
	}

	fileWriter, err := zipWriter.CreateHeader(header)
	if err != nil {
		return fmt.Errorf("failed to create a writer to add file %s to zip archive: %w", f.Name, err)
	}

	_, err = fileWriter.Write(buffer.Bytes())
	if err != nil {
		return fmt.Errorf("failed to write file %s to the new archive: %w", f.Name, err)
	}

	return nil
}

func getOutputFileName(cmd *cobra.Command, path string) string {
	filename := ""
	outputFileName, _ := cmd.Flags().GetString("outputFileName")
	if outputFileName == "" {
		filename = getDefaultOutputFileName(path)
	} else {
		filename = outputFileName
	}

	return filename
}

func getDefaultOutputFileName(path string) string {
	filename := filepath.Base(path)
	extension := filepath.Ext(path)
	name := filename[:len(filename)-len(extension)]

	return fmt.Sprintf("%s_cleaned.zip", name)
}
