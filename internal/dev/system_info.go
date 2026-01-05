package dev

import (
	"fmt"
	"runtime"
	"strings"
)

type Arch string
type Platform string

const Arm64 Arch = "arm64"
const Amd64 Arch = "amd64"

const Darwin Platform = "darwin"
const Linux Platform = "linux"

// type platforms struct {
// 	Darwin Platform
// }

// var Platforms = platforms{
// 	Darwin: Platform("darwin"),
// }

type SystemInfo struct {
	Arch     Arch
	Platform Platform
}

func GetSystemInfo() (SystemInfo, error) {
	url := runtime.GOARCH + "-" + runtime.GOOS
	systemInfo, err := GetSystemInfoFromUrl(url)
	if err != nil {
		return SystemInfo{}, err
	}

	return systemInfo, nil
}

func GetSystemInfoFromUrl(url string) (SystemInfo, error) {
	url = strings.ToLower(url)
	arch, err := parseArch(url)
	if err != nil {
		return SystemInfo{}, err
	}

	platform, err := parsePlatform(url)
	if err != nil {
		return SystemInfo{}, err
	}

	systemInfo := SystemInfo{
		Arch:     arch,
		Platform: platform,
	}

	return systemInfo, nil
}

func parseArch(url string) (Arch, error) {
	switch {
	case strings.Contains(url, "arm64") || strings.Contains(url, "aarch64"):
		return Arm64, nil
	case strings.Contains(url, "amd64") || strings.Contains(url, "x86_64"):
		return Amd64, nil
	default:
		return Arch(""), fmt.Errorf("failed to parse architecture from url %s", url)
	}
}

func parsePlatform(url string) (Platform, error) {
	switch {
	case strings.Contains(url, "darwin") || strings.Contains(url, "macos"):
		return Darwin, nil
	case strings.Contains(url, "linux"):
		return Linux, nil
	default:
		return Platform(""), fmt.Errorf("failed to parse platform from url %s", url)
	}
}
