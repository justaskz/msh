package config

import "embed"

//go:embed init.sh
var EmbedFiles embed.FS
