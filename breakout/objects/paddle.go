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

	p := &Paddle{
		position: position,
		sprite:   sprite,
		playArea: playArea,
	}
	p.Reset()

	return p
}

func (p *Paddle) Reset() {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	position := vector{
		X: (p.playArea.Rect().Width/2 + p.playArea.Rect().X) - halfW,
		Y: (p.playArea.Rect().MaxY() - 32) - halfH,
	}

	p.position = position
	p.direction = 0.0

}

func (b *Paddle) Direction() float64 {
	return b.direction
}

func (b *Paddle) Update(world *game.World, state game.State) {
	if dir := world.HorizontalAxis(); dir != 0.0 {
		b.speed = 9
		b.direction = dir
	} else if world.PressLeft() {
		b.speed = 9
		b.direction = -1.0
	} else if world.PressRight() {
		b.speed = 9
		b.direction = 1.0
	} else {
		b.speed = 0.0
		b.direction = 0.0
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

func (b *Paddle) Rect() rect {
	bounds := b.sprite.Bounds()

	return NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
