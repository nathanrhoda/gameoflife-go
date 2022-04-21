package main

import (
	"fmt"
	cell "main/cell"
)

type BOOM struct {
	x       int
	y       int
	isAlive bool
}

func main() {
	fmt.Println("Game Of Life")
	c := cell.CreateCell(1, 1)
	fmt.Println(c)

	b := BOOM{
		x:       1,
		y:       2,
		isAlive: true,
	}
	fmt.Println(b.isAlive)
}
