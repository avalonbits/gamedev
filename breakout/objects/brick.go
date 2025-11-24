package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Brick struct {
	position vector
	sprite   *ebiten.Image
	hitCount int
}

func NewBrick(x, y, hitCount int, sprite *ebiten.Image) *Brick {
	return &Brick{
		position: vector{X: float64(x), Y: float64(y)},
		sprite:   sprite,
		hitCount: hitCount,
	}
}

func (b *Brick) Update(world *game.World) {
}

func (b *Brick) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.position.X, b.position.Y)
	display.DrawImage(b.sprite, op)
}

func (b *Brick) Rect() game.Rect {
	bounds := b.sprite.Bounds()

	return game.NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (b *Brick) Intersects(bounds game.Bounds) bool {
	return b.Rect().Intersects(bounds.Rect())
}
