package cell_test

import (
	"main/cell"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCell(t *testing.T) {
	x := 1
	y := 0

	c := cell.CreateCell(x, y)
	assert.False(t, c.IsAlive)
	assert.Equal(t, x, c.X)
	assert.Equal(t, y, c.Y)
}
