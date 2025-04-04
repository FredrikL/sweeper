// Package main provides a Minesweeper game using the Bubble Tea TUI framework.
package main

import (
	"math/rand"
	"time"
)

// Game dimensions and difficulty
const (
	BoardWidth  = 10
	BoardHeight = 10
	MineCount   = 15
)

// Cell represents a single cell in the Minesweeper board
type Cell struct {
	IsMine     bool
	IsRevealed bool
	IsFlagged  bool
	Neighbors  int
}

// Creates and initializes a new game board
func createBoard() [][]Cell {
	board := make([][]Cell, BoardHeight)
	for i := range board {
		board[i] = make([]Cell, BoardWidth)
	}

	// Place mines randomly
	rand.Seed(time.Now().UnixNano())
	minesPlaced := 0
	for minesPlaced < MineCount {
		x := rand.Intn(BoardWidth)
		y := rand.Intn(BoardHeight)
		if !board[y][x].IsMine {
			board[y][x].IsMine = true
			minesPlaced++
		}
	}

	// Calculate neighbors
	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			if !board[y][x].IsMine {
				count := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dx == 0 && dy == 0 {
							continue
						}
						nx, ny := x+dx, y+dy
						if nx >= 0 && nx < BoardWidth && ny >= 0 && ny < BoardHeight && board[ny][nx].IsMine {
							count++
						}
					}
				}
				board[y][x].Neighbors = count
			}
		}
	}

	return board
}

// Reveals a cell and handles cascading reveals for empty cells
func revealCell(board [][]Cell, x, y int) (bool, bool) {
	if board[y][x].IsRevealed || board[y][x].IsFlagged {
		return false, false
	}

	board[y][x].IsRevealed = true

	// Check if a mine was revealed
	if board[y][x].IsMine {
		return true, false
	}

	// For empty cells, reveal neighbors
	if board[y][x].Neighbors == 0 {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				nx, ny := x+dx, y+dy
				if nx >= 0 && nx < BoardWidth && ny >= 0 && ny < BoardHeight {
					revealCell(board, nx, ny)
				}
			}
		}
	}

	// Check for win condition
	return false, checkWin(board)
}

// Toggle flag on a cell
func toggleFlag(board [][]Cell, x, y int) {
	if !board[y][x].IsRevealed {
		board[y][x].IsFlagged = !board[y][x].IsFlagged
	}
}

// Check if the player has won
func checkWin(board [][]Cell) bool {
	unrevealed := 0
	for y := 0; y < BoardHeight; y++ {
		for x := 0; x < BoardWidth; x++ {
			if !board[y][x].IsRevealed && !board[y][x].IsMine {
				unrevealed++
			}
		}
	}
	return unrevealed == 0
}
