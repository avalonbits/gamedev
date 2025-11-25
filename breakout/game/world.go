package game

import (
	"github.com/hajimehoshi/ebiten/v2"
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
	screenW int
	screenH int
	playW   int
	margin  int
	objects []Object
}

func NewWorld(
	title string,
	screenW int,
	screenH int,
	playW int,
	margin int,
	objectFn ...ObjectFactory,
) *World {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(screenW, screenH)

	w := &World{
		screenW: screenW,
		screenH: screenH,
		playW:   playW,
		margin:  margin,
	}
	for _, fn := range objectFn {
		w.objects = append(w.objects, fn(w))
	}

	return w
}

func (w *World) Width() int {
	return w.screenW
}

func (w *World) PlayWidth() int {
	return w.playW
}

func (w *World) Height() int {
	return w.screenH
}

func (w *World) Margin() int {
	return w.margin
}

func (w *World) Update() error {
	for _, obj := range w.objects {
		obj.Update(w)
	}
	return nil

}

func (w *World) Draw(screen *ebiten.Image) {
	for _, obj := range w.objects {
		obj.Draw(screen)
	}
}

func (w *World) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return w.screenW, w.screenH
}
