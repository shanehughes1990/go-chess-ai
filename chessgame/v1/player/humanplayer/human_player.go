package humanplayer

import (
	"fmt"

	"github.com/notnil/chess"
	"github.com/shanehughes1990/chess-ai/chessgame/v1"
	"github.com/thoas/go-funk"
)

// humanPlayer is a player that makes moves based on user input.
type humanPlayer struct{ name string }

// NewHumanPlayer creates a new humanPlayer.
func NewHumanPlayer(name string) chessgame.Player {
	return &humanPlayer{name: name}
}

// Name returns the name of the player.
func (p *humanPlayer) Name() string {
	return p.name
}

// IsHuman returns true if the player is a human.
func (p *humanPlayer) IsHuman() bool {
	return true
}

// MakeMove makes a move based on user input.
func (p *humanPlayer) MakeMove(game *chessgame.GameState, xy ...int) (*chess.Move, error) {
	// x, y := ebiten.CursorPosition()
	if xy == nil || len(xy) != 2 {
		return nil, fmt.Errorf("invalid coordinates: %v", xy)
	}

	x, y := xy[0], xy[1]

	// calculate the square coordinates based on the mouse click
	clickedSquareX := x / game.GetSquareSize()
	clickedSquareY := y / game.GetSquareSize()

	// deselect if outside the board
	if x < 0 || x >= 8*game.GetSquareSize() || y < 0 || y >= 8*game.GetSquareSize() {
		game.UnsetMove()
		return nil, nil
	}

	// deselect if the same square is clicked twice
	if selectedSquareX, selectedSquareY := game.GetSelectedSquare(); selectedSquareX == clickedSquareX && selectedSquareY == clickedSquareY {
		game.UnsetMove()
		return nil, nil
	}

	// update the selected square state
	game.SetSelectedSquare(clickedSquareX, clickedSquareY)

	// check if the clicked square is a valid move
	if selectedSquareX, selectedSquareY := game.GetSelectedSquare(); selectedSquareX >= 0 && selectedSquareY >= 0 {
		selectedSquare := chess.Square(selectedSquareY*8 + selectedSquareX)
		clickedSquare := chess.Square(clickedSquareY*8 + clickedSquareX)

		// get the available moves for the selected square
		piece := game.Game().Position().Board().Piece(selectedSquare)
		if piece != chess.NoPiece && piece.Color() == game.Game().Position().Turn() {
			game.SetAvailableMoves(funk.Filter(
				game.Game().ValidMoves(),
				func(move *chess.Move) bool { return move.S1() == selectedSquare },
			).([]*chess.Move))
		}

		// return the move if we clicked on a valid square for the selected piece
		for _, move := range game.GetAvailableMoves() {
			if move.S2() == clickedSquare {
				return move, nil
			}
		}
	}

	return nil, nil
}
