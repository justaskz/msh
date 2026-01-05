package dev

import (
	"fmt"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/justaskz/mashina/internal/archive"
	r "github.com/justaskz/mashina/internal/repos"
)

func Install(repo r.Repo) error {
	fmt.Println("Dowloading:", repo.Url)
	path, err := downloadFile(repo)
	if err != nil {
		return err
	}

	fmt.Println("Extracting", path)
	destination := filepath.Join(GetLibPath(), repo.Uid().String())
	err = archive.Extract(path, destination)
	if err != nil {
		return err
	}

	binPath, err := findBin(destination, repo.Name)
	if err != nil {
		return err
	}

	err = os.Chmod(binPath, fs.FileMode(0700))
	if err != nil {
		return err
	}

	err = UpdateSymlink(repo)
	if err != nil {
		return err
	}

	return nil
}

func downloadFile(repo r.Repo) (string, error) {
	saveDir := "/tmp/msh"
	err := os.MkdirAll(saveDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directory: %w", err)
	}

	fileName := repo.Url.BasePath()
	savePath := filepath.Join(saveDir, fileName)

	if fileExists(savePath) {
		return savePath, nil
	}

	resp, err := http.Get(repo.Url.String())
	if err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("bad status: %s", resp.Status)
	}

	outFile, err := os.Create(savePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save file: %w", err)
	}

	fmt.Println("File downloaded:", savePath)
	return savePath, nil
}

func fileExists(savePath string) bool {
	_, err := os.Stat(savePath)
	return !os.IsNotExist(err)
}
