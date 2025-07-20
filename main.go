package main

import (
	"conway-life/tui"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)


func main() {
	p := tea.NewProgram(tui.NewModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}