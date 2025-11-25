package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayArea struct {
	sprite *ebiten.Image
	margin float64
}

func NewPlayArea(world *game.World, sprite *ebiten.Image) *PlayArea {
	return &PlayArea{
		sprite: sprite,
		margin: float64(world.Margin()),
	}
}

func (pa *PlayArea) Update(world *game.World) {
}

func (pa *PlayArea) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pa.margin, pa.margin)
	display.DrawImage(pa.sprite, op)
}

func (pa *PlayArea) Rect() game.Rect {
	bounds := pa.sprite.Bounds()

	return game.NewRect(
		pa.margin,
		pa.margin,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (pa *PlayArea) Intersects(bounds game.Bounds) bool {
	return pa.Rect().Intersects(bounds.Rect())
}
