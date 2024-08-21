# Chess AI Game Engine

A Go-based chess game engine designed for flexibility and experimentation with various chess AI algorithms.

## Key Features

* **Modular Architecture:** Easily integrate and test different chess AI algorithms.
* **Interactive Gameplay:** Play chess against human opponents or AI bots.
* **Ebiten Integration:** Utilizes [hajimehoshi/ebiten](https://github.com/hajimehoshi/ebiten) for rendering and handling user input.
* **Chess Package Integration:** Leverages [notnil/chess](https://github.com/notnil/chess) package for core chess logic and move validation.
* **Self-Learning AI Support:** Built-in support for developing and training self-learning AI opponents using neural networks and reinforcement learning. _(Comming soon)_

# Usage

The following snippet shows you how to use the chess game engine to play a game of chess between a human player and the random AI player.

```go
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
```

# Implimenting a Custom AI Player

To impliment a custom AI player, you need to impliment the chessgame.Player interface. the following bot impliments a simple AI player that makes random moves.

The Player interface is as follows
```go
// Player represents a player in the game.
type Player interface {
	// Name returns the name of the player.
	Name() string
	// MakeMove is the method that decides the move for the player.
	//
	// Returning a move will cause the game engine to finalize the move.
	//
	// Returning nil will cause the game engine to wait for the player to make a move.
	MakeMove(game *GameState) (*chess.Move, error)
}
```

```go
package randomai

import (
	"math/rand"

	"github.com/notnil/chess"
	"github.com/shanehughes1990/chess-ai/chessgame/v1"
)

// randomAI is a player that makes random moves.
type randomAI struct{ name string }

// NewRandomAI creates a new randomAI Player.
func NewRandomAI(name string) chessgame.Player {
	return &randomAI{name: name}
}

// Name returns the name of the player.
func (p *randomAI) Name() string {
	return p.name
}

// MakeMove gets all the valid moves from the chess engine and picks one at random
//
// Returning a *chess.Move will result in finalizing the move in the chess engine.
//
// Otherwise you can return nil to skip the moveplayer, and will be empty)
func (p *randomAI) MakeMove(game *chessgame.GameState) (*chess.Move, error) {
	// Get all valid moves for the current player
	validMoves := game.Game().ValidMoves()
	if len(validMoves) == 0 {
		return nil, nil
	}

	// Choose a random move from the valid moves
	randomIndex := rand.Intn(len(validMoves))
	chosenMove := validMoves[randomIndex]

	return chosenMove, nil
}
```