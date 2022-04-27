package cell

type CellStruct struct {
	X       int
	Y       int
	IsAlive bool
}

func CreateCell(x int, y int) CellStruct {
	cell := CellStruct{
		X:       x,
		Y:       y,
		IsAlive: false,
	}

	return cell
}
