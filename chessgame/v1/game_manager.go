package chessgame

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/notnil/chess"
)

// GameManager is the main entry for the chess game engine.
//
// It is responsible for managing the game state and running the game loop.
type GameManager interface {
	// Start runs the game loop.
	Start() error
}

// gameManager is the implementation of the chess game.
type gameManager struct {
	boardSize                                                                         int
	lightSquareColor, darkSquareColor, highlightSquareColor, availableMoveSquareColor color.Color
	pieceImages                                                                       map[chess.Piece]*ebiten.Image
	gameState                                                                         *GameState
	gameEngine                                                                        ebiten.Game
}

// NewGameManager creates a new GameManager instance.
func NewGameManager(opts ...GameManagerOption) GameManager {
	chessGame := &gameManager{
		boardSize:                800,                                 // default board size
		lightSquareColor:         color.RGBA{R: 255, G: 206, B: 158},  // default light square color, light brown
		darkSquareColor:          color.RGBA{R: 209, G: 139, B: 71},   // default dark square color, dark brown
		highlightSquareColor:     color.RGBA{R: 255, G: 215, B: 0},    // default highlight square color, gold yellow
		availableMoveSquareColor: color.RGBA{0, 255, 0, 128},          // default available move square color, green
		pieceImages:              make(map[chess.Piece]*ebiten.Image), // map of the loaded piece images
		gameState:                &GameState{},                        // initialize the game state
	}

	// apply the options
	for _, opt := range opts {
		opt(chessGame)
	}

	// initialize the internal game struct and state
	chessGame.newGameEngine()
	chessGame.newGameState()

	return chessGame
}

// Start runs the game loop.
func (gm *gameManager) Start() error {
	// load the piece images here so we don't have to return an error on NewGameManager
	for piece, path := range pieceMap {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return err
		}

		gm.pieceImages[piece] = img
	}

	return ebiten.RunGame(gm.gameEngine)
}
