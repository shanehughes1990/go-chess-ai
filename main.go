package main

import (
	"github.com/shanehughes1990/chess-ai/chessgame/v1"
	"github.com/shanehughes1990/chess-ai/chessgame/v1/bots/randomai"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	manager := chessgame.NewGameManager(
		chessgame.WithWhitePlayer(chessgame.NewHumanPlayer("Player 1")),
		chessgame.WithBlackPlayer(randomai.NewRandomAI("Player 2")),
	)

	if err := manager.Start(); err != nil {
		logrus.WithError(err).Panic("failed to start chess game")
	}
}
