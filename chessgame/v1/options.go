package chessgame

// GameManagerOption is a func that sets an option on a ChessGame.
type GameManagerOption func(*gameManager)

// WithWhitePlayer sets the white player.
func WithWhitePlayer(player Player) GameManagerOption {
	return func(cg *gameManager) {
		cg.gameState.whitePlayer = player
	}
}

// WithBlackPlayer sets the black player.
func WithBlackPlayer(player Player) GameManagerOption {
	return func(cg *gameManager) {
		cg.gameState.blackPlayer = player
	}
}
