package objects

import (
	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Levels struct {
	levels    []level
	currLevel int
	playArea  *PlayArea
}

type level struct {
	bricks []*Brick
}

func NewLevels(levels []assets.Level, playArea *PlayArea) *Levels {
	brickLevels := []level{}

	for _, lvl := range levels {
		brickLevel := []*Brick{}
		for _, brick := range lvl.Bricks() {
			x, y := brick.Position()
			x += int(playArea.Rect().X)
			y += int(playArea.Rect().Y)

			brickLevel = append(brickLevel, NewBrick(x, y, brick.HitCount(), brick.Sprite()))
		}
		brickLevels = append(brickLevels, level{bricks: brickLevel})
	}

	return &Levels{
		levels:    brickLevels,
		currLevel: 0,
		playArea:  playArea,
	}
}

func (l *Levels) Update(world *game.World) {
}

func (l *Levels) Draw(display *ebiten.Image) {
	level := l.levels[l.currLevel]
	for _, brick := range level.bricks {
		brick.Draw(display)
	}
}

func (l *Levels) HitBrick(ball game.Rect) (int, bool, bool) {
	level := l.levels[l.currLevel]
	changeX := false
	changeY := false

	var hitCount int
	for _, brick := range level.bricks {
		if brick.sprite == nil || brick.hitCount <= 0 {
			continue
		}
		bounds := brick.Rect()
		v1 := ball.MaxY() >= bounds.Y && ball.MaxY() <= bounds.MaxY()
		v2 := ball.Y >= bounds.Y && ball.Y < bounds.MaxY()
		if !v1 && !v2 {
			continue
		}

		h1 := ball.MaxX() >= bounds.X && ball.MaxX() <= bounds.MaxX()
		h2 := ball.X >= bounds.X && ball.X <= bounds.MaxX()
		if !h1 && !h2 {
			continue
		}

		// It's a hit!
		brick.hitCount--
		hitCount = brick.hitCount
		changeX = !(ball.X >= bounds.X && ball.MaxX() <= bounds.MaxX())
		changeY = !(ball.Y >= bounds.Y && ball.MaxY() <= bounds.MaxY())

		break
	}

	return hitCount, changeX, changeY
}

func (l *Levels) Rect() game.Rect {
	return game.Rect{}
}

func (l *Levels) Intersects(bounds game.Bounds) bool {
	return false
}
