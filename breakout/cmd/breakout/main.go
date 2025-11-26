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
	)

	world.AddObject(objects.NewPlayArea(world, assets.DefaultBackground))
	world.AddObject(objects.NewBricks(assets.Levels, 16))
	world.AddObject(objects.NewBall(world, assets.Ball))
	world.AddObject(objects.NewPaddle(world, assets.Paddle))

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
