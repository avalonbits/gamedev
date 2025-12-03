package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type State interface {
	Update(world *World) State
	Draw(*ebiten.Image)
	Next(state State)
	Reset()
}

type StateFactory func(world *World) State

type World struct {
	screenW        int
	screenH        int
	state          State
	availableSlots []int
	next           int
	gamepads       []ebiten.GamepadID
	buttons        []ebiten.StandardGamepadButton
}

func NewWorld(
	title string,
	screenW int,
	screenH int,
) *World {
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowSize(screenW, screenH)
	ebiten.SetVsyncEnabled(true)

	w := &World{
		screenW:  screenW,
		screenH:  screenH,
		gamepads: make([]ebiten.GamepadID, 0, 8),
		buttons:  make([]ebiten.StandardGamepadButton, 0, ebiten.StandardGamepadButtonMax),
	}
	return w
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

func (w *World) Gamepad() ebiten.GamepadID {
	w.gamepads = inpututil.AppendJustConnectedGamepadIDs(w.gamepads[:0])
	if len(w.gamepads) == 0 {
		return 0
	}
	return w.gamepads[0]
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

func (w *World) PressLeft() bool {
	return ebiten.IsKeyPressed(ebiten.KeyLeft) ||
		ebiten.IsStandardGamepadButtonPressed(w.Gamepad(), ebiten.StandardGamepadButtonLeftLeft)
}

func (w *World) PressRight() bool {
	return ebiten.IsKeyPressed(ebiten.KeyRight) ||
		ebiten.IsStandardGamepadButtonPressed(w.Gamepad(), ebiten.StandardGamepadButtonLeftRight)
}

func (w *World) JustPressedUp() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyUp) ||
		inpututil.IsStandardGamepadButtonJustPressed(w.Gamepad(), ebiten.StandardGamepadButtonLeftTop) ||
		w.VerticalAxis() <= -1.0

}

func (w *World) JustPressedDown() bool {
	return inpututil.IsKeyJustPressed(ebiten.KeyDown) ||
		inpututil.IsStandardGamepadButtonJustPressed(w.Gamepad(), ebiten.StandardGamepadButtonLeftBottom) ||
		w.VerticalAxis() >= 1.0

}

var actionButtons = map[ebiten.StandardGamepadButton]bool{
	ebiten.StandardGamepadButtonRightTop:    true,
	ebiten.StandardGamepadButtonRightLeft:   true,
	ebiten.StandardGamepadButtonRightRight:  true,
	ebiten.StandardGamepadButtonRightBottom: true,
}

func (w *World) JustPressedAction() bool {
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		return true
	}

	w.buttons = inpututil.AppendJustPressedStandardGamepadButtons(w.Gamepad(), w.buttons[:0])
	for _, button := range w.buttons {
		if actionButtons[button] {
			return true
		}
	}

	return false
}

func (w *World) HorizontalAxis() float64 {
	value := ebiten.StandardGamepadAxisValue(w.Gamepad(), ebiten.StandardGamepadAxisLeftStickHorizontal)
	if value <= -0.1 || value >= 0.1 {
		return value
	} else {
		return 0.0
	}
}

func (w *World) VerticalAxis() float64 {
	value := ebiten.StandardGamepadAxisValue(w.Gamepad(), ebiten.StandardGamepadAxisLeftStickVertical)
	if value <= -0.1 || value >= 0.1 {
		return value
	} else {
		return 0.0
	}
}
