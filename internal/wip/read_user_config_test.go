package wip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadUserConfig(t *testing.T) {
	pack := Pack{Name: "helix", Version: "24.07"}
	expected := Config{Aliases: []string{}, Envs: []string{}, Packs: []Pack{pack}}

	path := "../../fixtures/config/mashina/config.yaml"
	actual := ReadUserConfig(path)

	assert.Equal(t, expected, actual)
}
