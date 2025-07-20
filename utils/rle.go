package utils

import (
	"bufio"
	"conway-life/game"
	"os"
	"strconv"
	"strings"
)

func LoadRLE(path string) (game.Grid, error)  {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	grid := game.InitializeGrid()
	scanner := bufio.NewScanner(file)
	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") || strings.HasPrefix(line, "x =") {
			continue
		}

		count := ""

		for _, r := range line {
			switch r {
			case 'b':
				n := parseCount(count)
				for i := 0; i < n; i++ {
					if y<game.Height && x <game.Width {
						grid[y][x] = game.Dead
					}
					x++
				}
				count = ""
			
			case 'o':
				n := parseCount(count)
				for i := 0; i < n; i++ {
					if y < game.Height && x < game.Width {
						grid[y][x] = game.Alive
					}
					x++
				}
				count = ""

			case '$':
				n := parseCount(count)
				if n==0 {
					n = 1
				}
				y += n
				x = 0
				count = ""

			case '!':
				break

			default:
				if r> '0' && r <= '9' {
					count += string(r)
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return grid, nil
}

func parseCount(s string) int {
	if s == ""{
		return 1
	}
	n, _ := strconv.Atoi(s)
	return n;
}