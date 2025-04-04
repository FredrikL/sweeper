// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles user input and game state updates
func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return handleKeypress(m, msg)
	case tea.MouseMsg:
		return handleMouseEvent(m, msg)
	}
	return m, nil
}

// handleKeypress processes keyboard inputs
func handleKeypress(m GameModel, msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "ctrl+c", "q":
		return m, tea.Quit
	case "up", "k":
		if m.cursorY > 0 {
			m.cursorY--
		}
	case "down", "j":
		if m.cursorY < BoardHeight-1 {
			m.cursorY++
		}
	case "left", "h":
		if m.cursorX > 0 {
			m.cursorX--
		}
	case "right", "l":
		if m.cursorX < BoardWidth-1 {
			m.cursorX++
		}
	case " ", "enter":
		if !m.gameOver && !m.gameWon {
			gameOver, gameWon := revealCell(m.board, m.cursorX, m.cursorY)
			m.gameOver = gameOver
			m.gameWon = gameWon
		}
	case "f":
		if !m.gameOver && !m.gameWon {
			toggleFlag(m.board, m.cursorX, m.cursorY)
		}
	}
	return m, nil
}

// handleMouseEvent processes mouse inputs
func handleMouseEvent(m GameModel, msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	// Only handle left and right clicks
	if msg.Type != tea.MouseLeft && msg.Type != tea.MouseRight {
		return m, nil
	}

	// Calculate grid position
	// Account for:
	// - Border padding (1 line + 2 columns)
	// - Title and spacing (2 lines)
	// - Column headers + spacing (3 lines)
	// - Row header (4 columns)
	// - Each cell being 2 characters wide

	adjustedX := msg.X - 6 // Subtract border padding + row headers (1+2+3=6)
	adjustedY := msg.Y - 6 // Subtract border, title, spacing, and headers (1+1+1+3=6)

	if adjustedX < 0 || adjustedY < 0 {
		return m, nil
	}

	gridX := adjustedX / 2 // Each cell is 2 characters wide
	gridY := adjustedY     // Each cell is 1 line high

	// Check if click is within the grid bounds
	if gridX >= 0 && gridX < BoardWidth && gridY >= 0 && gridY < BoardHeight {
		m.cursorX = gridX
		m.cursorY = gridY

		if msg.Type == tea.MouseLeft {
			if !m.gameOver && !m.gameWon {
				gameOver, gameWon := revealCell(m.board, gridX, gridY)
				m.gameOver = gameOver
				m.gameWon = gameWon
			}
		} else if msg.Type == tea.MouseRight {
			if !m.gameOver && !m.gameWon {
				toggleFlag(m.board, gridX, gridY)
			}
		}
	}
	return m, nil
}
