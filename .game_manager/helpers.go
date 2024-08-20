package gamemanager

import (
	"github.com/notnil/chess"
	"github.com/sirupsen/logrus"
)

// gameover checks if the game is over.
//
// this checks if the game is over by checking if the game is in checkmate or stalemate.
func (gm *GameManager) gameover() (chess.Outcome, bool) {
	return gm.state.game.Outcome(), gm.state.game.Outcome() != chess.NoOutcome
}

// restartGame restarts the game.
//
// this restarts the game by resetting the game state.
func (gm *GameManager) restartGameState() {
	newState := State{
		game:            chess.NewGame(),
		availableMoves:  []*chess.Move{},
		selectedSquareX: -1,
		selectedSquareY: -1,
	}

	if gm.state.player1 != nil {
		newState.player1 = gm.state.player1
	}

	if gm.state.player2 != nil {
		newState.player2 = gm.state.player2
	}

	// set the new state
	gm.state = newState
}

// getCurrentPlayer returns the current player.
func (gm *GameManager) getCurrentPlayer() Player {
	if gm.state.game.Position().Turn() == chess.White {
		return gm.state.player1
	}

	return gm.state.player2
}

// unsetSelectedSquare unsets the selected square.
//
// this resets the selected piece to -1 -1.
func (gm *GameManager) unsetSelectedSquare() {
	gm.state.selectedSquareX, gm.state.selectedSquareY = -1, -1
	gm.state.availableMoves = nil
}

// endTurn ends the current player's turn.
//
// this ends the current player's turn and switches to the next player.
func (gm *GameManager) endTurn(move *chess.Move) error {
	logrus.Infof("Ending turn for %s", gm.getCurrentPlayer().Name())
	err := gm.state.game.Move(move)
	if err != nil {
		// Handle invalid move (e.g., display an error message)
		return err
	}

	gm.unsetSelectedSquare()
	return nil
}
