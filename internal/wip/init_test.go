package wip

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	expected := "export XDG_CONFIG_HOME"
	actual := Init()

	assert.Contains(t, actual, expected)
}
