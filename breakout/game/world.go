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
	screenW        int
	screenH        int
	playW          int
	margin         int
	objects        []Object
	availableSlots []int
	next           int
}

func NewWorld(
	title string,
	screenW int,
	screenH int,
	playW int,
	margin int,
) *World {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(screenW, screenH)

	return &World{
		screenW: screenW,
		screenH: screenH,
		playW:   playW,
		margin:  margin,
	}
}

func (w *World) AppendObjects(obj Object) int {
	w.objects = append(w.objects, obj)
	return len(w.objects) - 1
}

func (w *World) AddObject(obj Object) int {
	if w.next == len(w.availableSlots) {
		w.next = 0
		w.availableSlots = w.availableSlots[:0]
		return w.AppendObjects(obj)
	}

	idx := w.availableSlots[w.next]
	w.objects[idx] = obj
	w.next++

	return idx
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
