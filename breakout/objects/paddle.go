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
	playArea  *PlayArea
}

func NewPaddle(sprite *ebiten.Image, playArea *PlayArea) *Paddle {
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := vector{
		X: (playArea.Rect().Width/2 + playArea.Rect().X) - halfW,
		Y: (playArea.Rect().MaxY() - 32) - halfH,
	}

	return &Paddle{
		position:  position,
		sprite:    sprite,
		direction: 1.0,
		playArea:  playArea,
	}
}

func (b *Paddle) Update(world *game.World) {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		b.speed = 8
		b.direction = -1.0
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		b.speed = 8
		b.direction = 1.0
	} else {
		b.speed = 0
	}

	rect := b.playArea.Rect()
	paddleW := float64(b.sprite.Bounds().Dx())
	maxX := rect.MaxX() - paddleW
	minX := rect.X
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
