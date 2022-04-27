package cell_test

import (
	"main/cell"
	"main/game"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCell(t *testing.T) {
	x := 1
	y := 0

	c := cell.CreateCell(x, y)
	assert.True(t, c.IsAlive)
	game.Launch()
	assert.NotNil(t, c)
}
