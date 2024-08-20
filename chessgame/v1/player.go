package chessgame

import "github.com/notnil/chess"

// Player represents a player in the game.
type Player interface {
	// Name returns the name of the player.
	Name() string
	// IsHuman returns true if the player is a human.
	IsHuman() bool
	// MakeMove
	//
	// When the player IsHuman, the xy coordinates will be passed to the function,
	// based off the mouse left click
	MakeMove(game *GameState, xy ...int) (*chess.Move, error)
}
