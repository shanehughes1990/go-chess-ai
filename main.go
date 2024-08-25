package main

import (
	"fmt"
	"time"

	"github.com/shanehughes1990/chess-ai/chessgame/v1"
	"github.com/shanehughes1990/chess-ai/chessgame/v1/bots/randomai"
	filetracer "github.com/shanehughes1990/chess-ai/chessgame/v1/tracer/file"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	manager, err := chessgame.NewGameManager(
		chessgame.WithWhitePlayer(chessgame.NewHumanPlayer("Player 1")),
		chessgame.WithBlackPlayer(randomai.NewRandomAI("Player 2")),
		chessgame.WithTracer(filetracer.NewFileTracer(fmt.Sprintf("game-%d.txt", time.Now().Unix()))),
	)
	if err != nil {
		logrus.WithError(err).Panic("failed to create chess game")
	}

	// wait for the game to finish
	if err := manager.Start(); err != nil {
		logrus.WithError(err).Panic("failed to start chess game")
	}
}
