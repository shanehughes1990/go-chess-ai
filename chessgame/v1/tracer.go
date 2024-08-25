package chessgame

// GameTracer is an interface that traces the game state to a backend.
type GameTracer interface {
	// Open opens the tracer.
	Open() error
	// Close closes the tracer.
	Close() error
	// ReadInGameState reads in the game state from the backend.
	ReadInGameState(game GameState) error
}

// NoOpTracer is a no-op default implimentation of GameTracer.
type NoOpTracer struct{}

// Open satisfies the GameTracer interface.
//
// This is a no-op.
func (t *NoOpTracer) Open() error {
	return nil
}

// Close satisfies the GameTracer interface.
//
// This is a no-op.
func (t *NoOpTracer) Close() error {
	return nil
}

// ReadInGameState satisfies the GameTracer interface.
//
// This is a no-op.
func (t *NoOpTracer) ReadInGameState(game GameState) error {
	return nil
}
