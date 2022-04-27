package game_test

import (
	"main/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLaunch(t *testing.T) {
	game.Launch()
	assert.True(t, false)
}
