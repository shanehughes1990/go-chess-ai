package gamemanager

import (
	"github.com/hajimehoshi/ebiten/v2"
)

// GameManagerOption is a function that modifies a GameManager.
type GameManagerOption func(*GameManager)

// WithWindowTitle sets the window title of the game.
func WithWindowTitle(title string) GameManagerOption {
	return func(gm *GameManager) {
		ebiten.SetWindowTitle(title)
	}
}

// WithWindowSize sets the window size of the game.
func WithWindowSize(windowWidth, windowHeight int) GameManagerOption {
	return func(gm *GameManager) {
		gm.windowWidth = windowWidth
		gm.windowHeight = windowHeight
	}
}

// WithTPS sets the target ticks per second of the game.
func WithTPS(tps int) GameManagerOption {
	return func(gm *GameManager) {
		gm.tps = tps
	}
}

// WithVSyncEnabled sets whether VSync is enabled.
func WithVSyncEnabled(enabled bool) GameManagerOption {
	return func(gm *GameManager) {
		gm.vsyncEnabled = enabled
	}
}

// WithRunnableOnUnfocused sets whether the game should continue running when unfocused.
func WithRunnableOnUnfocused(runnable bool) GameManagerOption {
	return func(gm *GameManager) {
		gm.runnableOnUnfocused = runnable
	}
}

// WithPlayer1 sets the human player.
func WithPlayer1(player Player) GameManagerOption {
	return func(gm *GameManager) {
		gm.state.player1 = player
	}
}

// WithPlayer2 sets the human player.
func WithPlayer2(player Player) GameManagerOption {
	return func(gm *GameManager) {
		gm.state.player2 = player
	}
}
