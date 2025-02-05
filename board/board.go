package board

import (
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type coords struct {
	x string
	y string
}

type Board struct {
	image  *ebiten.Image
	layout *coords
}

func (b *Board) Initialize(imagePath string) error {
	boardfile, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer boardfile.Close()

	img, _, err := ebitenutil.NewImageFromReader(boardfile)
	if err != nil {
		log.Fatal(err)
	}
	b.image = img
	b.layout = &coords{x: "0", y: "0"}
	return nil
}

func (b *Board) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	screen.DrawImage(b.image, op)
}

func (b *Board) Update() error {
	return nil
}
