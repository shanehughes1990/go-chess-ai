package gamemanager

import (
	"math/rand"

	"github.com/notnil/chess"
)

// AIPlayer is a player that makes random moves.
type AIPlayer struct{ name string }

// NewAIPlayer creates a new AIPlayer.
func NewAIPlayer(name string) *AIPlayer {
	return &AIPlayer{name: name}
}

func (p *AIPlayer) MakeMove(game *chess.Game, gm *GameManager) *chess.Move {
	// Get all valid moves for the current player
	validMoves := game.ValidMoves()
	if len(validMoves) == 0 {
		return nil
	}

	// Choose a random move from the valid moves
	randomIndex := rand.Intn(len(validMoves))
	chosenMove := validMoves[randomIndex]

	return chosenMove
}

func (p *AIPlayer) IsHuman() bool {
	return false
}

func (p *AIPlayer) Name() string {
	return "AI"
}
