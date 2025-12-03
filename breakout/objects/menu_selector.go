package objects

import (
	"time"

	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
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
	drawSprite bool
}

func NewMenuSelector(selector *ebiten.Image, song assets.SoundEffect, nextState func() game.State) *MenuSelector {
	song.ChangeVolume(1.0)
	return &MenuSelector{
		position:   vector{X: 510, Y: 415},
		sprite:     selector,
		nextY:      []float64{0.0, 58.0, 116},
		song:       song,
		nextState:  nextState,
		startSong:  game.NewTimer(500 * time.Millisecond),
		drawSprite: true,
	}
}

func (ms *MenuSelector) Update(world *game.World, state game.State) {
	if ms.transition != nil {
		ms.song.ChangeVolume(-0.0055)
		ms.drawSprite = ms.transition.Update()%30 < 15

		if ms.transition.IsReady() {
			ms.song.Stop()
			state.Next(ms.nextState())

			return
		}
	}

	if !ms.song.IsPlaying() {
		ms.startSong.Update()
		if ms.startSong.IsReady() {
			ms.song.Play()
		}
	}

	if world.JustPressedDown() {
		ms.currY = min(len(ms.nextY)-1, ms.currY+1)
	} else if world.JustPressedUp() {
		ms.currY = max(0, ms.currY-1)
	} else if world.JustPressedAction() {
		ms.transition = game.NewTimer(3000 * time.Millisecond)
	}
}

func (ms *MenuSelector) Reset() {
}

func (ms *MenuSelector) Draw(display *ebiten.Image) {
	if !ms.drawSprite {
		return
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ms.position.X, ms.position.Y+ms.nextY[ms.currY])
	display.DrawImage(ms.sprite, op)
}
