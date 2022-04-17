package cell_test

import (
	"cell"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCell(t *testing.T) {

	c := cell.Cell{
		x:       1,
		y:       2,
		isAlive: false,
	}

	alive := c.CreateCell()
	assert.Equal(t, true, alive)
}
