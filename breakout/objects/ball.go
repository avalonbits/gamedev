package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	sprite    *ebiten.Image
	position  vector
	direction vector
	maxY      float64
	speed     float64
	playArea  *PlayArea
}

func NewBall(sprite *ebiten.Image, playArea *PlayArea) *Ball {
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := vector{
		X: (playArea.Rect().Width/2 + playArea.Rect().X) - halfW,
		Y: (playArea.Rect().MaxY() - 32) - 256,
	}

	return &Ball{
		position: position,
		maxY:     position.Y,
		sprite:   sprite,
	}
}

func (b *Ball) Update(world *game.World) {
}

func (b *Ball) moveWithPaddle(world *game.World) {

}

func (b *Ball) moveFreely(world *game.World) {
}

func (b *Ball) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.position.X, b.position.Y)
	display.DrawImage(b.sprite, op)
}

func (b *Ball) Rect() game.Rect {
	bounds := b.sprite.Bounds()

	return game.NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (b *Ball) Intersects(bounds game.Bounds) bool {
	return b.Rect().Intersects(bounds.Rect())
}
