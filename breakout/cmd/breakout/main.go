package main

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/avalonbits/gamedev/breakout/states"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	world := game.NewWorld(
		"Breakoout",
		1280,
		720,
	)

	world.SetState(states.NewMenu())

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
