package object

import (
	"math"
	"math/rand/v2"

	"github.com/avalonbits/gamedev/spaceshoot/assets"
	"github.com/avalonbits/gamedev/spaceshoot/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Meteor struct {
	position      vector
	movement      vector
	rotation      float64
	rotationSpeed float64
	sprite        *ebiten.Image
}

func NewMeteor(screenW, screenH int) game.Object {
	sprite := assets.Meteors[rand.IntN(len(assets.Meteors))]

	// Figure out the target position — the screen center, in this case
	target := vector{
		X: float64(screenW / 2),
		Y: float64(screenH / 2),
	}

	// The distance from the center the meteor should spawn at — half the width
	r := float64(screenW / 2)

	// Pick a random angle — 2π is 360° — so this returns 0° to 360°
	angle := rand.Float64() * 2 * math.Pi

	// Figure out the spawn position by moving r pixels from the target at the chosen angle
	pos := vector{
		X: target.X + math.Cos(angle)*r,
		Y: target.Y + math.Sin(angle)*r,
	}

	// Randomized velocity
	velocity := 0.25 + rand.Float64()*1.5

	// Direction is the target minus the current position
	direction := vector{
		X: target.X - pos.X,
		Y: target.Y - pos.Y,
	}

	// Normalize the vector — get just the direction without the length
	normalizedDirection := direction.Normalize()

	// Multiply the direction by velocity
	movement := vector{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	return &Meteor{
		position:      pos,
		movement:      movement,
		rotationSpeed: 0.02 + rand.Float64()*0.04,
		sprite:        sprite,
	}
}

func (m *Meteor) Update(_ *game.World) {
	m.position.X += m.movement.X
	m.position.Y += m.movement.Y
	m.rotation += m.rotationSpeed
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	bounds := m.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	// Rotate across the center
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-halfW, -halfH)
	op.GeoM.Rotate(m.rotation)
	op.GeoM.Translate(halfW, halfH)

	op.GeoM.Translate(m.position.X, m.position.Y)

	screen.DrawImage(m.sprite, op)
}

func (m *Meteor) Rect() game.Rect {
	bounds := m.sprite.Bounds()

	return game.NewRect(
		m.position.X,
		m.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
