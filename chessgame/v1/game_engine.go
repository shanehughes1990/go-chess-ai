package chessgame

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v3"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/notnil/chess"
	"github.com/sirupsen/logrus"
)

// gameEngine is used internally to impliment the ebiten.Game interface.
type gameEngine struct{ *gameManager }

// newGameEngine assigns the gameEngine implimentation to the gameManager.
func (gm *gameManager) newGameEngine() {
	gm.gameEngine = &gameEngine{gm}
}

// Update impliments the ebiten.Game interface.
func (g *gameEngine) Update() error {
	logrus.Tracef("Current player: %s %s", g.gameState.currentPlayer.Name(), g.gameState.game.Position().Turn().String())
	// run the MakeMove function for the current player
	if err := g.makeMove(); err != nil {
		return err
	}

	return nil
}

// Draw impliments the ebiten.Game interface.
func (g *gameEngine) Draw(screen *ebiten.Image) {
	// Calculate the size of each square based on the screen dimensions
	g.gameState.squareSize = g.boardSize / 8
	// calculate the padding on either side of the screen
	g.gameState.padding = (screen.Bounds().Dx() - g.boardSize) / 2

	// draw the board to the screen
	g.drawBoard(screen)

	// draw the pieces to the screen
	g.drawPieces(screen)

	// draw the debug menu
	g.drawDebugMenu(screen)

	// draw the selected square and it's available moves
	g.drawHighlightSquare(screen)
	g.drawAvailableMoves(screen)
}

// Layout impliments the ebiten.Game interface.
func (g *gameEngine) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return 25% more width than boardSize
	return g.boardSize + (g.boardSize / 4), g.boardSize
}

// makeMove makes a move for the current player.
func (g *gameEngine) makeMove() error {
	move, err := g.gameState.currentPlayer.MakeMove(g.gameState)
	if err != nil {
		return err
	}

	if move != nil {
		if err := g.endTurn(move); err != nil {
			return err
		}
	}

	return nil
}

// endTurn ends the current players turn.
func (g *gameEngine) endTurn(move *chess.Move) error {
	// open the tracer
	if err := g.tracer.Open(); err != nil {
		return err
	}
	defer g.tracer.Close()

	// move the piece
	if err := g.gameState.game.Move(move); err != nil {
		return err
	}

	// trace the current game state
	if err := g.tracer.ReadInGameState(*g.gameState); err != nil {
		return err
	}

	// set the next player
	switch g.gameState.game.Position().Turn() {
	case chess.White:
		if g.gameState.whitePlayer == g.gameState.currentPlayer {
			logrus.Warnf("current player is already %s", g.gameState.currentPlayer.Name())
			return nil
		}

		g.gameState.currentPlayer = g.gameState.whitePlayer

	case chess.Black:
		if g.gameState.blackPlayer == g.gameState.currentPlayer {
			logrus.Warnf("current player is already %s", g.gameState.currentPlayer.Name())
			return nil
		}

		g.gameState.currentPlayer = g.gameState.blackPlayer
	}

	// unset the selected square
	g.gameState.unsetMove()

	return nil
}

// drawDebugMenu draws the debug menu to the screen.
func (g *gameEngine) drawDebugMenu(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.GeoM.Translate(10, 10)
	op.ColorScale.ScaleWithColor(color.Black)
	text.Draw(
		screen,
		fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()),
		text.NewGoXFace(bitmapfont.Face),
		op,
	)
	op.GeoM.Translate(0, 20)
	text.Draw(
		screen,
		fmt.Sprintf("FPS: %0.2f", ebiten.ActualFPS()),
		text.NewGoXFace(bitmapfont.Face),
		op,
	)
}

