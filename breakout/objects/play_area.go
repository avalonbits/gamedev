package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type PlayArea struct {
	sprite *ebiten.Image
	margin float64
}

func NewPlayArea(margin float64, sprite *ebiten.Image) *PlayArea {
	return &PlayArea{
		sprite: sprite,
		margin: margin,
	}
}

func (pa *PlayArea) Update(world *game.World, stateFn func(game.State)) {
}

func (pa *PlayArea) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(pa.margin, pa.margin)
	display.DrawImage(pa.sprite, op)
}

func (pa *PlayArea) Rect() rect {
	bounds := pa.sprite.Bounds()

	return NewRect(
		pa.margin,
		pa.margin,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
