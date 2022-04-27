package cell

type CellInterface interface {
	getX() int
	getY() int
	isAlive() bool
}

type CellStruct struct {
	X       int
	Y       int
	IsAlive bool
}

func CreateCell(x int, y int) CellStruct {
	cell := CellStruct{
		X:       x,
		Y:       y,
		IsAlive: true,
	}

	return cell
}
