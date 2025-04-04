// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"github.com/charmbracelet/lipgloss"
)

// renderVictoryScreen creates and returns a celebratory victory screen
func renderVictoryScreen() string {
	// Create a celebratory victory screen
	victoryStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("46")).
		Bold(true).
		Background(lipgloss.Color("236")).
		Padding(1, 3).
		Align(lipgloss.Center).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("39"))

	// Generate better sparkles that render in all terminals
	sparkle := lipgloss.NewStyle().Foreground(lipgloss.Color("226")).Render("*")
	star := lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Render("*")
	confetti1 := lipgloss.NewStyle().Foreground(lipgloss.Color("201")).Render(".")
	confetti2 := lipgloss.NewStyle().Foreground(lipgloss.Color("51")).Render("·")

	// New simpler trophy that renders better
	trophyLines := []string{
		"       ___       ",
		"     .'   '.     ",
		"    /       \\    ",
		"    |       |    ",
		"    |  WIN  |    ",
		"    |       |    ",
		"    |       |    ",
		"    '._____.'    ",
		"    /       \\    ",
		"   /         \\   ",
		"  /___________\\  ",
		"       | |       ",
		"      _|_|_      ",
		"     |_____|     ",
	}

	// Add sparkles to the trophy
	trophy := []string{}
	for _, line := range trophyLines {
		trophy = append(trophy,
			confetti1+" "+sparkle+line+sparkle+" "+confetti2)
	}

	// Create sparkle border with simpler characters that render in all terminals
	border := ""
	for i := 0; i < 38; i++ {
		if i%4 == 0 {
			border += sparkle
		} else if i%4 == 1 {
			border += confetti1
		} else if i%4 == 2 {
			border += star
		} else {
			border += confetti2
		}
		if i < 37 {
			border += " "
		}
	}

	trophyStr := ""
	for _, line := range trophy {
		trophyStr += line + "\n"
	}

	message := lipgloss.NewStyle().
		Foreground(lipgloss.Color("226")).
		Bold(true).
		Width(30).
		Align(lipgloss.Center).
		Render("* VICTORY! *")

	subMessage := lipgloss.NewStyle().
		Foreground(lipgloss.Color("255")).
		Bold(true).
		Render("All mines successfully found!")

	controls := lipgloss.NewStyle().
		Foreground(lipgloss.Color("242")).
		Render("Press 'q' to quit")

	// Add sparkle border at top and bottom
	return victoryStyle.Render(
		border + "\n\n" +
			message + "\n\n" +
			lipgloss.NewStyle().Foreground(lipgloss.Color("220")).Render(trophyStr) + "\n" +
			subMessage + "\n\n" +
			controls + "\n\n" +
			border)
}
