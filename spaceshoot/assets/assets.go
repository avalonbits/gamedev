package assets

import (
	"embed"
	"image"
	"io/fs"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

//go:embed *.png *.ttf meteors/*.png
var assets embed.FS

var (
	Player      = load("player.png")
	LaserSprite = load("laser.png")
	Meteors     = loadAll("meteors/*.png")
	ScoreFont   = loadFont("SpaceMono-Regular.ttf")
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
