package game

import "math/rand"

type Cell bool

const (
	Dead Cell = false
	Alive Cell = true
)

type Grid [][]Cell

const (
	Width = 50
	Height = 20
)

func InitializeGrid() Grid {
	grid := make(Grid, Height)
	for y := 0; y < Height; y++ {
		grid[y] = make([]Cell, Width)
		for x := 0; x < Width; x++ {
			if rand.Float64() < 0.25 {
				grid[y][x] = Alive
			} else {
				grid[y][x] = Dead
			}
		}
	}
	return grid
}