package objects

import (
	"time"

	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type MenuSelector struct {
	position   vector
	sprite     *ebiten.Image
	nextY      []float64
	currY      int
	song       assets.SoundEffect
	nextState  func() game.State
	startSong  *game.Timer
	transition *game.Timer
}

func NewMenuSelector(selector *ebiten.Image, song assets.SoundEffect, nextState func() game.State) *MenuSelector {
	return &MenuSelector{
		position:  vector{X: 510, Y: 415},
		sprite:    selector,
		nextY:     []float64{0.0, 58.0, 116},
		song:      song,
		nextState: nextState,
		startSong: game.NewTimer(500 * time.Millisecond),
	}
}

func (ms *MenuSelector) Update(world *game.World, stateFn func(game.State)) {
	if ms.transition != nil {
		ms.transition.Update()
		ms.song.ChangeVolume(-0.02)

		if ms.transition.IsReady() {
			ms.song.Stop()
			stateFn(ms.nextState())

			return
		}
	}

	if !ms.song.IsPlaying() {
		ms.startSong.Update()
		if ms.startSong.IsReady() {
			ms.song.Play()
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		ms.currY = (ms.currY + 1) % len(ms.nextY)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		ms.currY--
		if ms.currY < 0 {
			ms.currY = len(ms.nextY) - 1
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		ms.transition = game.NewTimer(1500 * time.Millisecond)
	}
}

func (ms *MenuSelector) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ms.position.X, ms.position.Y+ms.nextY[ms.currY])
	display.DrawImage(ms.sprite, op)
}
