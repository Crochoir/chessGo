package pieces

import (
	_ "image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Coords struct {
	x float64
	y float64
}

type Rank struct {
	rank   string
	motion [][]string
	image  *ebiten.Image
}

type Piece struct {
	rank     *Rank
	position *Coords
}

func (p *Piece) Initialize(imagePath string, x, y float64) error {
	if p.rank == nil {
		p.rank = &Rank{}
	}
	piecefile, err := os.Open(imagePath)
	if err != nil {
		log.Fatal(err)
	}
	defer piecefile.Close()

	img, _, err := ebitenutil.NewImageFromReader(piecefile)
	if err != nil {
		log.Fatal(err)
	}
	p.rank.image = img
	p.position = &Coords{x: x, y: y}
	return nil
}

func (p *Piece) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.75, 0.75)
	op.GeoM.Translate(float64(p.position.x), float64(p.position.y))
	screen.DrawImage(p.rank.image, op)

}

func (p *Piece) Update() error {
	return nil
}
