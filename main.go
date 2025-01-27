package main

import (
	"fmt"
	"bytes"
	"image"
	"image/color"
	"log"
	"github.com/Crochoir/chessGo/pieces"
	"github.com/Crochoir/chessGo/board"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth = 800
	screenHeight = 800
)

var (
	board *ebiten.Image
	piece *ebiten.Image
)

type Game struct {
	turn bool
}

func (g *Game) Update() error {
	if g.turn == true {
		g.turn == false
	} else {
		g.turn == true
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawnOptions{}
	screen.DrawImage(board, op)
}

func (g *Game) Layout(Width, Height int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(rect-8x8_png))
	if err != nil {
		log.Fatal(err)
	}
	boardImage = ebiten.NewImageFromImage(img)

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("chess")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
