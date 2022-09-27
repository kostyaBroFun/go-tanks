// Package gameloop provides module for loop of game tanks.
// User connects to GameLoop, then sends commands and
// receives state of game several times per second.
// #TODO: provide only changes of state.
package gameloop

import (
	"github.com/kostyaBroFun/go-tanks/internal/service/gameloop/game"
)

type GameLoop struct {
}

func (gl *GameLoop) Connect(player game.Player) <-chan game.State {
	panic("not implemented")
}

func (gl *GameLoop) CurrentState() game.State {
	panic("not implemented")
}

func (gl *GameLoop) Close() error {
	panic("not implemented")
}