// drawAvailableMoves draws the available moves for the selected piece.
func (g *gameEngine) drawAvailableMoves(screen *ebiten.Image) {
	for _, move := range g.gameState.availableMoves {
		toSquare := move.S2()
		x, y := int(toSquare%8)*g.gameState.squareSize, int(toSquare/8)*g.gameState.squareSize
		// vector.DrawFilledRect(screen, float32(x), float32(y), float32(g.gameState.SquareSize), float32(g.gameState.SquareSize), g.availableMoveSquareColor, true)

		// Add padding to the x coordinate
		x += g.gameState.padding

		lineWidth := 5 // Adjust the line width as needed

		// Draw the four sides of the square using thin rectangles
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(g.gameState.squareSize), float32(lineWidth), g.availableMoveSquareColor, true)                                  // Top
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(lineWidth), float32(g.gameState.squareSize), g.availableMoveSquareColor, true)                                  // Left
		vector.DrawFilledRect(screen, float32(x+g.gameState.squareSize-lineWidth), float32(y), float32(lineWidth), float32(g.gameState.squareSize), g.availableMoveSquareColor, true) // Right
		vector.DrawFilledRect(screen, float32(x), float32(y+g.gameState.squareSize-lineWidth), float32(g.gameState.squareSize), float32(lineWidth), g.availableMoveSquareColor, true) // Bottom
	}
}

// drawHighlightSquare draws a highlight on the selected square.
func (g *gameEngine) drawHighlightSquare(screen *ebiten.Image) {
	if g.gameState.selectedSquareX >= 0 && g.gameState.selectedSquareY >= 0 {
		selectedSquare := chess.Square(g.gameState.selectedSquareY*8 + g.gameState.selectedSquareX)
		piece := g.gameState.game.Position().Board().Piece(selectedSquare)

		// highlight the selected square if it is the current players turn
		if piece != chess.NoPiece && piece.Color() == g.gameState.game.Position().Turn() {
			x, y := g.gameState.selectedSquareX*g.gameState.squareSize, g.gameState.selectedSquareY*g.gameState.squareSize

			// Add padding to the x coordinate
			x += g.gameState.padding

			lineWidth := 5 // Adjust the line width as needed

			// Draw the four sides of the square using thin rectangles
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(g.gameState.squareSize), float32(lineWidth), g.highlightSquareColor, true)                                  // Top
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(lineWidth), float32(g.gameState.squareSize), g.highlightSquareColor, true)                                  // Left
			vector.DrawFilledRect(screen, float32(x+g.gameState.squareSize-lineWidth), float32(y), float32(lineWidth), float32(g.gameState.squareSize), g.highlightSquareColor, true) // Right
			vector.DrawFilledRect(screen, float32(x), float32(y+g.gameState.squareSize-lineWidth), float32(g.gameState.squareSize), float32(lineWidth), g.highlightSquareColor, true) // Bottom
		}
	}
}

// drawBoard draws the chess board to the screen.
func (g *gameEngine) drawBoard(screen *ebiten.Image) {

	// Iterate over the board and draw the squares
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x, y := i*g.gameState.squareSize, j*g.gameState.squareSize

			// adjust the x coordinate based on the padding
			x += g.gameState.padding

			// Determine the color based on the row and column
			var fillColor color.Color
			if (i+j)%2 == 0 { // light squares on any even row and column
				fillColor = g.lightSquareColor
			} else {
				fillColor = g.darkSquareColor
			}

			// Draw the square
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(g.gameState.squareSize), float32(g.gameState.squareSize), fillColor, true)
		}
	}
}

// drawPieces draws the chess pieces to the screen.
func (g *gameEngine) drawPieces(screen *ebiten.Image) {
	board := g.gameState.game.Position().Board() // get the current board state
	for i := 0; i < 8; i++ {                     // Iterate over ranks (rows)
		for j := 0; j < 8; j++ { // Iterate over files (columns)
			square := chess.Square(i*8 + j)
			piece := board.Piece(square)

			if piece != chess.NoPiece {
				// Calculate x and y, keeping white on top
				x := (j * g.gameState.squareSize) + g.gameState.padding
				y := i * g.gameState.squareSize

				// Center the piece image within the square
				pieceImg := g.pieceImages[piece]
				pieceWidth, pieceHeight := pieceImg.Bounds().Dx(), pieceImg.Bounds().Dy()
				offsetX := (g.gameState.squareSize - pieceWidth) / 2
				offsetY := (g.gameState.squareSize - pieceHeight) / 2

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x+offsetX), float64(y+offsetY))
				screen.DrawImage(pieceImg, op)
			}
		}
	}
}
