package assets

import (
	"embed"
	"image"
	"io"
	"io/fs"
	"math"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.png bricks/*.png levels/*.txt sounds/*.ogg
var assets embed.FS

var (
	IntroSong         = loadSound("sounds/intro-song.ogg")
	GameMenu          = loadImage("game-menu.png")
	MenuSelector      = loadImage("ball.png")
	Ball              = loadImage("ball_12x12.png")
	Paddle            = loadImage("paddle_blue.png")
	Bricks            = loadImages("bricks/*.png")
	Levels            = loadLevels("levels/*.txt")
	DefaultBackground = loadImage("default_background.png")
	PingSE            = loadSound("sounds/ping.ogg")
	PongSE            = loadSound("sounds/pong.ogg")
	ClingSE           = loadSound("sounds/cling.ogg")
)

type SoundEffect struct {
	player *audio.Player
}

func (se SoundEffect) Play() {
	if err := se.player.Rewind(); err != nil {
		panic(err)
	}
	se.player.Play()
}

func (se SoundEffect) IsPlaying() bool {
	return se.player.IsPlaying()
}

func (se SoundEffect) ChangeVolume(inc float64) {
	next := max(0.0, min(1.0, inc+se.player.Volume()))
	se.player.SetVolume(next)
}

func (se SoundEffect) Stop() {
	se.player.Pause()
}

var audioContext = audio.NewContext(44_100)

func loadSound(name string) SoundEffect {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	stream, err := vorbis.DecodeWithSampleRate(44_100, f)
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(stream)
	if err != nil {
		panic(err)
	}

	player := audioContext.NewPlayerFromBytes(data)
	player.SetVolume(1.0)

	return SoundEffect{player: player}
}

func loadImage(name string) *ebiten.Image {
	f, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func loadImages(pattern string) []*ebiten.Image {
	files, err := fs.Glob(assets, pattern)
	if err != nil {
		panic(err)
	}

	meteors := make([]*ebiten.Image, 0, len(files))
	for _, file := range files {
		meteors = append(meteors, loadImage(file))
	}

	return meteors
}

func loadFont(name string) font.Face {
	f, err := assets.ReadFile(name)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(f)
	if err != nil {
		panic(err)
	}

	face, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    48,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return face
}

type Level struct {
	idx    int
	bricks []Brick
}

type Brick struct {
	x, y     int
	power    int
	hitCount int
	sprite   *ebiten.Image
}

func (b Brick) Position() (int, int) {
	return b.x, b.y
}

func (b Brick) Sprite() *ebiten.Image {
	return b.sprite
}

func (b Brick) HitCount() int {
	return b.hitCount
}

func (l *Level) Index() int {
	return l.idx
}

func (l *Level) Bricks() []Brick {
	return l.bricks
}

func loadLevels(pattern string) []Level {
	files, err := fs.Glob(assets, pattern)
	if err != nil {
		panic(err)
	}

	levels := make([]Level, 0, len(files))
	for idx, fileName := range files {
		content, err := assets.ReadFile(fileName)
		if err != nil {
			panic(err)
		}

		levels = append(levels, Level{
			idx:    idx + 1,
			bricks: parseBricks(content),
		})
	}

	return levels
}

func parseBricks(content []byte) []Brick {
	brickAssets := loadImages("bricks/*.png")

	bricks := make([]Brick, 0, 13*13)
	cIdx := 0
	for y := 0; y < 13; y++ {
		for x := 0; x < 13; x++ {
			power := int(content[cIdx] - '0')
			cIdx++

			color := int(content[cIdx] - '0')
			cIdx += 2

			var sprite *ebiten.Image
			if color > 0 {
				sprite = brickAssets[color-1]
			}

			var hitCount int
			switch color {
			case 0:
				hitCount = 0
			case 6:
				hitCount = 3
			case 7:
				hitCount = math.MaxInt
			default:
				hitCount = 1
			}

			pixelX, pixelY := x, y
			if sprite != nil {
				pixelX *= sprite.Bounds().Max.X
				pixelY *= sprite.Bounds().Max.Y
			}
			bricks = append(bricks, Brick{
				x:        pixelX,
				y:        pixelY,
				hitCount: hitCount,
				power:    power,
				sprite:   sprite,
			})
		}
	}

	return bricks
}
