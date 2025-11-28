package objects

import (
	"math"
	"math/rand/v2"

	"github.com/avalonbits/gamedev/breakout/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	sprite   *ebiten.Image
	position vector
	movement vector
	velocity float64
	playArea *PlayArea
	paddle   *Paddle
	bricks   *Bricks
}

func NewBall(sprite *ebiten.Image, playArea *PlayArea, paddle *Paddle, bricks *Bricks) *Ball {
	bounds := sprite.Bounds()
	halfW := float64(bounds.Dx()) / 2

	position := vector{
		X: (playArea.Rect().Width/2 + playArea.Rect().X) - halfW,
		Y: (playArea.Rect().MaxY() - 32) - 256,
	}

	return &Ball{
		sprite:   sprite,
		position: position,
		velocity: 4,
		movement: vector{X: 0, Y: 1},
		playArea: playArea,
		paddle:   paddle,
	}
}

func (b *Ball) Update(world *game.World) {
	paddle := b.paddle.Rect()
	ball := b.Rect()

	collide := b.collidePaddle(ball, paddle)
	collide = collide || b.collidePlayArea(ball, b.playArea.Rect())

	b.position.X += (b.movement.X * b.velocity)
	b.position.Y += (b.movement.Y * b.velocity)
}

var paddleAngle = []float64{
	15 * math.Pi / 180,
	30 * math.Pi / 180,
	60 * math.Pi / 180,
	60 * math.Pi / 180,
	30 * math.Pi / 180,
	15 * math.Pi / 180,
}
var paddleAngleDirection = []float64{
	-1, -1, -1, 1, 1, 1,
}

func (b *Ball) collidePaddle(ball game.Rect, paddle game.Rect) bool {
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
	// the ball hit the paddle.
	//
	// The badlle will be devided in 6 segments.
	// - 2-Middle: +-60 degrees
	// - 2-outer:  +-30 degrees
	// - 2-edge:   +-15 degrees
	//
	// all of them we add [0,2] degrees of jitter

	pos := ball.MaxX() - paddle.X
	segmentSize := paddle.Width / 6
	idx := int(min(5, pos/segmentSize))
	angle := paddleAngle[idx] + rand.Float64()*2*math.Pi/180
	dirX := paddleAngleDirection[idx]

	b.movement = vector{
		X: math.Cos(angle) * dirX,
		Y: -math.Sin(angle),
	}.Normalize()

	return true
}

func (b *Ball) collidePlayArea(ball game.Rect, playArea game.Rect) bool {
	collide := true
	if ball.MaxX() >= playArea.MaxX() || ball.X <= playArea.X {
		b.movement.X = -b.movement.X
	} else if ball.Y <= playArea.Y {
		b.movement.Y = -b.movement.Y
	} else {
		collide = false
	}

	return collide
}

func (b *Ball) Draw(display *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(b.position.X, b.position.Y)
	display.DrawImage(b.sprite, op)
}

func (b *Ball) Rect() game.Rect {
	bounds := b.sprite.Bounds()

	return game.NewRect(
		b.position.X,
		b.position.Y,
		float64(bounds.Dx()),
		float64(bounds.Dy()),
	)
}

func (b *Ball) Intersects(bounds game.Bounds) bool {
	return b.Rect().Intersects(bounds.Rect())
}
