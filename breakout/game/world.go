package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type World struct {
	screenW int
	screenH int
}

func NewWorld(title string, screenW, screenH int) *World {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(screenW, screenH)
	return &World{
		screenW: screenW,
		screenH: screenH,
	}
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
}

func (w *World) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return w.screenW, w.screenH
}
