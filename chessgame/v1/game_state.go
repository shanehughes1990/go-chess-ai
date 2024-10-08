package chessgame

import "github.com/notnil/chess"

// GameState holds the state of the game.
type GameState struct {
	game                                    *chess.Game
	squareSize                              int
	padding                                 int
	availableMoves                          []*chess.Move
	selectedSquareX, selectedSquareY        int
	whitePlayer, blackPlayer, currentPlayer Player
}

// newGameState creates a new gameState.
//
// Will set the game to a new chess game.
// If player1 is not nil, it will set player1.
// If player2 is not nil, it will set player2.
func (gm *gameManager) newGameState() {
	gameState := &GameState{
		game:            chess.NewGame(),
		squareSize:      gm.boardSize / 8,
		padding:         0,
		availableMoves:  nil,
		selectedSquareX: -1,
		selectedSquareY: -1,
	}

	// if we have no state (new game), set the gameState and return
	if gm.gameState == nil {
		gm.gameState = gameState
		return
	}

	// if there's an existing state, set the players from the previous state and overwrite the game
	// var player1, player2 Player
	if gm.gameState.whitePlayer != nil {
		gameState.whitePlayer = gm.gameState.whitePlayer
	}

	if gm.gameState.blackPlayer != nil {
		gameState.blackPlayer = gm.gameState.blackPlayer
	}

	// set the current player
	gameState.currentPlayer = gameState.whitePlayer

	// set the gameState
	gm.gameState = gameState
}

// GetSquareSize returns the size of each square on the board.
func (g *GameState) GetSquareSize() int {
	return g.squareSize
}

// Game returns the current game.
func (g *GameState) Game() *chess.Game {
	return g.game
}

// UnsetMove unsets the selected move.
func (g *GameState) UnsetMove() {
	g.selectedSquareX, g.selectedSquareY, g.availableMoves = -1, -1, nil
}

// GetSelectedSquare returns the selected square.
func (g *GameState) GetSelectedSquare() (x int, y int) {
	return g.selectedSquareX, g.selectedSquareY
}

// SetSelectedSquare sets the selected square.
func (g *GameState) SetSelectedSquare(x, y int) {
	g.selectedSquareX, g.selectedSquareY = x, y
}

// GetAvailableMoves returns the available moves.
func (g *GameState) GetAvailableMoves() []*chess.Move {
	return g.availableMoves
}

// SetAvailableMoves sets the available moves.
func (g *GameState) SetAvailableMoves(moves []*chess.Move) {
	g.availableMoves = moves
}
