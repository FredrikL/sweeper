// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"github.com/charmbracelet/lipgloss"
)

// Synthwave color palette
var (
	// Background colors
	bgDark   = lipgloss.Color("#121036") // Dark purple/blue background
	bgMedium = lipgloss.Color("#1e1a44") // Medium purple background
	bgLight  = lipgloss.Color("#2d2a57") // Light purple background

	// Neon colors
	neonPink   = lipgloss.Color("#ff00ff") // Bright pink
	neonBlue   = lipgloss.Color("#00ffff") // Cyan/blue
	neonPurple = lipgloss.Color("#bf00ff") // Purple
	neonOrange = lipgloss.Color("#ff9900") // Orange
	neonYellow = lipgloss.Color("#ffff00") // Yellow
	neonGreen  = lipgloss.Color("#00ff99") // Green
	neonRed    = lipgloss.Color("#ff0066") // Red

	// Gradients and accents
	sunsetRed    = lipgloss.Color("#ff5555") // Sunset red
	sunsetOrange = lipgloss.Color("#ff8c42") // Sunset orange
	gridBlue     = lipgloss.Color("#29b6f6") // Grid blue
	hotPink      = lipgloss.Color("#fd3777") // Hot pink
)

// Cell styles
var (
	// UnrevealedStyle is used for cells that haven't been revealed yet
	UnrevealedStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(bgDark).
			Foreground(neonPurple)

	// RevealedStyle is used for cells that have been revealed
	RevealedStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(bgMedium).
			Foreground(neonBlue)

	// CursorStyle is used for the cell under the cursor
	CursorStyle = lipgloss.NewStyle().
			Width(2).
			Align(lipgloss.Center).
			Background(hotPink).
			Foreground(lipgloss.Color("black")).
			Bold(true)
)

// Number colors for the neighboring mine counts - synthwave style
var NumberStyles = []lipgloss.Style{
	lipgloss.NewStyle().Foreground(neonBlue),   // 1
	lipgloss.NewStyle().Foreground(neonGreen),  // 2
	lipgloss.NewStyle().Foreground(neonPink),   // 3
	lipgloss.NewStyle().Foreground(neonPurple), // 4
	lipgloss.NewStyle().Foreground(neonOrange), // 5
	lipgloss.NewStyle().Foreground(neonYellow), // 6
	lipgloss.NewStyle().Foreground(sunsetRed),  // 7
	lipgloss.NewStyle().Foreground(hotPink),    // 8
}

// Game status styles
var (
	// GameOverStyle is used for the game over message
	GameOverStyle = lipgloss.NewStyle().
			Foreground(sunsetRed).
			Background(bgDark).
			Bold(true).
			Padding(0, 1)

	// GameWonStyle is used for the victory message
	GameWonStyle = lipgloss.NewStyle().
			Foreground(neonGreen).
			Background(bgDark).
			Bold(true).
			Padding(0, 1)

	// TitleStyle is used for the game title
	TitleStyle = lipgloss.NewStyle().
			Foreground(neonPink).
			Bold(true).
			Underline(true).
			Padding(0, 1)

	// RetroSunsetStyle for borders and accents
	RetroSunsetStyle = lipgloss.NewStyle().
				Border(lipgloss.RoundedBorder()).
				BorderForeground(neonPink)
)
