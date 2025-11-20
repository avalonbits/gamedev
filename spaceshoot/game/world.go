package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update()
	Draw(*ebiten.Image)
}

type ObjectFactory func(screenW, screenH int) Object

type World struct {
	screenW          int
	screenH          int
	player           Object
	meteors          []Object
	meteorFn         ObjectFactory
	meteorSpawnTimer *Timer
}

func NewWorld(
	screenW int,
	screenH int,
	player Object,
	meteorFn ObjectFactory,
	meteorSpawn time.Duration,
) *World {
	return &World{
		screenW:          screenW,
		screenH:          screenH,
		player:           player,
		meteorFn:         meteorFn,
		meteorSpawnTimer: NewTimer(meteorSpawn),
	}

}
func (w *World) Update() error {
	w.player.Update()

	w.meteorSpawnTimer.Update()
	if w.meteorSpawnTimer.IsReady() {
		w.meteorSpawnTimer.Reset()

		w.meteors = append(w.meteors, w.meteorFn(w.screenW, w.screenH))
	}

	for _, m := range w.meteors {
		m.Update()
	}
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	w.player.Draw(screen)

	for _, m := range w.meteors {
		m.Draw(screen)
	}
}

func (w *World) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return int(w.screenW), int(w.screenH)
}
