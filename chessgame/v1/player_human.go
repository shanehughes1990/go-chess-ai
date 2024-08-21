package chessgame

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/notnil/chess"
	"github.com/thoas/go-funk"
)

// humanPlayer impliments the Player interface when the player is human.
type humanPlayer struct{ name string }

// NewHumanPlayer creates a new humanPlayer.
func NewHumanPlayer(name string) Player {
	return &humanPlayer{name: name}
}

// Name returns the name of the player.
func (p *humanPlayer) Name() string {
	return p.name
}

// MakeMove makes a move based on user input.
func (p *humanPlayer) MakeMove(game *GameState) (*chess.Move, error) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// adjust the x coordinate based on the padding
		x -= game.padding

		// calculate the square coordinates based on the mouse click
		clickedSquareX := x / game.squareSize
		clickedSquareY := y / game.squareSize

		// deselect if outside the board
		if x < 0 || x >= 8*game.squareSize || y < 0 || y >= 8*game.squareSize {
			game.unsetMove()
			return nil, nil
		}

		// deselect if the same square is clicked twice
		if game.selectedSquareX == clickedSquareX && game.selectedSquareY == clickedSquareY {
			game.unsetMove()
			return nil, nil
		}

		// update the selected square state
		game.selectedSquareX, game.selectedSquareY = clickedSquareX, clickedSquareY

		// check if the clicked square is a valid move
		if game.selectedSquareX >= 0 && game.selectedSquareY >= 0 {
			selectedSquare := chess.Square(game.selectedSquareY*8 + game.selectedSquareX)
			clickedSquare := chess.Square(clickedSquareY*8 + clickedSquareX)

			// get the available moves for the selected square
			piece := game.Game().Position().Board().Piece(selectedSquare)
			if piece != chess.NoPiece && piece.Color() == game.Game().Position().Turn() {
				game.availableMoves = funk.Filter(
					game.Game().ValidMoves(),
					func(move *chess.Move) bool { return move.S1() == selectedSquare },
				).([]*chess.Move)
			}

			// return the move if we clicked on a valid square for the selected piece
			for _, move := range game.availableMoves {
				if move.S2() == clickedSquare {
					return move, nil
				}
			}
		}
	}

	return nil, nil
}
