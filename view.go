// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"fmt"
)

// View renders the current game state
func (m GameModel) View() string {
	if m.gameWon {
		return renderVictoryScreen()
	}

	var s string

	// Game over screen
	if m.gameOver {
		s += GameOverStyle.Render("Game Over! Press 'q' to quit.\n")
	} else {
		s += TitleStyle.Render("Minesweeper (Press 'q' to quit)\n")
	}

	s += "\n"

	// Board
	s += renderBoard(m)

	// Instructions
	s += renderInstructions()

	return s
}

// renderBoard renders the game board
func renderBoard(m GameModel) string {
	var s string

	for y := 0; y < BoardHeight; y++ {
		row := ""
		for x := 0; x < BoardWidth; x++ {
			cell := m.board[y][x]
			style := UnrevealedStyle

			if x == m.cursorX && y == m.cursorY {
				style = CursorStyle
			}

			// Show all mines when game is over
			if m.gameOver && cell.IsMine {
				style = RevealedStyle
				row += style.Render("💣")
				continue
			}

			if cell.IsRevealed {
				style = RevealedStyle
			}

			content := "■"
			if !cell.IsRevealed {
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

			row += style.Render(content)
		}
		s += row + "\n"
	}

	return s
}

// renderInstructions renders the game controls instructions
func renderInstructions() string {
	var s string

	s += "\nControls:\n"
	s += "↑↓←→ or hjkl: Move cursor\n"
	s += "Space/Enter: Reveal cell\n"
	s += "f: Toggle flag\n"
	s += "Left Click: Reveal cell\n"
	s += "Right Click: Toggle flag\n"
	s += "q: Quit\n"

	return s
}
