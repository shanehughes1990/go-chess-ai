package gamemanager

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/notnil/chess"
	"github.com/shanehughes1990/chess-ai/pkg/pointer"
	"github.com/sirupsen/logrus"
)

// GameManager is a struct that holds the game state.
type GameManager struct {
	debugMenu                                         bool
	tps                                               int
	vsyncEnabled                                      bool
	runnableOnUnfocused                               bool
	boardSize                                         int
	windowWidth, windowHeight                         int
	squareSize                                        int
	lightSquareColor, darkSquareColor, highlightColor color.Color
	pieceImages                                       map[chess.Piece]*ebiten.Image // Map of the loaded piece images
	state                                             State
}

type State struct {
	game                             *chess.Game
	availableMoves                   []*chess.Move
	player1, player2                 Player
	selectedSquareX, selectedSquareY int
}

// NewGameManager creates a new GameManager instance.
func NewGameManager(opts ...GameManagerOption) (GameManager, error) {
	gm := &GameManager{
		debugMenu:           false,
		tps:                 ebiten.DefaultTPS,
		vsyncEnabled:        true,
		runnableOnUnfocused: true,
		boardSize:           800,
		windowWidth:         960,
		windowHeight:        800,
		lightSquareColor:    color.RGBA{0xEE, 0xEE, 0xEE, 0xFF}, // Light gray
		darkSquareColor:     color.RGBA{0x99, 0x99, 0x99, 0xFF}, // Dark gray
		highlightColor:      color.RGBA{255, 215, 0, 255},       // Gold yellow
		pieceImages:         make(map[chess.Piece]*ebiten.Image),
		state:               State{},
	}

	// Initialize the game state
	gm.restartGameState()

	// Apply the options
	for _, opt := range opts {
		opt(gm)
	}

	// Load the piece images
	for piece, path := range PieceImage {
		img, _, err := ebitenutil.NewImageFromFile(path)
		if err != nil {
			return GameManager{}, err
		}

		gm.pieceImages[piece] = img
	}

	// calculate the square size
	gm.squareSize = gm.boardSize / 8

	return pointer.Deref(gm), nil
}

// Update is called every tick of the game loop.
func (gm *GameManager) Update() error {
	// debug overlay hotkey
	logrus.Debugf("Current Player: %s, Game Turn: %s", gm.getCurrentPlayer().Name(), gm.state.game.Position().Turn().String())
	gm.updateKeyRJustPressed()

	// check if the game is over
	if _, gameOver := gm.gameover(); gameOver {
		gm.updateRestartButtonJustPressed()
		return nil
	}

	// make the current player's move
	move := gm.getCurrentPlayer().MakeMove(gm.state.game, gm)
	if move != nil {
		if err := gm.endTurn(move); err != nil {
			return err
		}
	}

	return nil
}

// Draw is called every frame of the game loop.
func (gm *GameManager) Draw(screen *ebiten.Image) {
	// Calculate the size of each square based on the screen dimensions
	// draw the board
	gm.drawBoard(screen, gm.squareSize)

	// draw the pieces
	gm.drawPieces(screen, gm.squareSize)

	// draw the game over screen
	gm.drawRestartButton(screen)

	// highlight the selected square
	if gm.state.selectedSquareX >= 0 && gm.state.selectedSquareY >= 0 {
		gm.drawHighlightSquare(screen, gm.squareSize)
	}

	// highlight the available moves
	gm.drawAvailableMoves(screen, gm.squareSize)

	// You can add other rendering elements here, like pieces, UI, etc.
	if gm.debugMenu {
		gm.drawDebugInfo(screen)
	}
}

// Layout is run every time the window is resized.
// TODO: Implement this function to handle window resizing.
func (gm *GameManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 800, 800
}

// StartGame starts the game.
func (gm *GameManager) StartGame() error {
	// Set the game engine options
	ebiten.SetWindowTitle("Chess AI")
	ebiten.SetWindowSize(gm.windowWidth, gm.windowHeight)
	ebiten.SetTPS(gm.tps)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(gm.vsyncEnabled)
	ebiten.SetRunnableOnUnfocused(gm.runnableOnUnfocused)

	// Start the game
	return ebiten.RunGame(gm)
}
