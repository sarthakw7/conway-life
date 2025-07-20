package tui

import (
	"conway-life/game"
	"conway-life/utils"
	"fmt"

	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	grid       game.Grid
	running    bool
	speed      time.Duration
	generation int
	stepOnce   bool
	editMode   bool
	cursorX    int
	cursorY    int
}

var (
	aliveStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	deadStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("236"))

	gridBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			Padding(1, 2).
			BorderForeground(lipgloss.Color("62"))

	infoStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			MarginTop(1)
	
	headingStyle = lipgloss.NewStyle().
			Bold(true). 
			Foreground(lipgloss.Color("205"))
)

func NewModel() model {
	grid := resetGrid("glider")
	
	return model{
		grid: grid,
		running: true,
		speed: 200 * time.Millisecond,
		generation: 0,
		editMode: false,
		cursorX: 0,
		cursorY: 0,
	}
}

func (m model) Init() tea.Cmd {
	return tea.Tick(m.speed, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

type tickMsg struct{}


func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd)  {
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "p":
			m.running = !m.running

		case "s":
			if !m.running {
				m.stepOnce = true
			}

		case "e":
			m.editMode = true
			m.running = false

		case "enter":
			if m.editMode {
				m.editMode = false
				m.running = true
			}

		case "w":
			_ = utils.SaveGrid(m.grid)
		
		case "l":
			loaded, err := utils.LoadGrid()
			if err == nil {
				m.grid = loaded
				m.generation = 0
			}

		case "r":
			loaded, err := utils.LoadRLE("patterns/gosper-glider-gun.rle")
			if err == nil {
				m.grid = loaded
				m.generation = 0
			}

		case "+":
			if m.speed > 50 * time.Millisecond {
				m.speed -= 50 * time.Millisecond
			}

		case "-":
			m.speed += 50 * time.Millisecond
			
		case "1":
			m.grid = resetGrid("glider")
			m.generation = 0

		case "2":
			m.grid = resetGrid("blinker")
			m.generation = 0

		case "3":
			m.grid = resetGrid("random")
			m.generation = 0
		}

		if m.editMode {
			switch msg.String() {
			case "up":
				if m.cursorY > 0 {
					m.cursorY--
				}
			
			case "down":
				if m.cursorY < game.Height-1 {
					m.cursorY++
				}

			case "left":
				if m.cursorX > 0 {
				m.cursorX--
			}

			case "right":
				if m.cursorX < game.Width-1 {
				m.cursorX++
			}

			case " ":
				if m.grid[m.cursorY][m.cursorX] == game.Alive {
					m.grid[m.cursorY][m.cursorX] = game.Dead
				} else {
					m.grid[m.cursorY][m.cursorX] = game.Alive
				}
			}

		}

	case tickMsg:
		if m.running || m.stepOnce {
			m.grid = game.NextGeneration(m.grid)
			m.generation++
			m.stepOnce = false
		}
	}

	return m, tea.Tick(m.speed, func(t time.Time) tea.Msg {
		return tickMsg{}
	})
}

func (m model) View() string {
	var gridView string
	for y := 0; y<game.Height; y++ {
		for x:= 0; x<game.Width; x++ {
			cellChar := " "
			if m.grid[y][x] == game.Alive {
				cellChar = "█"
			}

			if m.editMode && x == m.cursorX && y == m.cursorY {
				cellChar = lipgloss.NewStyle().
					Background(lipgloss.Color("5")). // magenta bg
					Foreground(lipgloss.Color("0")). // black fg
					Render(cellChar)
			} else if m.grid[y][x] == game.Alive {
				cellChar = aliveStyle.Render(cellChar)
			} else {
				cellChar = deadStyle.Render(cellChar)
			}

			gridView += cellChar
		}
		gridView += "\n"
	}
	styledGrid := gridBox.Render(gridView)

	mode := "LIVE MODE"
	if m.editMode {
		mode = "EDIT MODE — arrows to move, [space] toggle, [enter] resume"
	}
	modeLine := infoStyle.Render(mode)

	status := headingStyle.Render(fmt.Sprintf("Generation: %d  |  Speed: %v", m.generation, m.speed))
	controls := infoStyle.Render("[1] Glider   [2] Blinker   [3] Random [r] Load RLE\n[p] Pause   [+/-] Speed   [s] Step   [e] Edit   [w] Save   [l] Load   [q] Quit")




	return lipgloss.JoinVertical(lipgloss.Left, styledGrid, status,modeLine, controls)

}


func resetGrid(pattern string) game.Grid  {
	grid := game.InitializeGrid()
	switch pattern {
	case "glider":
		utils.Glider(grid, 1, 1)
	case "blinker":
		utils.Blinker(grid, 5, 5)
	case "random":
		utils.Random(grid, 0.25)
	}
	return grid
}
