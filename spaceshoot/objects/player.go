package objects

import (
	"math"
	"time"

	"github.com/avalonbits/gamedev/spaceshoot/game"
	"github.com/hajimehoshi/ebiten/v2"
)

const ()

type Player struct {
	position      vector
	sprite        *ebiten.Image
	rotation      float64
	shootCooldown *game.Timer
}

func NewPlayer(world *game.World, sprite *ebiten.Image, shootCooldown time.Duration) *Player {
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2
	halfH := float64(bounds.Dy()) / 2

	pos := vector{
		X: float64(world.Width()/2) - halfW,
		Y: float64(world.Height()/2) - halfH,
	}

	return &Player{
		position:      pos,
		sprite:        sprite,
		shootCooldown: game.NewTimer(shootCooldown),
	}
}

func (p *Player) Update(world *game.World) {
	speed := math.Pi / float64(ebiten.TPS()) // 180 degrees per second

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}

	p.shootCooldown.Update()
	if p.shootCooldown.IsReady() && ebiten.IsKeyPressed(ebiten.KeySpace) {
		const bulletSpawnOffset = 50.0

		p.shootCooldown.Reset()
		p.shootCooldown.Reset()

		bounds := p.sprite.Bounds()
		halfW := float64(bounds.Dx()) / 2
		halfH := float64(bounds.Dy()) / 2

		spawnPos := vector{
			p.position.X + halfW + math.Sin(p.rotation)*bulletSpawnOffset,
			p.position.Y + halfH + math.Cos(p.rotation)*-bulletSpawnOffset,
		}

		bullet := NewBullet(spawnPos, p.rotation)
		world.AddBullet(bullet)
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

func (p *Player) Rect() game.Rect {
	bounds := p.sprite.Bounds()

	return game.NewRect(
		p.position.X,
		p.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (p *Player) Intersects(bounds game.Bounds) bool {
	return p.Rect().Intersects(bounds.Rect())
}
