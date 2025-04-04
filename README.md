# Sweeper - A Terminal Minesweeper Game

A colorful, mouse-supported Terminal UI implementation of the classic Minesweeper game built with Go and the Bubble Tea framework.

## Demo

<!-- YouTube video embed with high-quality thumbnail -->
<a href="https://www.youtube.com/watch?v=JeNS1ZNHQs8" target="_blank">
  <img src="https://img.youtube.com/vi/JeNS1ZNHQs8/maxresdefault.jpg" alt="Minesweeper Game Demo" width="600" />
</a>

*Click the image above to watch the demo video*

## Features

- 🎮 Full terminal UI with mouse support
- 🎨 Beautiful colors and styling
- 🧠 Classic Minesweeper gameplay
- 🖱️ Both keyboard and mouse controls

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

- **Arrow keys / hjkl**: Move cursor
- **Space / Enter**: Reveal cell
- **f**: Toggle flag
- **Left Click**: Reveal cell
- **Right Click**: Toggle flag
- **q**: Quit the game

## Technical Details

This game is built using:

- [Go](https://golang.org/)
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - A TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Style definitions for terminal applications

## License

MIT 