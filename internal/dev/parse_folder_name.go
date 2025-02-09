package dev

import (
	"fmt"
	"strings"
)

func ParseFolderName(folderName string) (*string, *string, *string, error) {
	parts := strings.Split(folderName, "__")

	if len(parts) != 3 {
		return nil, nil, nil, fmt.Errorf("folder '%s' format is incorrect", folderName)
	}

	return &parts[0], &parts[1], &parts[2], nil
}
