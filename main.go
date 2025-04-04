// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(
		NewModel(),
		tea.WithMouseCellMotion(),
		tea.WithAltScreen(),
	)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Error running program: %v", err)
	}
}
