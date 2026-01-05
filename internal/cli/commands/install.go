package commands

import (
	"fmt"
	"os"

	p "github.com/justaskz/mashina/internal/models/path"
	rs "github.com/justaskz/mashina/internal/models/repos"
)

func Install() {
	// system := s.New()
	repos, err := rs.Read(p.Paths.ConfigPath, p.Paths.LibPath)
	handleError(err)

	for _, repo := range repos {
		if repo.State.IsInstalled() {
			continue
		}

		path := p.Paths.LibPath.String() + "/" + repo.Uid().String()
		os.MkdirAll(path, os.ModePerm)

		// release, err := releases.Get(repo, system)
		// handleError(err)

		// err := releases.Install(release)
		// handleError(err)
	}
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// func Install2() {
// 	systemInfo, err := dev.GetSystemInfo()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	repos, err := dev.GetReposWithState()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	var isVersionFound bool
// 	var isAssetFound bool
// 	for _, repo := range repos {
// 		if repo.State == r.Installed {
// 			continue
// 		}

// 		releases, err := github.GetReleases(repo.Owner, repo.Name)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		isVersionFound = false
// 		for _, release := range releases {
// 			releaseName := dev.ParseVersion(*release.Name)

// 			if releaseName == repo.Version {
// 				isVersionFound = true
// 				isAssetFound = false

// 				for _, asset := range release.Assets {
// 					url := *asset.BrowserDownloadURL
// 					assetSystemInfo, err := dev.GetSystemInfoFromUrl(url)

// 					if err != nil {
// 						continue
// 					}

// 					if strings.Contains(url, "pkg") {
// 						continue
// 					}

// 					if assetSystemInfo == systemInfo {
// 						isAssetFound = true
// 						fmt.Printf("Installing %s from %s \n", repo.Name, url)
// 						repo.Url = r.Url(url)
// 						dev.Install(repo)
// 						break
// 					}
// 				}
// 			}
// 		}

// 		if !isVersionFound {
// 			fmt.Printf("%s version %s was not found \n", repo.Name, repo.Version)
// 			continue
// 		}

// 		if !isAssetFound {
// 			fmt.Printf("No assets was found for %s \n", repo.Name)
// 		}
// 	}

// 	// get url for version
// 	// run install(repo)
// }

// func run(repos r.Repos, systemInfo dev.SystemInfo) {

// }
