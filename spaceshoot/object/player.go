package object

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	position vector
	sprite   *ebiten.Image
	rotation float64
}

func NewPlayer(sprite *ebiten.Image, screenW, screenH int) *Player {
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := vector{
		X: float64(screenW/2) - halfW,
		Y: float64(screenH/2) - halfH,
	}

	return &Player{
		position: pos,
		sprite:   sprite,
	}
}

func (p *Player) Update() {
	speed := math.Pi / float64(ebiten.TPS()) // 180 degrees per second

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	bounds := p.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// Rotate across the center
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
