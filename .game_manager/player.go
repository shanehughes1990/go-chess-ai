package gamemanager

import (
	"github.com/notnil/chess"
)

type Player interface {
	// MakeMove is called to make a move for the player.
	MakeMove(game *chess.Game, gm *GameManager) *chess.Move
	// IsHuman returns true if the player is human
	IsHuman() bool
	// Name
	Name() string
}
