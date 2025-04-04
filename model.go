// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	tea "github.com/charmbracelet/bubbletea"
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
