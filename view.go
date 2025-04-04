// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View renders the current game state with synthwave styling
func (m GameModel) View() string {
	if m.gameWon {
		return renderVictoryScreen()
	}

	var s string

	// Game title with synthwave styling
	titleText := "SYNTHWAVE MINESWEEPER"
	if m.gameOver {
		titleText = "GAME OVER! Press 'q' to quit"
	}

	title := lipgloss.NewStyle().
		Foreground(neonPink).
		Bold(true).
		Padding(0, 0, 1, 0).
		Render(titleText)

	s += title + "\n"

	// Board
	s += renderBoard(m)

	// Instructions
	s += renderInstructions()

	// Apply a simple border to the whole UI
	return lipgloss.NewStyle().
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(neonPink).
		Padding(1, 2).
		Render(s)
}

// renderBoard renders the game board with synthwave styling
func renderBoard(m GameModel) string {
	var s string

	// Create header style for the coordinates
	headerStyle := lipgloss.NewStyle().
		Foreground(neonBlue).
		Bold(true)

	// Column headers
	s += "   "
	for x := 0; x < BoardWidth; x++ {
		s += headerStyle.Render(fmt.Sprintf(" %d ", x%10))
	}
	s += "\n\n"

	// Board rows
	for y := 0; y < BoardHeight; y++ {
		// Row number
		s += headerStyle.Render(fmt.Sprintf(" %d ", y%10)) + " "

		for x := 0; x < BoardWidth; x++ {
			cell := m.board[y][x]
			style := UnrevealedStyle

			if x == m.cursorX && y == m.cursorY {
				style = CursorStyle
			} else if cell.IsRevealed {
				style = RevealedStyle
			}

			content := "■"
			if m.gameOver && cell.IsMine {
				content = "💣"
			} else if !cell.IsRevealed {
				if cell.IsFlagged {
					content = "⚑"
				}
			} else if cell.IsMine {
				content = "💣"
			} else if cell.Neighbors > 0 {
				content = fmt.Sprintf("%d", cell.Neighbors)
				content = NumberStyles[cell.Neighbors-1].Render(content)
			} else {
				content = " "
			}

			s += style.Render(content)
		}
		s += "\n"
	}

	return s
}

// renderInstructions renders the game controls instructions with synthwave styling
func renderInstructions() string {
	headerStyle := lipgloss.NewStyle().
		Foreground(neonPink).
		Bold(true)

	controlStyle := lipgloss.NewStyle().
		Foreground(neonBlue)

	keyStyle := lipgloss.NewStyle().
		Foreground(neonYellow).
		Bold(true)

	s := "\n" + headerStyle.Render("CONTROLS") + "\n\n"

	s += controlStyle.Render("Move: ") + keyStyle.Render("↑↓←→") + controlStyle.Render(" or ") + keyStyle.Render("hjkl") + "\n"
	s += controlStyle.Render("Reveal: ") + keyStyle.Render("Space") + " " + controlStyle.Render("or") + " " + keyStyle.Render("Enter") + "\n"
	s += controlStyle.Render("Flag: ") + keyStyle.Render("f") + "\n"
	s += controlStyle.Render("Mouse: ") + keyStyle.Render("Left click") + " = Reveal, " + keyStyle.Render("Right click") + " = Flag\n"
	s += controlStyle.Render("Quit: ") + keyStyle.Render("q") + "\n"

	return s
}
