package cell

type Cell struct {
	x       int
	y       int
	isAlive bool
}

func CreateCell(x int, y int) Cell {
	return Cell{
		x:       x,
		y:       y,
		isAlive: false,
	}
}
