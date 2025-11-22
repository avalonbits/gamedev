package assets

import (
	"embed"
	"fmt"
	"image"
	"io/fs"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.png bricks/*.png levels/*.txt
var assets embed.FS

var (
	Paddle = load("paddle_blue.png")
	Bricks = loadAll("bricks/*.png")
	Levels = loadLevels("levels/*.txt")
)

func load(name string) *ebiten.Image {
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

func loadAll(pattern string) []*ebiten.Image {
	files, err := fs.Glob(assets, pattern)
	if err != nil {
		panic(err)
	}

	meteors := make([]*ebiten.Image, 0, len(files))
	for _, file := range files {
		meteors = append(meteors, load(file))
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
	x, y   int
	sprite *ebiten.Image
}

func (b Brick) Position() (int, int) {
	return b.x, b.y
}

func (b Brick) Sprite() *ebiten.Image {
	return b.sprite
}

func (l *Level) Index() int {
	return l.idx
}

func (l *Level) Bricks() []Brick {
	return l.bricks
}

func loadLevels(pattern string) []Level {
	fmt.Println("load levels")
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
				sprite = Bricks[color-1]
			}
		}
		fmt.Print("\n")
	}

	return bricks
}
