package assets

import (
	"embed"
	"fmt"
	"image"
	"io/fs"

	_ "image/png"

	ebiten "github.com/hajimehoshi/ebiten/v2"
)

//go:embed *.png meteors/*.png
var sprites embed.FS

var (
	Player  = load("player.png")
	Meteors = loadAll("meteors/*.png")
)

func load(name string) *ebiten.Image {
	f, err := sprites.Open(name)
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
	files, err := fs.Glob(sprites, pattern)
	if err != nil {
		panic(err)
	}

	meteors := make([]*ebiten.Image, 0, len(files))
	for _, file := range files {
		meteors = append(meteors, load(file))
	}
	fmt.Println(files)

	return meteors
}
