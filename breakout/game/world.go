package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type State interface {
	Update(world *World) State
	Draw(*ebiten.Image)
}

type StateFactory func(world *World) State

type World struct {
	screenW        int
	screenH        int
	state          State
	availableSlots []int
	next           int
}

func NewWorld(
	title string,
	screenW int,
	screenH int,
) *World {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetVsyncEnabled(true)

	return &World{
		screenW: screenW,
		screenH: screenH,
	}
}

func (w *World) SetState(state State) {
	w.state = state
}

func (w *World) Width() int {
	return w.screenW
}

func (w *World) Height() int {
	return w.screenH
}

func (w *World) Update() error {
	w.state = w.state.Update(w)
	return nil
}

func (w *World) Draw(display *ebiten.Image) {
	w.state.Draw(display)
}

func (w *World) Layout(outsideWidth, outsideHeight int) (width, height int) {
	return w.screenW, w.screenH
}
