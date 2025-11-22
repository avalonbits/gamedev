package main

import (
	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/avalonbits/gamedev/breakout/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	world := game.NewWorld("Breakoout", 1280, 720, func(_ *game.World) game.Object {
		return objects.NewBrick(0, 0, 1, assets.Bricks[0])
	})

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
