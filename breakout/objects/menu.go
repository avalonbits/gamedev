package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Menu struct {
	position vector
	sprite   *ebiten.Image
}

func NewMenu(background *ebiten.Image) *Menu {
	return &Menu{
		sprite: background,
	}
}

func (m *Menu) Update(world *game.World, _ game.State) {
}

func (m *Menu) Reset() {
}

func (m *Menu) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	display.DrawImage(m.sprite, op)
}
