package objects

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"

	"github.com/avalonbits/gamedev/breakout/assets"
	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	sprite     *ebiten.Image
	position   vector
	movement   vector
	velocity   float64
	playArea   *PlayArea
	paddle     *Paddle
	levels     *Levels
	ping       assets.SoundEffect
	pong       assets.SoundEffect
	cling      assets.SoundEffect
	speedTimer *game.Timer
	nextState  func() game.State
}

func NewBall(
	sprite *ebiten.Image,
	playArea *PlayArea,
	paddle *Paddle,
	levels *Levels,
	ping assets.SoundEffect,
	pong assets.SoundEffect,
	cling assets.SoundEffect,
	nextState func() game.State,
) *Ball {
	b := &Ball{
		sprite:     sprite,
		playArea:   playArea,
		paddle:     paddle,
		levels:     levels,
		ping:       ping,
		pong:       pong,
		cling:      cling,
		speedTimer: game.NewTimer(30 * time.Second),
		nextState:  nextState,
	}
	b.Restart()

	return b
}

func (b *Ball) Reset() {
	bounds := b.sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := vector{
		X: (b.playArea.Rect().Width/2 + b.playArea.Rect().X) - halfW,
		Y: (b.playArea.Rect().MaxY() - 32) - 192,
	}

	b.position = position
	b.movement = vector{X: 0, Y: 1}
}

func (b *Ball) Restart() {
	b.Reset()
	b.velocity = 5
}

func (b *Ball) Update(world *game.World, state game.State) {
	b.speedTimer.Update()
	if b.speedTimer.IsReady() {
		b.speedTimer.Reset()
		b.velocity++
	}

	ball := b.Rect()
	playArea := b.playArea.Rect()

	collide := b.collidePaddle()
	collide = collide || b.collidePlayArea(ball, playArea, state)
	collide = collide || b.collideBricks(ball, state)
	if collide {
		fmt.Println(b.velocity)
	}

	b.position.X += (b.movement.X * b.velocity)
	b.position.X = max(playArea.X, min(b.position.X, playArea.MaxX()-ball.Width))
	b.position.Y += (b.movement.Y * b.velocity)
	b.position.Y = max(b.position.Y, playArea.Y)
}

var paddleAngle = []float64{
	40 * math.Pi / 180,
	50 * math.Pi / 180,
	65 * math.Pi / 180,
	85 * math.Pi / 180,
	65 * math.Pi / 180,
	50 * math.Pi / 180,
	40 * math.Pi / 180,
}

func (b *Ball) collidePaddle() bool {
	ball := b.Rect()
	paddle := b.paddle.Rect()

	if b.movement.Y < 0 {
		// We are already going up, no need to check collision
		return false
	}

	verticalPaddle := paddle.Y <= ball.MaxY() && paddle.Y >= ball.Y
	if !verticalPaddle {
		return false
	}

	horizontalPaddle := ball.MaxX() >= paddle.X && ball.X <= paddle.MaxX()
	if !horizontalPaddle {
		return false
	}

	// Now that we know it collided, we need to determine the angle, which is a function of where
	// the ball hit the paddle + [0,2] degrees of jitter.

	segmentCount := float64(len(paddleAngle))
	pos := ball.MaxX() - paddle.X - ball.Width/2 // center of the ball
	segmentSize := paddle.Width / segmentCount

	idx := int(min(segmentCount-1, pos/segmentSize))
	angle := paddleAngle[idx] + rand.Float64()*2*math.Pi/180
	paddleDir := b.paddle.Direction()

	dirX := 1.0
	if b.movement.X < 0.0 {
		dirX = -1.0
	}
	if paddleDir != 0 && dirX != paddleDir {
		dirX = -dirX
	}

	b.movement = vector{
		X: math.Cos(angle) * dirX,
		Y: -math.Sin(angle),
	}.Normalize()

	b.ping.Play()

	return true
}

func (b *Ball) collidePlayArea(ball rect, playArea rect, state game.State) bool {
	if ball.MaxY() >= playArea.MaxY() {
		state.Reset()
		return true
	}

	collide := true
	if ball.MaxX() >= playArea.MaxX() || ball.X <= playArea.X {
		fmt.Println(b.movement, ball)
		b.movement.X = -b.movement.X
	} else if b.movement.Y <= 0 && ball.Y <= playArea.Y {
		b.movement.Y = -b.movement.Y
	} else {
		collide = false
	}

	return collide
}

func (b *Ball) collideBricks(ball rect, state game.State) bool {
	hitCount, xhit, yhit, levelOver := b.levels.HitBrick(ball)
	if levelOver {
		if b.levels.Next() {
			b.Restart()
			state.Reset()
		} else {
			state.Next(b.nextState())
		}
		return false
	}

	if !xhit && !yhit {
		return false
	}

	if xhit && yhit {
		xhit = math.Abs(b.movement.X) >= math.Abs(b.movement.Y)
		yhit = !xhit
	}

	if yhit {
		b.movement.Y = -b.movement.Y
	} else if xhit {
		b.movement.X = -b.movement.X
	}

	if hitCount == 0 {
		b.pong.Play()
	} else {
		b.cling.Play()
	}

	return true
}

func (b *Ball) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.position.X, b.position.Y)
	display.DrawImage(b.sprite, op)
}

func (b *Ball) Rect() rect {
	bounds := b.sprite.Bounds()

	return NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}
