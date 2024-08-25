package filetracer

import (
	"os"

	"github.com/shanehughes1990/chess-ai/chessgame/v1"
)

// FileTracer implements the GameTracer interface.
type FileTracer struct {
	filename string
	*os.File
}

// NewFileTracer initializes a new FileTracer.
func NewFileTracer(filename string) chessgame.GameTracer {
	return &FileTracer{filename: filename}
}

// Open opens the tracer.
func (t *FileTracer) Open() error {
	// Open the file
	f, err := os.OpenFile(t.filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	t.File = f
	return nil
}

// ReadInGameState reads in the game state from the backend.
func (t *FileTracer) ReadInGameState(game chessgame.GameState) error {
	// write the current game state to a new line in the file
	_, err := t.File.WriteString(game.Game().Position().Board().String() + "\n")
	if err != nil {
		return err
	}

	return nil
}

// Close closes the tracer.
func (t *FileTracer) Close() error {
	return t.File.Close()
}
