package chessgame

import "github.com/notnil/chess"

// Player represents a player in the game.
type Player interface {
	// Name returns the name of the player.
	Name() string
	// MakeMove is the method that decides the move for the player.
	//
	// Returning a move will cause the game engine to finalize the move.
	//
	// Returning nil will cause the game engine to wait for the player to make a move.
	MakeMove(game *GameState) (*chess.Move, error)
}
