package main

import (
	"conway-life/game"
	"conway-life/render"
	"fmt"
	"time"
)



func main() {
	grid := game.InitializeGrid()

	for i := 0; i <= 50; i++ {
		render.ClearScreen()
		render.PrintGrid(grid)
		fmt.Println("Generation:", i)
		time.Sleep(200 * time.Millisecond)
		grid = game.NextGeneration(grid)
	}
}