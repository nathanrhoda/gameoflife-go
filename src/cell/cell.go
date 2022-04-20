package cell

type CellClass struct {
	x       int
	y       int
	isAlive bool
}

func CreateCell(x int, y int) CellClass {
	return CellClass{
		x:       x,
		y:       y,
		isAlive: false,
	}
}
