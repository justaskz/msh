package dev

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"

	r "github.com/justaskz/mashina/internal/repos"
	"github.com/ulikunitz/xz"
)

func ExtractTarXZ(repo r.Repo, filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	xzReader, err := xz.NewReader(file)
	if err != nil {
		return "", fmt.Errorf("failed to create xz reader: %w", err)
	}

	tarReader := tar.NewReader(xzReader)

	destDir := filepath.Join(GetOptPath(), string(repo.Uid()))

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return "", fmt.Errorf("error reading tar file: %w", err)
		}

		// Create the full path for extraction
		outPath := filepath.Join(destDir, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			err := os.MkdirAll(outPath, os.FileMode(header.Mode))
			if err != nil {
				return "", fmt.Errorf("failed to create directory: %w", err)
			}

		case tar.TypeReg:
			// Create parent directories if needed
			if err := os.MkdirAll(filepath.Dir(outPath), os.ModePerm); err != nil {
				return "", fmt.Errorf("failed to create parent directories: %w", err)
			}

			// Create file
			outFile, err := os.Create(outPath)
			if err != nil {
				return "", fmt.Errorf("failed to create file: %w", err)
			}
			defer outFile.Close()

			// Copy file content
			_, err = io.Copy(outFile, tarReader)
			if err != nil {
				return "", fmt.Errorf("failed to extract file: %w", err)
			}

			// Set correct permissions
			err = os.Chmod(outPath, os.FileMode(header.Mode))
			if err != nil {
				return "", fmt.Errorf("failed to set file permissions: %w", err)
			}
		}
	}

	fmt.Println("Extraction complete:", destDir)
	return destDir, nil
}
