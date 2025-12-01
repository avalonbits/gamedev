package objects

import (
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuSelector struct {
	position  vector
	sprite    *ebiten.Image
	nextY     []float64
	currY     int
	nextState func() game.State
}

func NewMenuSelector(selector *ebiten.Image, nextState func() game.State) *MenuSelector {
	return &MenuSelector{
		position:  vector{X: 510, Y: 415},
		sprite:    selector,
		nextY:     []float64{0.0, 58.0, 116},
		nextState: nextState,
	}
}

func (ms *MenuSelector) Update(world *game.World, stateFn func(game.State)) {
	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		ms.currY = (ms.currY + 1) % len(ms.nextY)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		ms.currY--
		if ms.currY < 0 {
			ms.currY = len(ms.nextY) - 1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		stateFn(ms.nextState())
	}
}

func (ms *MenuSelector) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ms.position.X, ms.position.Y+ms.nextY[ms.currY])
	display.DrawImage(ms.sprite, op)
}
