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
	)

	playArea := objects.NewPlayArea(16, assets.DefaultBackground)
	world.AddObject(playArea)

	bricks := objects.NewBricks(assets.Levels, playArea)
	world.AddObject(bricks)

	paddle := objects.NewPaddle(assets.Paddle, playArea)
	world.AddObject(paddle)

	world.AddObject(objects.NewBall(assets.Ball, playArea, paddle, bricks))

	if err := ebiten.RunGame(world); err != nil {
		panic(err)
	}
}
