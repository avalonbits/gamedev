package main

import (
	"math"

	"github.com/avalonbits/gamedev/spaceshoot/assets"
	"github.com/hajimehoshi/ebiten/v2"
)

type vector struct {
	X float64
	Y float64
}

type Game struct {
	playerSprite   *ebiten.Image
	playerPosition vector
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	width, height := g.playerSprite.Bounds().Dx(), g.playerSprite.Bounds().Dy()
	halfW, halfH := float64(width/2), float64(height/2)

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(45.0 * math.Pi / 180.0)
	op.GeoM.Translate(halfW, halfH)

	screen.DrawImage(g.playerSprite, op)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	g := &Game{
		playerSprite:   assets.Load("player.png"),
		playerPosition: vector{X: 100, Y: 100},
	}

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
