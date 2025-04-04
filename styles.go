// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Define styles for the game UI
var (
	// Cell styles
	UnrevealedStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(lipgloss.Color("236")).
			Foreground(lipgloss.Color("255"))

	RevealedStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(lipgloss.Color("235")).
			Foreground(lipgloss.Color("255"))

	CursorStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(lipgloss.Color("62")).
			Foreground(lipgloss.Color("255"))

	// Number colors
	NumberStyles = []lipgloss.Style{
		lipgloss.NewStyle().Foreground(lipgloss.Color("39")),  // 1
		lipgloss.NewStyle().Foreground(lipgloss.Color("46")),  // 2
		lipgloss.NewStyle().Foreground(lipgloss.Color("196")), // 3
		lipgloss.NewStyle().Foreground(lipgloss.Color("93")),  // 4
		lipgloss.NewStyle().Foreground(lipgloss.Color("208")), // 5
		lipgloss.NewStyle().Foreground(lipgloss.Color("51")),  // 6
		lipgloss.NewStyle().Foreground(lipgloss.Color("255")), // 7
		lipgloss.NewStyle().Foreground(lipgloss.Color("240")), // 8
	}

	// Game status styles
	GameOverStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("196")).
			Bold(true)

	GameWonStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("46")).
			Bold(true)

	TitleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("62")).
			Bold(true)
)
