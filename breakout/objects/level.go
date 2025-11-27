package objects

import (
	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Bricks struct {
	levels    []assets.Level
	currLevel int
	playArea  *PlayArea
}

func NewBricks(levels []assets.Level, playArea *PlayArea) *Bricks {
	return &Bricks{
		levels:    levels,
		currLevel: 0,
		playArea:  playArea,
	}
}

func (b *Bricks) Update(world *game.World) {
}

func (b *Bricks) Draw(display *ebiten.Image) {
	level := b.levels[b.currLevel]
	for _, brick := range level.Bricks() {
		if brick.Sprite() == nil {
			continue
		}
		b.drawBrick(display, brick)
	}

}

func (b *Bricks) drawBrick(display *ebiten.Image, brick assets.Brick) {
	op := &ebiten.DrawImageOptions{}
	x, y := brick.Position()
	playArea := b.playArea.Rect()
	op.GeoM.Translate(float64(x)+playArea.X, float64(y)+playArea.Y)
	display.DrawImage(brick.Sprite(), op)
}

func (b *Bricks) Rect() game.Rect {
	return game.Rect{}
}

func (b *Bricks) Intersects(bounds game.Bounds) bool {
	return b.Rect().Intersects(bounds.Rect())
}
