package main

import (
	"runtime"
	"time"

	"github.com/avalonbits/gamedev/spaceshoot/assets"
	"github.com/avalonbits/gamedev/spaceshoot/game"
	"github.com/avalonbits/gamedev/spaceshoot/object"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 600
)

func main() {
	runtime.GOMAXPROCS(1)
	g := game.NewWorld(
		ScreenWidth,
		ScreenHeight,
		func(screenW, screenH int) game.Object {
			return object.NewPlayer(
				screenW,
				screenH,
				assets.Player,
				500*time.Millisecond,
			)
		},
		object.NewMeteor,
		1000*time.Millisecond,
	)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
