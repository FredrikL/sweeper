// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"github.com/charmbracelet/lipgloss"
)

// renderVictoryScreen creates and returns a celebratory victory screen with synthwave theme
func renderVictoryScreen() string {
	// Create a synthwave victory screen
	victoryStyle := lipgloss.NewStyle().
		Foreground(neonGreen).
		Bold(true).
		Background(bgDark).
		Padding(1, 2).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(neonPink)

	// Generate synthwave grid lines and decorations
	grid := lipgloss.NewStyle().Foreground(gridBlue).Render
	sun := lipgloss.NewStyle().Foreground(sunsetOrange).Render
	neonAccent := lipgloss.NewStyle().Foreground(neonPurple).Render

	// Synthwave trophy/sun
	trophyLines := []string{
		"         .         ",
		"      \\  |  /      ",
		"       \\ | /       ",
		"     ----⦿----    ",
		"       / | \\       ",
		"      /  |  \\      ",
		"                   ",
		"   /\\         /\\   ",
		"  /  \\_______/  \\  ",
		" /               \\ ",
		"/_________________\\",
		"                   ",
		"  ┌─────────────┐  ",
		"  │    WINNER   │  ",
		"  └─────────────┘  ",
	}

	// Create synthwave grid border with neon accents
	gridLine := ""
	for i := 0; i < 42; i++ {
		if i%6 == 0 {
			gridLine += neonAccent("•")
		} else {
			gridLine += grid("·")
		}
	}

	// Horizontal grid line
	horizon := ""
	for i := 0; i < 42; i++ {
		horizon += sun("_")
	}

	trophyStr := ""
	for _, line := range trophyLines {
		trophyStr += neonAccent("│ ") + line + neonAccent(" │") + "\n"
	}

	message := lipgloss.NewStyle().
		Foreground(neonPink).
		Bold(true).
		Width(40).
		Align(lipgloss.Center).
		Render("~ RADICAL VICTORY ~")

	subMessage := lipgloss.NewStyle().
		Foreground(neonBlue).
		Bold(true).
		Render("All mines defused successfully!")

	controls := lipgloss.NewStyle().
		Foreground(bgLight).
		Render("Press 'q' to quit")

	return victoryStyle.Render(
		gridLine + "\n" +
			horizon + "\n\n" +
			message + "\n\n" +
			lipgloss.NewStyle().Foreground(neonYellow).Render(trophyStr) + "\n" +
			subMessage + "\n\n" +
			controls + "\n\n" +
			horizon + "\n" +
			gridLine)
}
