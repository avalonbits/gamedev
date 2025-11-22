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
	bricks  []Object
}

func NewWorld(
	title string,
	screenW int,
	screenH int,
	brickFn ObjectFactory,
) *World {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(screenW, screenH)

	w := &World{
		screenW: screenW,
		screenH: screenH,
	}
	w.bricks = w.createBricks(brickFn)

	return w
}

func (w *World) createBricks(brickFn ObjectFactory) []Object {
	bricks := []Object{}
	bricks = append(bricks, brickFn(w))
	return bricks
}

func (w *World) Width() int {
	return w.screenW
}

func (w *World) Height() int {
	return w.screenH
}

func (w *World) Update() error {
	return nil

}

func (w *World) Draw(screen *ebiten.Image) {
	for _, brick := range w.bricks {
		brick.Draw(screen)
	}
}

func (w *World) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return w.screenW, w.screenH
}
