// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// GameModel represents the state of the Minesweeper game
type GameModel struct {
	board    [][]Cell
	cursorX  int
	cursorY  int
	gameOver bool
	gameWon  bool
}

// NewModel creates an initialized game model
func NewModel() GameModel {
	return GameModel{
		board:    createBoard(),
		cursorX:  0,
		cursorY:  0,
		gameOver: false,
		gameWon:  false,
	}
}

// Init initializes the Bubble Tea model
func (m GameModel) Init() tea.Cmd {
	return nil
}

// Update handles user input and game state updates
func (m GameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
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
	case tea.MouseMsg:
		// Only handle left and right clicks
		if msg.Type != tea.MouseLeft && msg.Type != tea.MouseRight {
			return m, nil
		}

		// Calculate grid position
		// Each cell is 2 characters wide, and we have a 2-line header
		gridX := msg.X / 2
		gridY := msg.Y - 2

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
	}
	return m, nil
}

// View renders the current game state
func (m GameModel) View() string {
	var s string

	// Special end game victory screen
	if m.gameWon {
		// Create a celebratory victory screen
		victoryStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("46")).
			Bold(true).
			Italic(true).
			Background(lipgloss.Color("236")).
			Padding(1, 3).
			Align(lipgloss.Center).
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("39"))

		// Generate a trophy/celebration art
		trophy := []string{
			"  ___________  ",
			" '._==_==_=_.' ",
			" .-\\:      /-. ",
			"| (|:.     |) |",
			" '-|:.     |-' ",
			"   \\::.    /   ",
			"    '::. .'    ",
			"      ) (      ",
			"    _.' '._    ",
			"   '-------'   ",
		}

		trophyStr := ""
		for _, line := range trophy {
			trophyStr += line + "\n"
		}

		message := lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			Bold(true).
			Render("VICTORY!")

		subMessage := lipgloss.NewStyle().
			Foreground(lipgloss.Color("255")).
			Render("All mines successfully found!")

		controls := lipgloss.NewStyle().
			Foreground(lipgloss.Color("242")).
			Render("Press 'q' to quit")

		return victoryStyle.Render(
			message + "\n\n" +
				lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Render(trophyStr) + "\n" +
				subMessage + "\n\n" +
				controls)
	}

	// Game over screen
	if m.gameOver {
		s += GameOverStyle.Render("Game Over! Press 'q' to quit.\n")
	} else {
		s += TitleStyle.Render("Minesweeper (Press 'q' to quit)\n")
	}

	s += "\n"

	// Board
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

	// Instructions
	s += "\nControls:\n"
	s += "↑↓←→ or hjkl: Move cursor\n"
	s += "Space/Enter: Reveal cell\n"
	s += "f: Toggle flag\n"
	s += "Left Click: Reveal cell\n"
	s += "Right Click: Toggle flag\n"
	s += "q: Quit\n"

	return s
}
