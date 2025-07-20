package game

func CountLiveNeighbors(grid Grid, y, x int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			ny := (y + dy + Height) % Height
			nx := (x + dx + Width) % Width

			if dy == 0 && dx == 0 {
				continue
			}
			if ny >= 0 && ny < Height && nx >= 0 && nx < Width {
				if grid[ny][nx] == Alive {
					count++
				}
			}
		}
	}
	return count
}

func NextGeneration(current Grid) Grid {
	newGrid := make(Grid, Height)
	for y:= 0; y < Height; y++ {
		newGrid[y] = make([]Cell, Width)
		for x := 0; x < Width; x++ {
			liveNeighbbors := CountLiveNeighbors(current, y, x)
			if current[y][x] == Alive {
				if liveNeighbbors < 2 || liveNeighbbors > 3 {
					newGrid[y][x] = Dead
				} else {
					newGrid[y][x] = Alive
				}
			} else {
				if liveNeighbbors == 3 {
					newGrid[y][x] = Alive
				} else {
					newGrid[y][x] = Dead
				}
			}
		}
	}
	return newGrid
}