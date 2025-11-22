package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type HiddenItem int

const (
	NOHTING = iota
)

type Brick struct {
	position  vector
	layers    []*ebiten.Image
	currLayer int
	hidden    HiddenItem
}

func NewBrick(x, y int, layers []*ebiten.Image, hidden HiddenItem) game.Object {
	return &Brick{
		position:  vector{X: float64(x), Y: float64(y)},
		layers:    layers,
		hidden:    hidden,
		currLayer: len(layers) - 1,
	}
}

func (b *Brick) Update(world *game.World) {
}

func (b *Brick) Draw(display *ebiten.Image) {
}

func (b *Brick) Rect() game.Rect {
	bounds := b.layers[b.currLayer].Bounds()

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
