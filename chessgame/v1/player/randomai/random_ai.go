package randomai

import (
	"math/rand"

	"github.com/notnil/chess"
	"github.com/shanehughes1990/chess-ai/chessgame/v1"
)

// randomAI is a player that makes random moves.
type randomAI struct{ name string }

// NewRandomAI creates a new randomAI Player.
func NewRandomAI(name string) chessgame.Player {
	return &randomAI{name: name}
}

// Name returns the name of the player.
func (p *randomAI) Name() string {
	return p.name
}

// IsHuman returns true if the player is a human.
func (p *randomAI) IsHuman() bool {
	return false
}

// MakeMove makes a random move for the player.
func (p *randomAI) MakeMove(game *chessgame.GameState, xy ...int) (*chess.Move, error) {
	// Get all valid moves for the current player
	validMoves := game.Game().ValidMoves()
	if len(validMoves) == 0 {
		return nil, nil
	}

	// Choose a random move from the valid moves
	randomIndex := rand.Intn(len(validMoves))
	chosenMove := validMoves[randomIndex]

	return chosenMove, nil
}
