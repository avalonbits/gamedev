package states

import (
	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/avalonbits/gamedev/breakout/objects"
	"github.com/hajimehoshi/ebiten/v2"
)

type Object interface {
	Update(*game.World, game.State)
	Draw(*ebiten.Image)
	Reset()
}

type main struct {
	objects   []Object
	nextState game.State
	reset     bool
}

func (m *main) Update(world *game.World) game.State {
	for _, obj := range m.objects {
		obj.Update(world, m)
	}
	if m.reset {
		m.reset = false
		for _, obj := range m.objects {
			obj.Reset()
		}
	}
	return m.nextState
}

func (m *main) Draw(display *ebiten.Image) {
	for _, obj := range m.objects {
		obj.Draw(display)
	}
}

func (m *main) Next(state game.State) {
	m.nextState = state
}

func (m *main) Reset() {
	m.reset = true
}

type Game struct {
	*main
}

func NewGame() game.State {
	playArea := objects.NewPlayArea(16, assets.DefaultBackground)
	levels := objects.NewLevels(assets.Levels, playArea)
	paddle := objects.NewPaddle(assets.Paddle, playArea)
	ball := objects.NewBall(
		assets.Ball, playArea, paddle, levels, assets.PingSE, assets.PongSE, assets.ClingSE, NewMenu,
	)

	g := Game{
		main: &main{
			objects: []Object{playArea, levels, paddle, ball},
		},
	}
	g.nextState = g

	return &g
}

type Menu struct {
	*main
}

func NewMenu() game.State {
	menu := Menu{
		main: &main{
			objects: []Object{
				objects.NewMenu(assets.GameMenu),
				objects.NewMenuSelector(assets.MenuSelector, assets.IntroSong, NewGame),
			},
		},
	}
	menu.nextState = menu

	return &menu
}
