package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Paddle struct {
	position  vector
	sprite    *ebiten.Image
	speed     float64
	direction float64
}

func NewPaddle(world *game.World, sprite *ebiten.Image) *Paddle {
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := vector{
		X: float64(world.PlayWidth()/2+world.Margin()) - halfW,
		Y: float64(world.Height()-world.Margin()-32) - halfH,
	}

	return &Paddle{
		position:  position,
		sprite:    sprite,
		direction: 1.0,
	}
}

func (b *Paddle) Update(world *game.World) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		b.speed = min(8, b.speed+1.5)
		b.direction = -1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		b.speed = min(8, b.speed+1.5)
		b.direction = 1.0
	} else {
		b.speed = max(0, b.speed-2)
	}

	paddleW := b.sprite.Bounds().Max.X - b.sprite.Bounds().Min.X
	maxX := float64(world.PlayWidth() + world.Margin() - paddleW)
	minX := float64(world.Margin())
	b.position.X = max(minX, min(maxX, b.position.X+b.speed*b.direction))
}

func (b *Paddle) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.position.X, b.position.Y)
	display.DrawImage(b.sprite, op)
}

func (b *Paddle) Rect() game.Rect {
	bounds := b.sprite.Bounds()

	return game.NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (b *Paddle) Intersects(bounds game.Bounds) bool {
	return b.Rect().Intersects(bounds.Rect())
}
