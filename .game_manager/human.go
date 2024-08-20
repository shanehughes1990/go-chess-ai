package gamemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/notnil/chess"
	"github.com/thoas/go-funk"
)

// HumanPlayer is a struct that represents a human player.
type HumanPlayer struct{ name string }

// NewHumanPlayer creates a new HumanPlayer instance.
func NewHumanPlayer(name string) HumanPlayer {
	return HumanPlayer{name: name}
}

// Update is called every tick of the game loop.
func (p HumanPlayer) MakeMove(game *chess.Game, gm *GameManager) *chess.Move {
	if funk.All(
		funk.Equal(gm.getCurrentPlayer().IsHuman(), true),
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft),
	) {
		x, y := ebiten.CursorPosition()

		// Calculate the square coordinates based on mouse position
		clickedSquareX := x / gm.squareSize
		clickedSquareY := y / gm.squareSize

		// deselect if the same square is clicked twice
		if funk.All(
			funk.Equal(gm.state.selectedSquareX, clickedSquareX),
			funk.Equal(gm.state.selectedSquareY, clickedSquareY),
			funk.NotEqual(gm.state.selectedSquareX, -1),
			funk.NotEqual(gm.state.selectedSquareY, -1),
		) {
			gm.unsetSelectedSquare()
			return nil
		}

		// Update the selected square
		gm.state.selectedSquareX = clickedSquareX
		gm.state.selectedSquareY = clickedSquareY

		// deselect if outside the board
		if gm.state.selectedSquareX < 0 || gm.state.selectedSquareX >= 8 || gm.state.selectedSquareY < 0 || gm.state.selectedSquareY >= 8 {
			gm.unsetSelectedSquare()
			return nil
		}

		// Check if the clicked square is a valid move
		if gm.state.selectedSquareX >= 0 && gm.state.selectedSquareY >= 0 {
			selectedSquare := chess.Square(gm.state.selectedSquareY*8 + gm.state.selectedSquareX)
			clickedSquare := chess.Square(clickedSquareY*8 + clickedSquareX)

			for _, move := range gm.state.availableMoves {
				if move.S2() == clickedSquare {
					return move // Return the valid move
				}
			}

			// Get available moves for the newly selected piece
			piece := game.Position().Board().Piece(selectedSquare)
			if piece != chess.NoPiece && piece.Color() == game.Position().Turn() {
				gm.state.availableMoves = funk.Filter(
					gm.state.game.ValidMoves(),
					func(move *chess.Move) bool {
						return move.S1() == selectedSquare
					},
				).([](*chess.Move))
			} else {
				gm.state.availableMoves = nil // Clear available moves if no piece or wrong turn
			}
		}
	}

	return nil // No valid move selected yet
}

// IsHuman returns true if the player is human
func (hp HumanPlayer) IsHuman() bool {
	return true
}

// Name returns the name of the player
func (hp HumanPlayer) Name() string {
	return hp.name
}
