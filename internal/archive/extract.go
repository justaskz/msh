package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func Extract(path, destination string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	rootDir, err := findRootDir(tarReader)
	if err != nil {
		return err
	}

	file.Seek(0, io.SeekStart)
	gzipReader, _ = gzip.NewReader(file)
	tarReader = tar.NewReader(gzipReader)

	err = extract(tarReader, rootDir, destination)
	if err != nil {
		return err
	}

	return nil
}

func findRootDir(tarReader *tar.Reader) (string, error) {
	var rootDir string
	var firstEntry bool = true

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return "", err
		}

		if strings.HasPrefix(filepath.Base(header.Name), "._") {
			continue
		}

		cleanName := filepath.Clean(header.Name)
		parts := strings.Split(cleanName, string(os.PathSeparator))

		if firstEntry {
			rootDir = parts[0]
			firstEntry = false
		} else if len(parts) > 1 && parts[0] != rootDir {
			rootDir = ""
		}
	}

	return rootDir, nil
}

func extract(tarReader *tar.Reader, rootDir, destination string) error {
	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if strings.HasPrefix(filepath.Base(header.Name), "._") {
			continue
		}

		extractPath := filepath.Join(destination, header.Name)
		if rootDir != "" {
			extractPath = filepath.Join(destination, strings.TrimPrefix(header.Name, rootDir+"/"))
		}

		if header.Typeflag == tar.TypeDir {
			if err := os.MkdirAll(extractPath, os.ModePerm); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(extractPath), os.ModePerm); err != nil {
				return err
			}
			outFile, err := os.Create(extractPath)
			if err != nil {
				return err
			}
			defer outFile.Close()

			if _, err := io.Copy(outFile, tarReader); err != nil {
				return err
			}
		}
	}

	return nil
}
