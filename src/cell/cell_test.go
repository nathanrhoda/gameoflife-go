package cell_test

import (
	"main/cell"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCell(t *testing.T) {
	x := 0
	y := 0

	c := cell.CreateCell(x, y)

	c = c
	assert.False(t, c.isAlive)

}
