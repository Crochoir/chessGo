package board

import (
<<<<<<< HEAD
	_ "image/png"
	"log"
	"os"
=======
	"image/color"
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269

	"github.com/hajimehoshi/ebiten/v2"
)

type coords struct {
<<<<<<< HEAD
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
=======
	X string
	Y string
}

type Board struct {
	layout *coords
	width  int
	height int
}

func (b *Board) Initialize(width, height int) {
	b.width = width
	b.height = height
	b.layout = &coords{X: "0", Y: "0"}
}

func (b *Board) Draw(screen *ebiten.Image) {
	// Define the size of the board cells
	cellWidth := b.width / 8
	cellHeight := b.height / 8

	// Create images for light and dark squares
	lightSquare := ebiten.NewImage(cellWidth, cellHeight)
	darkSquare := ebiten.NewImage(cellWidth, cellHeight)

	lightSquare.Fill(color.RGBA{255, 255, 255, 255})
	darkSquare.Fill(color.RGBA{125, 125, 125, 255})

	// Loop to draw each square on the board
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			var squareImage *ebiten.Image
			if (row+col)%2 == 0 {
				squareImage = lightSquare
			} else {
				squareImage = darkSquare
			}

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(col*cellWidth), float64(row*cellHeight))
			screen.DrawImage(squareImage, op)
		}
	}
>>>>>>> c7fa942b5252e2d25417aa39e682288a11c76269
}

func (b *Board) Update() error {
	return nil
}
