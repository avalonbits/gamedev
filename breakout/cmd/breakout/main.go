package main

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	world := game.NewWorld("Breakoout", 1024, 788)

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
