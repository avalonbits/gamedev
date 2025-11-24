package main

import (
	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/avalonbits/gamedev/breakout/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	world := game.NewWorld(
		"Breakoout",
		1280,
		720,
		13*64,
		16,
		func(world *game.World) game.Object {
			return objects.NewPaddle(world, assets.Paddle)
		},
		func(_ *game.World) game.Object {
			return objects.NewBricks(assets.Levels, 16)
		},
	)

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
