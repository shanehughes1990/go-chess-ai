package gamemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/sirupsen/logrus"
)

// updateKeyRJustPressed handles the key press event for the R key.
func (gm *GameManager) updateKeyRJustPressed() {
	if inpututil.IsKeyJustPressed(ebiten.KeyR) {
		gm.debugMenu = !gm.debugMenu
		if gm.debugMenu {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(logrus.InfoLevel)
		}
	}
}

// updateRestartButtonJustPressed handles the mouse click event for the restart button.
func (gm *GameManager) updateRestartButtonJustPressed() {
	_, gameover := gm.gameover()
	if gameover && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		buttonWidth := 200
		buttonHeight := 100
		screenWidth, screenHeight := gm.windowWidth, gm.windowHeight
		centerX, centerY := float64(screenWidth/2), float64(screenHeight/2)
		buttonX := int(centerX) - buttonWidth/2
		buttonY := int(centerY)

		// Check if the click is within the button's bounds
		if x >= buttonX && x <= buttonX+buttonWidth && y >= buttonY && y <= buttonY+buttonHeight {
			gm.restartGameState()
		}
	}
}
