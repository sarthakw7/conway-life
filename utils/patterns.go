package utils

import "conway-life/game"

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