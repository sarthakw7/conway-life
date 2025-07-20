package utils

import (
	"conway-life/game"
	"math/rand"
)

func Glider(grid game.Grid, startX, startY int)  {
	grid[startY][startX+1] = game.Alive
	grid[startY+1][startX+2] = game.Alive
	grid[startY+2][startX] = game.Alive
	grid[startY+2][startX+1] = game.Alive
	grid[startY+2][startX+2] = game.Alive
}

func Blinker(grid game.Grid, startX, startY int) {
	grid[startY][startX] = game.Alive
	grid[startY][startX+1] = game.Alive
	grid[startY][startX+2] = game.Alive
}


func Random(grid game.Grid, density float64)  {
	for y := 0; y < game.Height; y++ {
		for x := 0; x < game.Width; x++ {
			if rand.Float64() < density {
				grid[y][x] = game.Alive
			} else {
				grid[y][x] = game.Dead
			}
		}
		
	}
}