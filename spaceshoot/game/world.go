package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/avalonbits/gamedev/spaceshoot/assets"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

type Bounds interface {
	Rect() Rect
	Intersects(Bounds) bool
}

type Object interface {
	Bounds

	Update(world *World)
	Draw(*ebiten.Image)
}

type ObjectFactory func(world *World) Object

type World struct {
	screenW          int
	screenH          int
	score            int
	player           Object
	meteors          []Object
	bullets          []Object
	meteorFn         ObjectFactory
	playerFn         ObjectFactory
	meteorSpawnTimer *Timer
}

func NewWorld(
	screenW int,
	screenH int,
	playerFn ObjectFactory,
	meteorFn ObjectFactory,
	meteorSpawn time.Duration,
) *World {
	ebiten.SetWindowTitle("Space Shooter")
	ebiten.SetWindowSize(screenW, screenH)

	world := World{
		screenW:          screenW,
		screenH:          screenH,
		playerFn:         playerFn,
		meteorFn:         meteorFn,
		meteorSpawnTimer: NewTimer(meteorSpawn),
	}

	world.player = playerFn(&world)
	return &world
}

func (w *World) Width() int {
	return w.screenW
}

func (w *World) Height() int {
	return w.screenH
}

func (w *World) AddBullet(bullet Object) {
	w.bullets = append(w.bullets, bullet)
}

func (w *World) Update() error {
	w.player.Update(w)

	w.meteorSpawnTimer.Update()
	if w.meteorSpawnTimer.IsReady() {
		w.meteorSpawnTimer.Reset()
		w.meteors = append(w.meteors, w.meteorFn(w))
	}

	for _, m := range w.meteors {
		m.Update(w)
	}

	for _, b := range w.bullets {
		b.Update(w)
	}

	// Check for meteor/bullet collisions
	for i, m := range w.meteors {
		for j, b := range w.bullets {
			if m.Intersects(b) {
				w.meteors = append(w.meteors[:i], w.meteors[i+1:]...)
				w.bullets = append(w.bullets[:j], w.bullets[j+1:]...)
				w.score++
			}
		}
		if m.Intersects(w.player) {
			w.Reset()
			break
		}

	}
	return nil
}

func (w *World) Draw(screen *ebiten.Image) {
	w.player.Draw(screen)

	for _, m := range w.meteors {
		m.Draw(screen)
	}

	for _, b := range w.bullets {
		b.Draw(screen)
	}

	text.Draw(screen, fmt.Sprintf("%06d", w.score), assets.ScoreFont, w.screenW/2-100, 50, color.White)
}

func (w *World) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return w.screenW, w.screenH
}

func (w *World) Reset() {
	w.player = w.playerFn(w)
	w.meteors = w.meteors[:0]
	w.bullets = w.bullets[:0]
	w.score = 0
	w.meteorSpawnTimer.Reset()
}
