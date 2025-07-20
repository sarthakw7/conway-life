package render

import (
	"fmt"
	"conway-life/game"
)

func PrintGrid(grid game.Grid) {
	for y := 0; y < game.Height; y++ {
		for x := 0; x < game.Width; x++ {
			if grid[y][x] == game.Alive {
				fmt.Print("X")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}