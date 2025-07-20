package main

import (
	"conway-life/game"
	"conway-life/render"
	"conway-life/input"
	"fmt"
	"time"
)



func main() {
	grid := game.InitializeGrid()

	if err := input.InitKeyboard(); err != nil {
		panic(err)
	}
	defer input.CloseKeyboard()

	delay := 200 * time.Millisecond
	running := true
	genration := 0

	for {
		render.ClearScreen()
		render.PrintGrid(grid)
		fmt.Println("Generation:", genration)
		fmt.Println("Controls: [p] pause/resume | [+/-] | [q] quit")

		if running {
			grid = game.NextGeneration(grid)
			genration++
		}

		time.Sleep(delay)

		if char, _ := input.ReadKey(); char != 0 {
			switch char {
			case 'q':
				fmt.Println("Exiting...")
				return
			
			case 'p':
				running = !running
			
			case '+':
				if delay > 50*time.Millisecond {
					delay -= 50 * time.Millisecond
				}

			case '-':
				delay += 50 * time.Millisecond
			}
		}
	}
}