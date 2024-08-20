package gamemanager

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

// drawBoard draws the chess board to the screen.
//
// this renders the chess board to the screen using the light and dark square colors.
func (gm *GameManager) drawBoard(screen *ebiten.Image, squareSize int) {
	// Iterate over the board and draw the squares
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			x, y := i*squareSize, j*squareSize

			// Determine the color based on the row and column
			var fillColor color.Color
			if (i+j)%2 == 0 { // light squares on any even row and column
				fillColor = gm.lightSquareColor
			} else {
				fillColor = gm.darkSquareColor
			}

			// Draw the square
			vector.DrawFilledRect(screen, float32(x), float32(y), float32(squareSize), float32(squareSize), fillColor, true)
		}
	}
}

// drawDebugInfo draws debug information to the screen.
//
// this renders an overlay on the screen with debug information such as TPS and FPS.
func (gm *GameManager) drawDebugInfo(screen *ebiten.Image) {
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

// drawHighlightSquare draws a highlight on the selected square.
//
// this renders a box around the selected square to highlight it.
func (gm *GameManager) drawHighlightSquare(screen *ebiten.Image, squareSize int) {
	logrus.Debugf("drawHightlightSquare: x=%d, y=%d", gm.state.selectedSquareX, gm.state.selectedSquareY)
	if gm.state.selectedSquareX == -1 || gm.state.selectedSquareY == -1 {
		return
	}

	// Get the piece on the selected square
	selectedSquare := chess.Square(gm.state.selectedSquareY*8 + gm.state.selectedSquareX)
	piece := gm.state.game.Position().Board().Piece(selectedSquare)

	// Only highlight if the piece belongs to the current player's turn
	if piece != chess.NoPiece && piece.Color() == gm.state.game.Position().Turn() {
		x, y := gm.state.selectedSquareX*squareSize, gm.state.selectedSquareY*squareSize
		lineWidth := 5 // Adjust the line width as needed

		// Draw the four sides of the square using thin rectangles
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(squareSize), float32(lineWidth), gm.highlightColor, true)                      // Top
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(lineWidth), float32(squareSize), gm.highlightColor, true)                      // Left
		vector.DrawFilledRect(screen, float32(x+squareSize-lineWidth), float32(y), float32(lineWidth), float32(squareSize), gm.highlightColor, true) // Right
		vector.DrawFilledRect(screen, float32(x), float32(y+squareSize-lineWidth), float32(squareSize), float32(lineWidth), gm.highlightColor, true) // Bottom
	}
}

// drawAvailableMoves draws the available moves to the screen.
//
// this renders the available moves to the screen as highlighted squares.
func (gm *GameManager) drawAvailableMoves(screen *ebiten.Image, squareSize int) {
	for _, move := range gm.state.availableMoves {
		toSquare := move.S2()
		x, y := int(toSquare%8)*squareSize, int(toSquare/8)*squareSize
		moveHighlightColor := color.RGBA{0, 255, 0, 128} // Semi-transparent green
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(squareSize), float32(squareSize), moveHighlightColor, true)
	}
}

// drawPieces draws the chess pieces to the screen.
//
// this renders the chess pieces to the screen based on the current game state.
func (gm *GameManager) drawPieces(screen *ebiten.Image, squareSize int) {
	board := gm.state.game.Position().Board() // get the current board state
	for i := 0; i < 8; i++ {                  // Iterate over ranks (rows)
		for j := 0; j < 8; j++ { // Iterate over files (columns)
			square := chess.Square(i*8 + j)
			piece := board.Piece(square)

			if piece != chess.NoPiece {
				// Calculate x and y, keeping white on top
				x := j * squareSize
				y := i * squareSize

				// Center the piece image within the square
				pieceImg := gm.pieceImages[piece]
				pieceWidth, pieceHeight := pieceImg.Bounds().Dx(), pieceImg.Bounds().Dy()
				offsetX := (squareSize - pieceWidth) / 2
				offsetY := (squareSize - pieceHeight) / 2

				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(float64(x+offsetX), float64(y+offsetY))
				screen.DrawImage(pieceImg, op)
			}
		}
	}
}

// drawRestartButton draws the restart button to the screen.
//
// this renders a restart button to the screen when the game is over.
func (gm *GameManager) drawRestartButton(screen *ebiten.Image) {
	// if the game is not over, return
	outcome, gameover := gm.gameover()
	if !gameover {
		return
	}

	buttonWidth := 200
	buttonHeight := 100

	screenWidth, screenHeight := screen.Bounds().Dx(), screen.Bounds().Dy()
	centerX, centerY := float64(screenWidth/2), float64(screenHeight/2)
	buttonX := int(centerX) - buttonWidth/2
	buttonY := int(centerY)
	buttonCenterX := buttonX + buttonWidth/2
	buttonThirdY := buttonY + buttonHeight/3
	// buttonCenterY := buttonY + buttonHeight/2

	// Draw the button background (a filled rectangle)
	buttonColor := color.RGBA{0, 128, 0, 255} // Green
	vector.DrawFilledRect(screen, float32(buttonX), float32(buttonY), float32(buttonWidth), float32(buttonHeight), buttonColor, true)

	// Draw the button text
	op := &text.DrawOptions{}
	op.GeoM.Scale(1.5, 1.5)
	// position the text in the center of the button
	op.GeoM.Translate(float64(buttonCenterX), float64(buttonThirdY)-10)
	op.PrimaryAlign = text.AlignCenter
	text.Draw(
		screen,
		"Restart",
		text.NewGoXFace(bitmapfont.Face),
		op,
	)

	// print the game over message
	var textValue string
	switch outcome {
	case chess.WhiteWon:
		textValue = "White wins!"
	case chess.BlackWon:
		textValue = "Black wins!"
	case chess.Draw:
		textValue = "Draw!"
	default:
		textValue = "No Outcome"
	}
	op.GeoM.Translate(0, 20)
	text.Draw(
		screen,
		fmt.Sprintf("Game Over: %s", textValue),
		text.NewGoXFace(bitmapfont.Face),
		op,
	)
}

// // drawBoardMap draws the board map to the screen.
// //
// // this renders an overlay to show the board coordinates.
// func (gm *GameManager) drawBoardMap(screen *ebiten.Image, squareSize int) {
// 	// Render the square map on the side
// 	squareMapX := gm.boardSize + 20 // Adjust positioning as needed
// 	squareMapY := 20
// 	face := text.Font{} // Replace with your actual font face
// 	textOptions := &text.DrawOptions{}
// 	textOptions.ColorM.Scale(0, 0, 0, 1) // Set text color to black

// 	for i := 7; i >= 0; i-- { // Iterate rows in reverse for chess notation
// 		for j := 0; j < 8; j++ {
// 			square := chess.Square(i*8 + j)
// 			file := chess.Files[square/8]
// 			rank := 8 - (square % 8)
// 			squareName := fmt.Sprintf("%c%d", file, rank)

// 			text.Draw(screen, squareName, face, textOptions)
// 			textOptions.GeoM.Translate(float64(squareSize), 0)
// 		}
// 		textOptions.GeoM.Translate(-8*float64(squareSize), float64(squareSize)) // Move to the next row
// 	}
// }
