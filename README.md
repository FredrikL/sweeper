<a href="https://www.youtube.com/watch?v=JeNS1ZNHQs8" target="_blank">
  <img src="https://img.youtube.com/vi/JeNS1ZNHQs8/maxresdefault.jpg" alt="Minesweeper Game Demo" width="600" />
</a>

# Sweeper - A Terminal Minesweeper Game

A colorful, **fully mouse-supported** Terminal UI implementation of the classic Minesweeper game built with Go and the Bubble Tea framework.

## Features

- 🖱️ **Full mouse support** - click to reveal cells, right-click to flag
- 🎮 Rich terminal UI with vibrant colors and animations
- 🎨 Beautiful victory screen with sparkles and trophy
- 🧠 Classic Minesweeper gameplay with intuitive controls
- 💣 All mines revealed on game over

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/sweeper.git
cd sweeper

# Build the game
go build

# Run the game
./sweeper
```

## How to Play

The goal is to reveal all cells that don't contain mines. The numbers indicate how many mines are adjacent to that cell.

### Controls

- **Mouse**:
  - **Left Click**: Reveal cell
  - **Right Click**: Toggle flag

- **Keyboard**:
  - **Arrow keys / hjkl**: Move cursor
  - **Space / Enter**: Reveal cell
  - **f**: Toggle flag
  - **q**: Quit the game

## Technical Details

This game is built using:

- [Go](https://golang.org/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - A TUI framework with mouse support
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal applications

## License

MIT 