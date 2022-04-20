package cell_test

import (
	"cell"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCell(t *testing.T) {
	x := 0
	y := 0

	c := cell.CreateCell(x, y)
	assert.Equal(t, false, c.isAlive)
}
